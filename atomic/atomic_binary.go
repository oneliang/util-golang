package atomic

import (
	"errors"
	"fmt"
	"github.com/oneliang/util-golang/common"
	"sync"
)

const (
	length_exist = 1
)

type IndexWrapper struct {
	Value uint
	lock  sync.Mutex
}

func NewIndexWrapper(value uint) *IndexWrapper {
	return &IndexWrapper{
		Value: value,
		lock:  sync.Mutex{},
	}
}

type AtomicBinary[DATA interface{}] struct {
	initializeSize uint
	expandSize     uint
	dataLength     uint
	// private use, but must use for construct
	byteArrayToData func(byteArray []byte) (data DATA)
	dataToByteArray func(data DATA) (byteArray []byte)
	// private variable
	binaryDataLength     uint
	byteArrayWrapperList []*common.ByteArrayWrapper
	autoExpandLock       *sync.Mutex
}

func NewAtomicBinaryDefault[DATA interface{}](
	initializeSize uint,
	dataLength uint,
	byteArrayToData func(byteArray []byte) (data DATA),
	dataToByteArray func(data DATA) (byteArray []byte),
) *AtomicBinary[DATA] {
	return NewAtomicBinary[DATA](initializeSize, 10000, dataLength, byteArrayToData, dataToByteArray)
}

func NewAtomicBinary[DATA interface{}](
	initializeSize uint,
	expandSize uint,
	dataLength uint,
	byteArrayToData func(byteArray []byte) (data DATA),
	dataToByteArray func(data DATA) (byteArray []byte),
) *AtomicBinary[DATA] {
	binaryDataLength := length_exist + dataLength
	var atomicBinary = &AtomicBinary[DATA]{
		initializeSize:   initializeSize,
		expandSize:       expandSize,
		dataLength:       dataLength,
		byteArrayToData:  byteArrayToData,
		dataToByteArray:  dataToByteArray,
		binaryDataLength: binaryDataLength,
	}
	atomicBinary.byteArrayWrapperList = append(atomicBinary.byteArrayWrapperList, common.NewByteArrayWrapper(atomicBinary.initializeSize*binaryDataLength))
	return atomicBinary
}

func (this *AtomicBinary[DATA]) getSuitableRealIndexAndByteArrayWrapper(index uint) (uint, *common.ByteArrayWrapper) {
	if 0 <= index && index < this.initializeSize {
		return index, this.byteArrayWrapperList[0]
	} else {
		var expandIndexInList = (index-this.initializeSize)/this.expandSize + 1
		var realIndex = (index - this.initializeSize) % this.expandSize
		if int(expandIndexInList) < len(this.byteArrayWrapperList) {
			return realIndex, this.byteArrayWrapperList[expandIndexInList]
		} else { //auto expand
			this.autoExpandLock.Lock()
			defer this.autoExpandLock.Unlock()
			//double check, because size will be changed in previous execute
			if int(expandIndexInList) < len(this.byteArrayWrapperList) {
				return realIndex, this.byteArrayWrapperList[expandIndexInList]
			} else {
				var needToIncreaseSize = int(expandIndexInList) - (len(this.byteArrayWrapperList) - 1)
				for i := 0; i < needToIncreaseSize; i++ {
					this.byteArrayWrapperList = append(this.byteArrayWrapperList, common.NewByteArrayWrapper(this.expandSize*this.binaryDataLength))
				}
				return realIndex, this.byteArrayWrapperList[expandIndexInList]
			}

		}
	}
}
func (this *AtomicBinary[DATA]) OperateCreate(indexWrapper *IndexWrapper, create func() DATA) (DATA, error) {
	return this.Operate(indexWrapper, create, nil)
}

func (this *AtomicBinary[DATA]) Operate(indexWrapper *IndexWrapper, create func() DATA, update func(oldData DATA) (newData DATA)) (DATA, error) {
	realIndex, byteArrayWrapper := this.getSuitableRealIndexAndByteArrayWrapper(indexWrapper.Value)
	indexWrapper.lock.Lock()
	defer indexWrapper.lock.Unlock()
	byteOffset := realIndex * this.binaryDataLength
	existByteArray, _ := byteArrayWrapper.Read(byteOffset, 1)
	existByte := existByteArray[0]
	if existByte > 0 { //exist
		oldDataByteArray, _ := byteArrayWrapper.Read(byteOffset+length_exist, this.dataLength)
		oldData := this.byteArrayToData(oldDataByteArray)
		if update != nil {
			newData := update(oldData)
			newDataByteArray := this.dataToByteArray(newData)
			if len(newDataByteArray) != int(this.dataLength) {
				return newData, errors.New(fmt.Sprintf("new data size is not equal %s when update", this.dataLength))
			}
			_ = byteArrayWrapper.Write(byteOffset+length_exist, newDataByteArray)
			return newData, nil
		} else {
			return oldData, nil
		}
	} else {
		_ = byteArrayWrapper.Write(byteOffset, []byte{1})
		var newData = create()
		if update != nil {
			newData = update(newData)
		}
		var newDataByteArray = this.dataToByteArray(newData)
		if len(newDataByteArray) != int(this.dataLength) {
			return newData, errors.New(fmt.Sprintf("new data size is not equal %s when create", this.dataLength))
		}
		_ = byteArrayWrapper.Write(byteOffset+length_exist, newDataByteArray)
		return newData, nil
	}
}

func (this *AtomicBinary[DATA]) Get(indexWrapper *IndexWrapper) DATA {
	realIndex, byteArrayWrapper := this.getSuitableRealIndexAndByteArrayWrapper(indexWrapper.Value)
	byteOffset := realIndex * this.binaryDataLength
	dataByteArray, _ := byteArrayWrapper.Read(byteOffset+length_exist, this.dataLength)
	return this.byteArrayToData(dataByteArray)
}
