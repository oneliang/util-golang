package concurrent

import (
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
)

type ResourceQueueThread[T any] struct {
	loopThread        *LoopThread
	needToStop        bool
	resourceChannel   chan T
	resourceProcessor func(resource T)
	realStopCallback  func()
	logger            logging.Logger
}

func NewResourceQueueThread[T any](resourceProcessor func(resource T), realStopCallback func()) *ResourceQueueThread[T] {
	resourceQueueThread := &ResourceQueueThread[T]{
		needToStop:        false,
		resourceChannel:   make(chan T),
		resourceProcessor: resourceProcessor,
		realStopCallback:  realStopCallback,
		logger:            logging.LoggerManager.GetLoggerByPattern("concurrent.ResourceQueueThread"),
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
		this.resourceProcessor(resource)
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
	} else {
		this.needToStop = false
		close(this.resourceChannel)
		if this.realStopCallback != nil {
			this.realStopCallback()
		}
	}
}
func (this *ResourceQueueThread[T]) AddResource(resource T) {
	this.resourceChannel <- resource
}
