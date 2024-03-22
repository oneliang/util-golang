package test

import (
	"github.com/oneliang/util-golang/logging"
	loggingExt "github.com/oneliang/util-golang/logging_ext"
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	fileLogger := loggingExt.NewFileLogger(logging.LevelConstants.VERBOSE, "log", "default.log", &loggingExt.FileLoggerConfig{
		Rule:       loggingExt.Rule.MINUTE,
		RemainDays: 1,
	})
	for i := 0; i < 10; i++ {
		fileLogger.Info("loop index:%d", i)
		time.Sleep(10 * time.Second)
	}
}
