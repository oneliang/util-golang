package goroutine

import (
	"fmt"
	"github.com/oneliang/util-golang/concurrent"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
)

type Task func() error
type Pool struct {
	resourceQueueThread *concurrent.ResourceQueueThread[Task]
	taskQueue           chan Task
	logger              logging.Logger
}

func NewPool(goroutineSize int) *Pool {
	pool := &Pool{
		taskQueue: make(chan Task),
		logger:    logging.LoggerManager.GetLoggerByPattern("goroutine.Pool"),
	}
	pool.resourceQueueThread = concurrent.NewResourceQueueThread[Task](func(resource Task) {
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
				case task, ok := <-pool.taskQueue:
					if !ok {
						continue
					}
					pool.logger.Info("go task hashcode:%+v", task)
					err := task()
					if err != nil {
						pool.logger.Error(constants.STRING_ERROR, err)
					}
				default:
				}
			}
		}()
	}

	return pool
}

func (this *Pool) AddTask(task Task) {
	this.resourceQueueThread.AddResource(task)
}

func (this *Pool) Start() {
	this.resourceQueueThread.Start()
}

func (this *Pool) Stop() {
	this.resourceQueueThread.Stop()
}
