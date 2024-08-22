package goroutine

import (
	"fmt"
	"github.com/oneliang/util-golang/concurrent"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
)

type Task func(params ...any) error

type taskWrapper struct {
	task   *Task
	params []any
}
type Pool struct {
	resourceQueueThread *concurrent.ResourceQueueThread[taskWrapper]
	taskQueue           chan taskWrapper
	logger              logging.Logger
}

func NewPool(goroutineSize int) *Pool {
	pool := &Pool{
		taskQueue: make(chan taskWrapper),
		logger:    logging.LoggerManager.GetLoggerByPattern("goroutine.Pool"),
	}
	pool.resourceQueueThread = concurrent.NewResourceQueueThread[taskWrapper](func(resource taskWrapper) {
		pool.logger.Info("pool.taskQueue <- resource:%+v", resource)
		pool.taskQueue <- resource
	}, func() {
		fmt.Println("finished")
		close(pool.taskQueue)
	})

	for i := 0; i < goroutineSize; i++ {
		go func() {
			for {
				select {
				case taskItem, ok := <-pool.taskQueue:
					if !ok {
						continue
					}
					pool.logger.Info("go task hashcode:%+v", taskItem)
					if taskItem.task != nil {
						err := (*taskItem.task)(taskItem.params...)
						if err != nil {
							pool.logger.Error(constants.STRING_ERROR, err)
						}
					} else {
						pool.logger.Error("go task is nil", nil)
					}
				default:
				}
			}
		}()
	}

	return pool
}

func (this *Pool) AddTask(task Task, params ...any) {
	this.resourceQueueThread.AddResource(taskWrapper{
		task:   &task,
		params: params,
	})
}

func (this *Pool) Start() {
	this.resourceQueueThread.Start()
}

func (this *Pool) Stop() {
	this.resourceQueueThread.Stop()
}
