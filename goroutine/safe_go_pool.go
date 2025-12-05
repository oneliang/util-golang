package goroutine

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"log"
)

// SafeGoPool 用于安全地并发执行任务，具备 panic 恢复 + 并发控制 + errgroup 管理能力
type SafeGoPool struct {
	errorGroup *errgroup.Group
	ctx        context.Context
	sem        *semaphore.Weighted
}

// NewSafeGoPool 创建一个带最大并发数限制的 SafeGoPool
func NewSafeGoPool(ctx context.Context, maxConcurrent int) *SafeGoPool {
	errorGroup, gctx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(int64(maxConcurrent))
	return &SafeGoPool{
		errorGroup: errorGroup,
		ctx:        gctx,
		sem:        sem,
	}
}

// Go 启动一个安全 goroutine，捕获 panic 并限制并发数
func (this *SafeGoPool) Go(fn func() error) {
	this.errorGroup.Go(func() (err error) {
		if err := this.sem.Acquire(this.ctx, 1); err != nil {
			return err
		}
		defer this.sem.Release(1)

		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGoPool] panic recovered: %v", r)
				err = fmt.Errorf("internal panic: %v", r)
			}
		}()

		return fn()
	})
}

// Wait 等待所有任务完成
func (this *SafeGoPool) Wait() error {
	return this.errorGroup.Wait()
}
