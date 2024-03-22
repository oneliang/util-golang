package logging

import (
	"fmt"
)

type BaseLogger struct {
	*AbstractLogger
}

// NewBaseLogger .
func NewBaseLogger(level *Level) *BaseLogger {
	return &BaseLogger{&AbstractLogger{
		Level: level,
		LogFunction: func(levelName string, message string, err error) {
			fmt.Println(message)
		},
	}}
}

var (
	DEFAULT_LOGGER = NewBaseLogger(LevelConstants.VERBOSE)
)
