package test

import (
	"github.com/oneliang/util-golang/logging"
	"testing"
)

func TestLogger(t *testing.T) {
	logging.LoggerManager.RegisterLoggerByPattern("*", logging.DEFAULT_LOGGER)
	logger := logging.LoggerManager.GetLoggerByPattern("logger_test")
	logger.Verbose("%s, %s", "aaa", "bbb")

	baseLogger := logging.NewBaseLogger(logging.LevelConstants.DEBUG)
	baseLogger.Info("%s, %s", "a", "b")
}
