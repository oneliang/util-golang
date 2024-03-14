package concurrent

import (
	"fmt"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
)

type ResourceQueueThread[T interface{}] struct {
	loopThread      *LoopThread
	needToStop      bool
	resourceChannel chan T
	logger          logging.Logger
}

func NewResourceQueueThread[T interface{}]() *ResourceQueueThread[T] {
	resourceQueueThread := &ResourceQueueThread[T]{
		needToStop:      false,
		resourceChannel: make(chan T),
		logger:          logging.LoggerManager.GetLoggerByPattern("ResourceQueueThread"),
	}

	resourceQueueThread.loopThread = NewLoopThread(func() {
		resourceQueueThread.run()
	})
	return resourceQueueThread
}

func (this *ResourceQueueThread[T]) Start() {
	err := this.loopThread.Start()
	if err != nil {
		this.logger.Error(constants.STRING_ERROR, err)
		return
	}
}
func (this *ResourceQueueThread[T]) run() {
	select {
	case resource := <-this.resourceChannel:
		fmt.Println(fmt.Sprintf("%+v", resource))
	default:
		if this.needToStop {
			this.realStop()
		}
	}
}

func (this *ResourceQueueThread[T]) Stop() {
	this.needToStop = true
}
func (this *ResourceQueueThread[T]) StopNow() {
	this.realStop()
}
func (this *ResourceQueueThread[T]) realStop() {
	err := this.loopThread.Stop()
	if err != nil {
		return
	}
}
func (this *ResourceQueueThread[T]) AddResource(resource T) {
	this.resourceChannel <- resource
}
