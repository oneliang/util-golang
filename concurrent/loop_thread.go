package concurrent

import (
	"errors"
	"fmt"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/logging"
	"sync"
)

const (
	state_inactive int8 = 0
	state_start    int8 = 1
	state_stop     int8 = 2
)

type Runnable func()

type LoopThread struct {
	operationLock *sync.Mutex
	runnable      Runnable
	state         int8
	stopChannel   chan bool
	logger        logging.Logger
}

func NewLoopThread(runnable Runnable) *LoopThread {
	return &LoopThread{
		operationLock: &sync.Mutex{},
		runnable:      runnable,
		state:         state_inactive,
		stopChannel:   make(chan bool, 1),
		logger:        logging.LoggerManager.GetLoggerByPattern("concurrent.LoopThread"),
	}
}

func (this *LoopThread) Start() error {
	if err := this.checkStopped(); err != nil {
		this.logger.Error(constants.STRING_ERROR, err)
		return err
	}
	defer this.operationLock.Unlock()
	this.operationLock.Lock()
	if this.state == state_inactive {
		this.state = state_start
		go this.loopRun(this.stopChannel)
	} else {
		err := errors.New(fmt.Sprintf("[%+v] is running", this))
		this.logger.Error(constants.STRING_ERROR, err)
		return err
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
		err := errors.New(fmt.Sprintf("[%v] is stopped", this))
		this.logger.Error(constants.STRING_ERROR, err)
		return err
	}
	return nil
}
