package logging

import (
	"github.com/oneliang/util-golang/constants"
	"regexp"
	"strings"
)

// LoggerManager global logger manager
var LoggerManager = loggerManager{
	loggerNameMap:        make(map[string]Logger),
	loggerPatternNameMap: make(map[string]Logger),
}

type loggerManager struct {
	loggerNameMap        map[string]Logger
	loggerPatternNameMap map[string]Logger
}

func (this *loggerManager) GetLogger(name string) Logger {
	logger, ok := this.loggerNameMap[name]
	if ok {
		return logger
	} else {
		return DEFAULT_LOGGER
	}
}

func (this *loggerManager) GetLoggerByPattern(name string) Logger {
	var logger Logger = nil
	for pattern, patternLogger := range this.loggerPatternNameMap {
		ok, err := regexp.MatchString(pattern, name)
		if err != nil {
			continue
		}
		if ok {
			logger = patternLogger
			break
		}
	}
	if logger == nil {
		logger = DEFAULT_LOGGER
	}
	return logger
}

func (this *loggerManager) RegisterLogger(name string, logger Logger) {
	this.loggerNameMap[name] = logger
}

func (this *loggerManager) RegisterLoggerByPattern(pattern string, logger Logger) {
	var fixPattern = pattern
	if pattern == constants.SYMBOL_WILDCARD {
		fixPattern = strings.Replace(pattern, constants.SYMBOL_WILDCARD, "[\\S|\\s]*", -1)
	}
	this.loggerPatternNameMap[fixPattern] = logger
}

func (this *loggerManager) UnRegisterLogger(name string) {
	delete(this.loggerNameMap, name)
}
