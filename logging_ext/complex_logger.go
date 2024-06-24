package loggingExt

import (
	"github.com/oneliang/util-golang/concurrent"
	"github.com/oneliang/util-golang/logging"
)

type ComplexLogger struct {
	*logging.AbstractLogger
	loggerList     []*logging.AbstractLogger
	async          bool
	logQueueThread *concurrent.ResourceQueueThread[*logMessage]
}

type logMessage struct {
	levelName string
	message   string
	err       error
}

func NewComplexLogger(level *logging.Level, loggerList []*logging.AbstractLogger, async bool) *ComplexLogger {
	complexLogger := &ComplexLogger{
		AbstractLogger: &logging.AbstractLogger{
			Level: level,
		},
		loggerList: loggerList,
		async:      async,
	}
	complexLogger.logQueueThread = concurrent.NewResourceQueueThread[*logMessage](func(resource *logMessage) {
		complexLogger.realLog(resource.levelName, resource.message, resource.err)
	}, nil)
	complexLogger.LogFunction = func(levelName string, message string, err error) {
		complexLogger.log(levelName, message, err)
	}
	return complexLogger
}

func (this *ComplexLogger) realLog(levelName string, message string, err error) {
	for _, logger := range this.loggerList {
		logger.LogFunction(levelName, message, err)
	}
}

func (this *ComplexLogger) log(levelName string, message string, err error) {
	if this.async {
		this.logQueueThread.AddResource(&logMessage{
			levelName: levelName,
			message:   message,
			err:       err,
		})
	} else {
		this.realLog(levelName, message, err)
	}
}

func (this *ComplexLogger) Destroy() {
	for _, logger := range this.loggerList {
		logger.Destroy()
	}
	this.AbstractLogger.Destroy()
	this.logQueueThread.Stop()
}
