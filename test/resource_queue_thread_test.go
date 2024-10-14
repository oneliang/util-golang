package test

import (
	"fmt"
	"github.com/oneliang/util-golang/concurrent"
	"testing"
	"time"
)

func TestResourceQueueThread(t *testing.T) {

	resourceQueueThread := concurrent.NewResourceQueueThread[func()](func(resource func()) {
		//fmt.Println(fmt.Sprintf("%+v", resource))
		resource()
	}, nil)
	resourceQueueThread.Start()
	for i := 0; i < 100; i++ {
		resourceQueueThread.AddResource(func() { fmt.Println(i) })
	}
	time.Sleep(100 * time.Second)
	//resourceQueueThread.AddResource(func() { fmt.Println(2) })
	time.Sleep(3 * time.Second)
	resourceQueueThread.Stop()
	//time.Sleep(4 * time.Second)
}
