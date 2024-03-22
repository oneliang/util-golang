package logging

import (
	"fmt"
	"github.com/oneliang/util-golang/base"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/constants"
	"time"
)

type LogFunction func(levelName string, message string, err error)
type AbstractLogger struct {
	Level       *Level
	LogFunction LogFunction
}

// Verbose .
func (this *AbstractLogger) Verbose(message string, args ...any) {
	this.logByLevel(LevelConstants.VERBOSE, message, nil, args...)
}

// Debug .
func (this *AbstractLogger) Debug(message string, args ...any) {
	this.logByLevel(LevelConstants.DEBUG, message, nil, args...)
}

// Info .
func (this *AbstractLogger) Info(message string, args ...any) {
	this.logByLevel(LevelConstants.INFO, message, nil, args...)
}

// Warning .
func (this *AbstractLogger) Warning(message string, args ...any) {
	this.logByLevel(LevelConstants.WARNING, message, nil, args...)
}

// Error .
func (this *AbstractLogger) Error(message string, err error, args ...any) {
	this.logByLevel(LevelConstants.ERROR, message, err, args...)
}

// Fatal .
func (this *AbstractLogger) Fatal(message string, args ...any) {
	this.logByLevel(LevelConstants.FATAL, message, nil, args...)
}

// LogByLevel .
func (this *AbstractLogger) logByLevel(level *Level, message string, err error, args ...any) {
	if level.ordinal >= this.Level.ordinal {
		this.log(level.name, message, err, args...)
	}
}

// Log .
func (this *AbstractLogger) log(levelName string, message string, err error, args ...any) {
	logContent := GenerateLogContent(levelName, true, message, err, args)
	this.LogFunction(levelName, logContent, err)
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
