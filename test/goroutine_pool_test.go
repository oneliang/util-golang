package test

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/goroutine"
	"log"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	pool := goroutine.NewPool(1)
	pool.Start()

	pool.AddTask(func(params ...any) error {
		log.Println(fmt.Sprintf("%v", params))
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 1))
		panic("aa")
		return nil
	}, 1, 2, 3)

	pool.AddTask(func(params ...any) error {
		log.Println(fmt.Sprintf("%v", params))
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 2))
		return nil
	}, 4, 5)

	pool.AddTask(func(params ...any) error {
		log.Println(fmt.Sprintf("%v", params))
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 3))
		return nil
	}, 6, 7)

	pool.AddTask(func(params ...any) error {
		log.Println(fmt.Sprintf("%v", params))
		fmt.Println(fmt.Sprintf("goroutine id:%d, %d", common.GetGoroutineId(), 4))
		return nil
	}, 8, 9, 10)

	time.Sleep(100 * time.Second)
	pool.Stop()
}
