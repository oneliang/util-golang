package concurrent

import (
	"errors"
	"fmt"
	"sync"
)

const (
	state_inactive int8 = 0
	state_start    int8 = 1
	state_stop     int8 = 2
)

type LoopThread struct {
	operationLock sync.Mutex
	runnable      Runnable
	state         int8
	stopChannel   chan bool
}
type Runnable func()

func NewLoopThread(runnable Runnable) *LoopThread {
	return &LoopThread{
		runnable:    runnable,
		state:       state_inactive,
		stopChannel: make(chan bool, 1),
	}
}

func (this *LoopThread) Start() error {
	if err := this.checkStopped(); err != nil {
		return err
	}
	defer this.operationLock.Unlock()
	this.operationLock.Lock()
	if this.state == state_inactive {
		this.state = state_start
		go this.loopRun(this.stopChannel)
	} else {
		return errors.New(fmt.Sprintf("[%v] is running", this))
	}
	return nil
}

// loopRun is private method, only execute one times
func (this *LoopThread) loopRun(stopChannel chan bool) {
	if this.state == state_start {
		for {
			select {
			case stop, ok := <-stopChannel:
				if stop {
					return
				}
				if !ok {
					return
				}
			default:
				this.runnable()
			}
		}
	}
}

func (this *LoopThread) Stop() error {
	if err := this.checkStopped(); err != nil {
		return err
	}
	defer this.operationLock.Unlock()
	this.operationLock.Lock()
	this.stopChannel <- true
	this.state = state_stop
	close(this.stopChannel)
	return nil
}

func (this *LoopThread) checkStopped() error {
	if this.state == state_stop {
		return errors.New(fmt.Sprintf("[%v] is stopped", this))
	}
	return nil
}
