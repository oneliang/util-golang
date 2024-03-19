package logging

import (
	"fmt"
	"github.com/oneliang/util-golang/base"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/constants"
	"time"
)

type LogFunction func(levelName string, message string, err error, args ...any)
type AbstractLogger struct {
	Level       *level
	LogFunction LogFunction
}

// Verbose .
func (this *AbstractLogger) Verbose(message string, args ...any) {
	this.logByLevel(Level.VERBOSE, message, nil, args...)
}

// Debug .
func (this *AbstractLogger) Debug(message string, args ...any) {
	this.logByLevel(Level.DEBUG, message, nil, args...)
}

// Info .
func (this *AbstractLogger) Info(message string, args ...any) {
	this.logByLevel(Level.INFO, message, nil, args...)
}

// Warning .
func (this *AbstractLogger) Warning(message string, args ...any) {
	this.logByLevel(Level.WARNING, message, nil, args...)
}

// Error .
func (this *AbstractLogger) Error(message string, err error, args ...any) {
	this.logByLevel(Level.ERROR, message, err, args...)
}

// Fatal .
func (this *AbstractLogger) Fatal(message string, args ...any) {
	this.logByLevel(Level.FATAL, message, nil, args...)
}

// LogByLevel .
func (this *AbstractLogger) logByLevel(level *level, message string, err error, args ...any) {
	if level.ordinal >= this.Level.ordinal {
		this.log(level.name, message, err, args...)
	}
}

// Log .
func (this *AbstractLogger) log(levelName string, message string, err error, args ...any) {
	this.LogFunction(levelName, message, err, args...)
}

// Destroy .
func (this *AbstractLogger) Destroy() {
}

func GenerateLogContent(levelName string, printGoroutineId bool, message string, err error, args ...any) string {
	var errorString = ""
	if err != nil {
		errorString = "\n" + fmt.Sprintf("%+v", base.WithStack(err))
	}
	timeString := time.Now().Format(constants.TIME_LAYOUT_YEAR_MONTH_DAY_HOUR_MINUTE_SECOND_MILLISECOND)
	if printGoroutineId {
		goroutineId := common.GetGoroutineId()
		return fmt.Sprintf("[%s][%s][GO_ID:%d][%s]%s", timeString, levelName, goroutineId, fmt.Sprintf(message, args...), errorString)
	} else {
		return fmt.Sprintf("[%s][%s][%s]%s", timeString, levelName, fmt.Sprintf(message, args...), errorString)
	}
}
