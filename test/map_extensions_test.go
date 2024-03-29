package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"runtime"
	"testing"
)

type Bean struct {
	Id   int
	Name string
}

func TestMapExtensions(t *testing.T) {
	var length = 100000000
	testBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		testBytes[i] = byte(i)
	}
	fmt.Println(len(testBytes))
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
	inputMap := make(map[int]*Bean)
	inputMap[0] = &Bean{
		Id:   0,
		Name: "0",
	}
	inputMap[1] = &Bean{
		Id:   1,
		Name: "1",
	}
	inputMap[2] = &Bean{
		Id:   2,
		Name: "2",
	}
	otherMap := make(map[int]*Bean)
	otherMap[0] = &Bean{
		Id:   0,
		Name: "0",
	}
	otherMap[1] = &Bean{
		Id:   1,
		Name: "1",
	}
	otherMap[2] = &Bean{
		Id:   2,
		Name: "2",
	}
	list := common.MapDiffersDefault[int, *Bean](inputMap, otherMap)
	fmt.Println(list)
	for _, item := range list {
		fmt.Println(item)
	}
}
