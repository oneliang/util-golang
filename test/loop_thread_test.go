package test

import (
	"fmt"
	"github.com/oneliang/util-golang/concurrent"
	"testing"
	"time"
)

func TestLoopThread(t *testing.T) {
	loopThread := concurrent.NewLoopThread(func() {
		fmt.Println("running")
		time.Sleep(1 * time.Second)
	})
	_ = loopThread.Start()
	time.Sleep(10 * time.Second)
	_ = loopThread.Stop()
	time.Sleep(1 * time.Second)
}
