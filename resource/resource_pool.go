package resource

import (
	"errors"
	"fmt"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/concurrent"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
	"sync"
	"time"
)

type ResourceSource[T any] interface {
	GetResource() *T

	DestroyResource(resource *T) error
}

type resourceStatus[T any] struct {
	resource         *T
	inUse            bool
	lastNotInUseTime int64
}

type stableResourceStatus[T any] struct {
	resource         *T
	lastNotInUseTime int64
}
type ResourcePool[T any] struct {
	Name                    string
	minResourceSize         uint
	maxResourceSize         uint
	maxStableResourceSize   uint
	resourceAliveTime       time.Duration
	threadSleepTime         time.Duration
	resourceSource          ResourceSource[T]
	hasBeenInitialized      bool
	initializeLock          sync.Mutex
	resourceLock            sync.Mutex
	stableResourceLock      sync.Mutex
	resourceStatusArray     []*resourceStatus[T]
	resourceCurrentSize     uint
	loopThread              *concurrent.LoopThread
	stableResourceStatusMap map[int]*stableResourceStatus[T]
	logger                  logging.Logger
}
type ResourcePoolConfig struct {
	MinResourceSize       uint
	MaxResourceSize       uint
	MaxStableResourceSize uint
	ResourceAliveTime     time.Duration
	ThreadSleepTime       time.Duration
}

var (
	defaultResourceConfig = &ResourcePoolConfig{
		MinResourceSize:       1,
		MaxResourceSize:       1,
		MaxStableResourceSize: 1,
		ResourceAliveTime:     0,
		ThreadSleepTime:       5 * time.Minute,
	}
)

func NewResourcePool[T any](name string, resourceSource ResourceSource[T], resourcePoolConfig *ResourcePoolConfig) *ResourcePool[T] {
	var config = resourcePoolConfig
	if config == nil {
		config = defaultResourceConfig
	}
	return &ResourcePool[T]{
		Name:                    name,
		minResourceSize:         config.MinResourceSize,
		maxResourceSize:         config.MaxResourceSize,
		maxStableResourceSize:   config.MaxStableResourceSize,
		resourceAliveTime:       config.ResourceAliveTime,
		threadSleepTime:         config.ThreadSleepTime,
		resourceSource:          resourceSource,
		hasBeenInitialized:      false,
		stableResourceStatusMap: make(map[int]*stableResourceStatus[T]),
		logger:                  logging.LoggerManager.GetLoggerByPattern("ResourcePool"),
	}
}

func (this *ResourcePool[T]) initialize() {
	if this.hasBeenInitialized {
		return
	}
	//will invoke second when return
	defer this.initializeLock.Unlock()
	//will invoke first when return
	defer func() {
		this.hasBeenInitialized = true
	}()

	this.initializeLock.Lock()
	if this.hasBeenInitialized { //double check
		return //return will trigger finally, but use unlock safety
	}
	this.resourceStatusArray = make([]*resourceStatus[T], this.maxResourceSize)
	for i := 0; i < int(this.minResourceSize); i++ {
		resource := this.resourceSource.GetResource()
		resourceStatusItem := &resourceStatus[T]{
			resource:         resource,
			inUse:            false,
			lastNotInUseTime: 0,
		}
		this.resourceStatusArray[i] = resourceStatusItem
		this.resourceCurrentSize++
	}
	if this.maxStableResourceSize == 0 {
		this.maxStableResourceSize = 1
	}
	this.loopThread = concurrent.NewLoopThread(func() {
		this.run()
	})
	err := this.loopThread.Start()
	if err != nil {
		this.logger.Error(constants.STRING_ERROR, err)
	}
}

// GetResource .
func (this *ResourcePool[T]) GetResource() (*T, error) {
	if !this.hasBeenInitialized {
		this.initialize()
	}
	defer this.resourceLock.Unlock()
	this.resourceLock.Lock()
	var resource *T = nil
	this.logger.Info("resource pool name:%s, resource current size:%d", this.Name, this.resourceCurrentSize)
	if this.resourceCurrentSize > 0 {
		for _, resourceStatusItem := range this.resourceStatusArray {
			if resourceStatusItem == nil || resourceStatusItem.inUse {
				continue
			}
			resource = resourceStatusItem.resource
			resourceStatusItem.inUse = true
			break
		}
	}
	if resource != nil {
		return resource, nil
	}
	if this.resourceCurrentSize < this.maxResourceSize {
		for index := 0; index < len(this.resourceStatusArray); index++ {
			resourceStatusItem := this.resourceStatusArray[index]
			if resourceStatusItem != nil {
				continue
			}
			resource = this.resourceSource.GetResource()
			oneResourceStatus := &resourceStatus[T]{
				resource:         resource,
				inUse:            true,
				lastNotInUseTime: 0,
			}
			this.resourceStatusArray[index] = oneResourceStatus
			this.resourceCurrentSize++
			break
		}
	} else {
		err := errors.New(fmt.Sprintf("The resource pool is max, current:%d", this.resourceCurrentSize))
		this.logger.Error(constants.STRING_BLANK, err)
		return nil, err
	}
	if resource == nil {
		return nil, errors.New("resource can not be nil")
	}
	return resource, nil
}

// GetStableResource .
func (this *ResourcePool[T]) GetStableResource() *T {
	if !this.hasBeenInitialized {
		this.initialize()
	}
	var stableResource *T = nil
	defer this.stableResourceLock.Unlock()
	this.stableResourceLock.Lock()
	this.logger.Info("resource pool name:%s, stable resource current size:%d, stable resource map:%v", this.Name, len(this.stableResourceStatusMap), this.stableResourceStatusMap)
	var stableResourceStatusKey = this.getStableResourceStatusKey()
	stableResourceStatusItem, ok := this.stableResourceStatusMap[stableResourceStatusKey]
	if ok { //exist
		stableResource = stableResourceStatusItem.resource
	} else { //not exist
		stableResource = this.resourceSource.GetResource()
		this.stableResourceStatusMap[stableResourceStatusKey] = &stableResourceStatus[T]{
			resource:         stableResource,
			lastNotInUseTime: 0, //initialize the not in use time
		}
	}
	return stableResource
}

// getStableResourceStatusKey .
func (this *ResourcePool[T]) getStableResourceStatusKey() int {
	var currentThreadHashCode = common.GetGoroutineId()
	return int(uint(currentThreadHashCode) % this.maxStableResourceSize)
}

// run .
func (this *ResourcePool[T]) run() {
	time.Sleep(this.threadSleepTime)
	this.logger.Debug("--The resource pool is:'%s', before clean resources number:%s, stable resource number:%s", this.Name, this.resourceCurrentSize, len(this.stableResourceStatusMap))
	this.clean()
	this.logger.Debug("--The resource pool is:'%s', after clean resources number:%s, stable resource number:%s", this.Name, this.resourceCurrentSize, len(this.stableResourceStatusMap))
}
func (this *ResourcePool[T]) clean() {
	defer this.resourceLock.Unlock()
	this.resourceLock.Lock()
	for index, resourceStatusItem := range this.resourceStatusArray {
		if resourceStatusItem == nil || resourceStatusItem.inUse {
			continue
		}
		var lastTime = resourceStatusItem.lastNotInUseTime
		var currentTime = time.Now().UnixMilli()
		var resource = resourceStatusItem.resource
		if currentTime-lastTime >= int64(this.resourceAliveTime) {
			this.realDestroyResource(index, resource)
		}
		//stable resource
		defer this.stableResourceLock.Unlock()
		this.stableResourceLock.Lock()
		for stableResourceKey, stableResourceStatusItem := range this.stableResourceStatusMap {
			lastTime = stableResourceStatusItem.lastNotInUseTime
			currentTime = time.Now().UnixMilli()
			resource = stableResourceStatusItem.resource
			if currentTime-lastTime >= int64(this.resourceAliveTime) {
				this.realDestroyStableResource(stableResourceKey, resource)
			}
		}
	}
}

// ReleaseResource .
func (this *ResourcePool[T]) ReleaseResource(resource *T, destroy bool) {
	if resource == nil {
		return
	}
	for index, resourceStatusItem := range this.resourceStatusArray {
		if resourceStatusItem == nil {
			continue
		}
		//find the resource and set in use false
		if resource == resourceStatusItem.resource {
			if destroy {
				defer this.resourceLock.Unlock()
				this.resourceLock.Lock()
				this.realDestroyResource(index, resource)
			} else {
				resourceStatusItem.inUse = false
				resourceStatusItem.lastNotInUseTime = time.Now().UnixMilli()
			}
			break
		}
	}
}

// ReleaseStableResource .
func (this *ResourcePool[T]) ReleaseStableResource(stableResource *T, destroy bool) {
	if stableResource == nil {
		return
	}
	var stableResourceStatusKey = this.getStableResourceStatusKey()
	stableResourceStatusItem, ok := this.stableResourceStatusMap[stableResourceStatusKey]
	if ok { // exist
		if stableResourceStatusItem != nil && stableResource == stableResourceStatusItem.resource {
			if destroy {
				this.stableResourceLock.Unlock()
				this.stableResourceLock.Lock()
				this.realDestroyStableResource(stableResourceStatusKey, stableResource)
			} else {
				stableResourceStatusItem.lastNotInUseTime = time.Now().UnixMilli()
			}
		} else {
			this.logger.Error("release stable resource, stable resource status is null or stable resource is not the same, stable resource status:%v, stable resource:%v", nil, stableResourceStatusItem, stableResource)
		}
	} else {
		this.logger.Error("release stable resource, this stable resource maybe haven't got from method named stableResource ? getStableResource and releaseStableResource maybe not in same thread, or stableResource had been cleaned,  Stable resource:%s", nil, stableResource)
	}
}

func (this *ResourcePool[T]) realDestroyResource(index int, resource *T) {
	//destroy
	err := this.resourceSource.DestroyResource(resource)
	if err != nil {
		this.logger.Error("Read destroy resource had error:%v", err, err)
	}
	this.resourceStatusArray[index] = nil
	this.resourceCurrentSize--
}

func (this *ResourcePool[T]) realDestroyStableResource(stableResourceStatusKey int, resource *T) {
	//destroy
	err := this.resourceSource.DestroyResource(resource)
	if err != nil {
		this.logger.Error("Read destroy stable resource had error:%v", err, err)
	}
	delete(this.stableResourceStatusMap, stableResourceStatusKey)
}

func (this *ResourcePool[T]) Destroy() {
	this.clean()
	if this.loopThread != nil {
		err := this.loopThread.Stop()
		if err != nil {
			this.logger.Error(constants.STRING_ERROR, err)
		}
		this.loopThread = nil
		this.hasBeenInitialized = false
	}
}

func (this *ResourcePool[T]) UseResource(block func(resource *T, err error), destroy bool) {
	resource, err := this.GetResource()
	block(resource, err)
	defer this.ReleaseResource(resource, destroy)
}

func (this *ResourcePool[T]) UseStableResource(block func(resource *T), destroy bool) {
	resource := this.GetStableResource()
	block(resource)
	defer this.ReleaseStableResource(resource, destroy)
}
