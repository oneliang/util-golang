package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/goroutine"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	pool := goroutine.NewPool(1)
	pool.Start()

	pool.AddTask(func() error {
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 1))
		return nil
	})

	pool.AddTask(func() error {
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 2))
		return nil
	})

	pool.AddTask(func() error {
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 3))
		return nil
	})

	pool.AddTask(func() error {
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 4))
		return nil
	})

	time.Sleep(10000)
	pool.Stop()
}
