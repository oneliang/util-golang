package logging

import (
	"fmt"
)

type BaseLogger struct {
	*AbstractLogger
}

// NewBaseLogger .
func NewBaseLogger(level *level) *BaseLogger {
	return &BaseLogger{&AbstractLogger{
		Level: level,
		LogFunction: func(levelName string, message string, err error, args ...any) {
			fmt.Println(GenerateLogContent(levelName, true, message, err, args...))
		},
	}}
}

var (
	DEFAULT_LOGGER = NewBaseLogger(Level.VERBOSE)
)
