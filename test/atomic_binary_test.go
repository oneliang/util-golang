package test

import (
	"fmt"
	"github.com/oneliang/util-golang/atomic"
	"github.com/oneliang/util-golang/common"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestAtomicBinary(t *testing.T) {
	var length uint = 1000000
	//var channelList []chan int
	//for i := 0; i < int(length); i++ {
	//	channelList = append(channelList, make(chan int, 1))
	//}
	//fmt.Println(len(channelList))
	//
	//var syncMutexList []sync.Mutex
	//for i := 0; i < int(length); i++ {
	//	syncMutexList = append(syncMutexList, sync.Mutex{})
	//}
	//fmt.Println(len(syncMutexList))
	var begin = time.Now()
	atomicBinary := atomic.NewAtomicBinaryDefault[int](length, 4, func(byteArray []byte) (data int) {
		return common.Primitive.ByteArrayToInt(byteArray)
		//return 0
	}, func(data int) (byteArray []byte) {
		return common.Primitive.IntToByteArray(data)
		//return []byte{0, 0, 0, 0}
	})
	indexWrapper := atomic.NewIndexWrapper(1)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2500000)
	for i := 0; i < 2500000; i++ {
		go func() {
			atomicBinary.Operate(indexWrapper, func() int {
				return 2500001
			}, func(oldData int) (newData int) {
				oldData--
				return oldData
			})
			defer waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println(atomicBinary.Get(indexWrapper))
	var end = time.Now()
	fmt.Println(fmt.Sprintf("cost:%d", end.UnixMilli()-begin.UnixMilli()))
	//var byteArrayWrapper = common.NewByteArrayWrapper(length)
	//byteArrayWrapper.Write(100, []byte{1, 2, 3, 4})
	//readBytes, _ := byteArrayWrapper.Read(100, 4)
	//fmt.Println(len(readBytes))
	//fmt.Println(readBytes)

	//testMap := make(map[int]int, length)
	//for i := 0; i < length; i++ {
	//	testMap[i] = i
	//}
	//fmt.Println(len(testMap))
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem) // 读取内存统计信息到变量mem

	// 输出不同部分的内存信息
	fmt.Printf("Alloc = %v\n", mem.Alloc)           // 已分配但未释放的字节数
	fmt.Printf("TotalAlloc = %v\n", mem.TotalAlloc) // 从开始运行到现在为止所有对象（包括GC）分配的字节数
	fmt.Printf("Sys = %v\n", mem.Sys)               // 系统调用分配的字节数
	fmt.Printf("NumGC = %v\n", mem.NumGC)           // GC发生次数
	return
	atomic.NewAtomicBinaryDefault[int](100, 4, func(byteArray []byte) (data int) {
		return 0
	}, func(data int) (byteArray []byte) {
		return make([]byte, 0)
	})
}
