package logging

import (
	"fmt"
	"github.com/oneliang/util-golang/common"
)

type BaseLogger struct {
	AbstractLogger
}

// NewBaseLogger .
func NewBaseLogger(level *level) *BaseLogger {
	return &BaseLogger{AbstractLogger{
		Level: level,
		LogFunction: func(levelName string, message string, err error, args ...any) {
			fmt.Println(GenerateLogString(levelName, true, message, err, args...))
		},
	}}
}
func GenerateLogString(levelName string, printGoroutineId bool, message string, err error, args ...any) string {
	var errorString = ""
	if err != nil {
		errorString = "\n" + fmt.Sprintf("%+v", common.WithStack(err))
	}
	if printGoroutineId {
		goroutineId := common.GetGoroutineId()
		return fmt.Sprintf("[GOID:%d][%s][%s]%s", goroutineId, levelName, fmt.Sprintf(message, args...), errorString)
	} else {
		return fmt.Sprintf("[%s][%s]%s", levelName, fmt.Sprintf(message, args...), errorString)
	}
}

var (
	DEFAULT_LOGGER = NewBaseLogger(Level.VERBOSE)
)
