package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"runtime"
	"testing"
)

func TestByteArrayWrapper(t *testing.T) {
	var length uint = 100000
	var byteArrayWrapper = common.NewByteArrayWrapper(length)
	byteArrayWrapper.Write(100, []byte{1, 2, 3, 4})
	readBytes, _ := byteArrayWrapper.Read(100, 4)
	fmt.Println(len(readBytes))
	fmt.Println(readBytes)

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
}
