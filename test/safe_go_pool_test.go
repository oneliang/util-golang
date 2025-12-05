package test

import (
	"context"
	"fmt"
	"github.com/oneliang/util-golang/goroutine"
	"testing"
)

func TestSafeGoPool(t *testing.T) {
	safeGoPool := goroutine.NewSafeGoPool(context.Background(), 2)
	safeGoPool.Go(func() error {
		fmt.Println("i am running")
		panic("a")
		return nil
	})
	err := safeGoPool.Wait()
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
}
