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
	resourceQueueThread.AddResource(func() { fmt.Println(1) })
	time.Sleep(2 * time.Second)
	resourceQueueThread.AddResource(func() { fmt.Println(2) })
	resourceQueueThread.Stop()
	time.Sleep(3 * time.Second)
}
