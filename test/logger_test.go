package test

import (
	"github.com/oneliang/util-golang/logging"
	"testing"
)

func TestLogger(t *testing.T) {
	logging.LoggerManager.RegisterLoggerByPattern("*", logging.DEFAULT_LOGGER)
	logger := logging.LoggerManager.GetLoggerByPattern("aaa")
	logger.Verbose(":%s", "aaa")

	baseLogger := logging.NewBaseLogger(logging.Level.DEBUG)
	baseLogger.Info("%s", "a")
}
