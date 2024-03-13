package test

import (
	"fmt"
	"github.com/oneliang/util-golang/resource"
	"testing"
	"time"
)

type Connection struct {
	Name string
}
type Source struct {
	count int
}

func (this *Source) GetResource() *Connection {
	this.count++
	return &Connection{
		Name: fmt.Sprintf("connection_index:%d", this.count),
	}
}
func (this *Source) DestroyResource(resource *Connection) error {
	return nil
}
func TestResourcePool(t *testing.T) {
	resourcePool := resource.NewResourcePool[Connection]("resourcePool", &Source{}, &resource.ResourcePoolConfig{
		MaxResourceSize: 1,
	})
	for i := 0; i < 50; i++ {
		go func() {
			resourceItem, err := resourcePool.GetResource()
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			//fmt.Println(resourceItem.Name)
			resourcePool.ReleaseResource(resourceItem, false)
			//fmt.Println(resourceItem)
		}()
	}
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		go func() {
			resourceItem, err := resourcePool.GetResource()
			if err != nil {
				panic(err)
			}
			fmt.Println(resourceItem.Name)
			resourcePool.ReleaseResource(resourceItem, false)
			//fmt.Println(resourceItem)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("-------------------")
	for i := 0; i < 5; i++ {
		go func() {
			resourceItem := resourcePool.GetStableResource()
			fmt.Println(resourceItem.Name)
			resourcePool.ReleaseStableResource(resourceItem, false)
			//fmt.Println(resourceItem)
		}()
	}

	time.Sleep(5 * time.Second)
}
