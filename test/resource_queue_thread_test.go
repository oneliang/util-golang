package test

import (
	"fmt"
	"github.com/oneliang/util-golang/concurrent"
	"testing"
	"time"
)

func TestResourceQueueThread(t *testing.T) {

	resourceQueueThread := concurrent.NewResourceQueueThread[any](func(resource any) {
		fmt.Println(fmt.Sprintf("%+v", resource))
	})
	resourceQueueThread.Start()
	resourceQueueThread.AddResource(1)
	time.Sleep(2 * time.Second)
	resourceQueueThread.AddResource(2)
	resourceQueueThread.Stop()
	time.Sleep(3 * time.Second)
}
