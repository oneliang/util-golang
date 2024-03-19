package test

import (
	"github.com/oneliang/util-golang/logging"
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	fileLogger := logging.NewFileLogger(logging.Level.VERBOSE, "log", "default.log", &logging.FileLoggerConfig{
		Rule:       logging.Rule.MINUTE,
		RemainDays: 1,
	})
	for i := 0; i < 10; i++ {
		fileLogger.Info("loop index:%d", i)
		time.Sleep(10 * time.Second)
	}
}
