package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"testing"
)

func TestQueue(t *testing.T) {
	// 示例1：整数队列
	fmt.Println("=== 整数队列示例 ===")
	intQueue := common.NewQueue[int](10)

	// 入队
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)
	intQueue.EnqueueBatch(30, 40, 50)

	fmt.Printf("队列长度: %d, 容量: %d\n", intQueue.Size(), intQueue.Capacity())

	// 查看队首队尾
	if front, ok := intQueue.Front(); ok {
		fmt.Printf("队首: %d\n", front)
	}

	if back, ok := intQueue.Back(); ok {
		fmt.Printf("队尾: %d\n", back)
	}

	// 出队操作
	fmt.Println("出队元素:")
	for !intQueue.IsEmpty() {
		if item, ok := intQueue.Dequeue(); ok {
			fmt.Printf("  %d", item)
		}
	}
	fmt.Println()

	// 示例2：字符串队列
	fmt.Println("\n=== 字符串队列示例 ===")
	strQueue := common.NewQueue[string](5)

	strQueue.Enqueue("Hello")
	strQueue.Enqueue("World")
	strQueue.Enqueue("Go")

	fmt.Println("遍历队列:")
	strQueue.ForEach(func(s string) {
		fmt.Printf("  %s", s)
	})
	fmt.Println()

	// 转换为切片
	slice := strQueue.ToSlice()
	fmt.Printf("转换为切片: %v\n", slice)

	// 示例3：自定义类型队列
	fmt.Println("\n=== 自定义类型队列示例 ===")
	type Person struct {
		Name string
		Age  int
	}

	personQueue := common.NewQueue[Person](3)
	personQueue.Enqueue(Person{"Alice", 30})
	personQueue.Enqueue(Person{"Bob", 25})
	personQueue.Enqueue(Person{"Charlie", 35})

	// 清空队列
	fmt.Printf("清空前长度: %d\n", personQueue.Size())
	personQueue.Clear()
	fmt.Printf("清空后长度: %d, 是否为空: %v\n", personQueue.Size(), personQueue.IsEmpty())

	// 性能演示：大量操作时的内存优化
	fmt.Println("\n=== 性能优化演示 ===")
	perfQueue := common.NewQueue[int](0)

	// 入队1000个元素
	for i := 0; i < 1000; i++ {
		perfQueue.Enqueue(i)
	}

	fmt.Printf("入队1000个元素后 - 长度: %d, 容量: %d\n",
		perfQueue.Size(), perfQueue.Capacity())

	// 出队500个元素（触发内存清理）
	for i := 0; i < 500; i++ {
		perfQueue.Dequeue()
	}

	fmt.Printf("出队500个元素后 - 长度: %d, 容量: %d\n",
		perfQueue.Size(), perfQueue.Capacity())

	// 再出队250个元素（可能触发容量缩减）
	for i := 0; i < 250; i++ {
		perfQueue.Dequeue()
	}

	fmt.Printf("再出队250个元素后 - 长度: %d, 容量: %d\n",
		perfQueue.Size(), perfQueue.Capacity())
}
