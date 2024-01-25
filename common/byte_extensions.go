package common

import (
	"errors"
)

type ByteArrayWrapper struct {
	Size      uint
	byteArray []byte
}

// NewByteArrayWrapper
func NewByteArrayWrapper(size uint) *ByteArrayWrapper {
	var byteArrayWrapper = &ByteArrayWrapper{
		Size:      size,
		byteArray: make([]byte, size),
	}
	return byteArrayWrapper
}

// Read thread unsafe
func (this *ByteArrayWrapper) Read(offset uint, size uint) ([]byte, error) {
	var end = offset + size
	if end > this.Size {
		return nil, errors.New("offset+size is large than Size")
	}
	return this.byteArray[offset:end], nil
}

// Write thread unsafe
func (this *ByteArrayWrapper) Write(offset uint, byteArray []byte) error {
	var end = offset + uint(len(byteArray))
	if end > this.Size {
		return errors.New("offset+size is large than Size")
	}
	copy(this.byteArray[offset:end], byteArray)
	return nil
}
