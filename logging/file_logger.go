package logging

import (
	"bufio"
	"fmt"
	"github.com/oneliang/util-golang/common"
	"github.com/oneliang/util-golang/constants"
	"github.com/oneliang/util-golang/file"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type ruleEnum struct {
	MINUTE *rule
	HOUR   *rule
	DAY    *rule
}

type rule struct {
	interval            int64
	directoryNameFormat string
	filenameFormat      string
}

var (
	Rule = initRule()

	defaultFileLoggerConfig = &FileLoggerConfig{
		Rule:       Rule.DAY,
		RemainDays: 7,
	}
	filePermission fs.FileMode = 0755
)

func initRule() *ruleEnum {
	rules := &ruleEnum{}
	rules.DAY = &rule{
		interval:            constants.TIME_MILLISECONDS_OF_DAY,
		directoryNameFormat: constants.TIME_LAYOUT_YEAR_MONTH_DAY,
		filenameFormat:      constants.TIME_LAYOUT_YEAR + constants.SYMBOL_UNDERLINE + constants.TIME_LAYOUT_MONTH + constants.SYMBOL_UNDERLINE + constants.TIME_LAYOUT_DAY,
	}
	rules.HOUR = &rule{
		interval:            constants.TIME_MILLISECONDS_OF_HOUR,
		directoryNameFormat: constants.TIME_LAYOUT_YEAR_MONTH_DAY,
		filenameFormat:      rules.DAY.filenameFormat + constants.SYMBOL_UNDERLINE + constants.TIME_LAYOUT_HOUR,
	}
	rules.MINUTE = &rule{
		interval:            constants.TIME_MILLISECONDS_OF_MINUTE,
		directoryNameFormat: constants.TIME_LAYOUT_YEAR_MONTH_DAY,
		filenameFormat:      rules.HOUR.filenameFormat + constants.SYMBOL_UNDERLINE + constants.TIME_LAYOUT_MINUTE,
	}
	return rules
}

type FileLogger struct {
	*AbstractLogger
	directoryAbsolutePath string
	filename              string
	rule                  *rule
	retainDays            uint
	currentBeginTime      int64
	logLock               sync.Mutex
	currentLogFile        *os.File
	currentLogFileWriter  *bufio.Writer
}

type FileLoggerConfig struct {
	Rule       *rule
	RemainDays uint
}

func NewFileLogger(level *level, directory string, filename string, fileLoggerConfig *FileLoggerConfig) *FileLogger {
	var config = fileLoggerConfig
	if config == nil {
		config = defaultFileLoggerConfig
	}
	directoryAbsolutePath, _ := filepath.Abs(directory)
	fileLogger := &FileLogger{
		directoryAbsolutePath: directoryAbsolutePath,
		filename:              filename,
		rule:                  config.Rule,
		retainDays:            config.RemainDays,
	}
	fileLogger.AbstractLogger = &AbstractLogger{
		Level: level,
		LogFunction: func(levelName string, message string, err error, args ...any) {
			_ = fileLogger.log(levelName, true, message, err, args...)
		},
	}
	fileLogger.init()
	return fileLogger
}
func (this *FileLogger) init() {
	this.currentBeginTime = common.GetZeroTime(time.Now().UnixMilli(), this.rule.interval)
	fullFilename, err := this.newFile(this.directoryAbsolutePath, this.currentBeginTime, this.filename, this.rule)
	if err != nil {
		panic(err)
	}
	this.currentLogFile, err = os.OpenFile(fullFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, filePermission)
	if err != nil {
		panic(err)
		return
	}
	this.currentLogFileWriter = bufio.NewWriter(this.currentLogFile)
}

func (this *FileLogger) newFile(directoryAbsolutePath string, currentBeginTime int64, filename string, rule *rule) (string, error) {
	beginDate := time.UnixMilli(currentBeginTime)
	subDirectoryName := beginDate.Format(rule.directoryNameFormat)
	filenamePrefix := beginDate.Format(rule.filenameFormat)
	subDirectoryAbsolutePath := filepath.Join(directoryAbsolutePath, subDirectoryName)
	err := os.MkdirAll(subDirectoryAbsolutePath, 0755)
	if err != nil {
		panic(err)
	}
	fullFilename := filepath.Join(subDirectoryAbsolutePath, filenamePrefix+constants.SYMBOL_UNDERLINE+filename)
	err = os.WriteFile(fullFilename, nil, 0755)
	if err != nil {
		return constants.STRING_BLANK, err
	}
	return fullFilename, nil
}

func (this *FileLogger) log(levelName string, printGoroutineId bool, message string, err error, args ...any) error {
	logContent := GenerateLogContent(levelName, printGoroutineId, message, err, args...) + constants.STRING_CRLF
	currentTime := time.Now().UnixMilli()
	timeInterval := currentTime - this.currentBeginTime
	if timeInterval >= this.rule.interval {
		this.logLock.Lock()
		defer this.logLock.Unlock()

		currentTime = time.Now().UnixMilli()
		timeInterval = currentTime - this.currentBeginTime
		//double check, current day may be change, day internal is the same when first in, but second time is not the same
		if timeInterval >= this.rule.interval {
			// zip log file before destroy
			this.Destroy()
			//compress current log file with zip
			err = this.zipCurrentLogFileAndRemoveIt(this.directoryAbsolutePath, this.currentBeginTime, this.filename, this.rule)
			if err != nil {
				fmt.Println(err)
				return err
			}
			//update current begin time
			this.currentBeginTime += this.rule.interval
			//delete expire file
			this.deleteExpireFile(this.directoryAbsolutePath, this.currentBeginTime, this.rule)
			//set to new file output stream
			fullFilename, err := this.newFile(this.directoryAbsolutePath, this.currentBeginTime, this.filename, this.rule)
			if err != nil {
				fmt.Println(fmt.Sprintf(":%s %v", fullFilename, err))
				return err
			}
			//reset
			this.currentLogFile, err = os.OpenFile(fullFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, filePermission)
			if err != nil {
				fmt.Println(err)
				return err
			}
			this.currentLogFileWriter = bufio.NewWriter(this.currentLogFile)
			fmt.Println(fmt.Sprintf("log, fullFilename:%s, currentLogFile:%s", fullFilename, this.currentLogFile))
		}
	}
	writeLogContent(this.currentLogFileWriter, logContent)
	return nil
}

func writeLogContent(fileWriter *bufio.Writer, logContent string) {
	if fileWriter == nil {
		return
	}
	_, err := fileWriter.WriteString(logContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fileWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (this *FileLogger) zipCurrentLogFileAndRemoveIt(directoryAbsolutePath string, currentBeginTime int64, filename string, rule *rule) error {
	beginDate := time.UnixMilli(currentBeginTime)
	subDirectoryName := beginDate.Format(rule.directoryNameFormat)
	filenamePrefix := beginDate.Format(rule.filenameFormat)
	subDirectoryAbsolutePath := filepath.Join(directoryAbsolutePath, subDirectoryName)
	err := os.MkdirAll(subDirectoryAbsolutePath, 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	oldLogFilename := filepath.Join(subDirectoryAbsolutePath, filenamePrefix+constants.SYMBOL_UNDERLINE+filename)
	zipFilename := filepath.Join(subDirectoryAbsolutePath, filenamePrefix+constants.SYMBOL_UNDERLINE+filename+constants.SYMBOL_DOT+constants.FILE_ZIP)
	err = file.Zip(zipFilename, oldLogFilename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.Remove(oldLogFilename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this *FileLogger) deleteExpireFile(directoryAbsolutePath string, currentBeginTime int64, rule *rule) {
	if this.retainDays >= 0 {
		for i := 30; i > int(this.retainDays); i-- {
			subDirectoryName := time.UnixMilli(common.GetDayZeroTimePrevious(currentBeginTime, i)).Format(rule.directoryNameFormat)
			subDirectoryAbsolutePath := filepath.Join(directoryAbsolutePath, subDirectoryName)
			err := os.RemoveAll(subDirectoryAbsolutePath)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	} else {
		//retain days < 0
	}
}

func (this *FileLogger) Destroy() {
	this.destroyCurrentFileOutputStream()
}

func (this *FileLogger) destroyCurrentFileOutputStream() {
	err := this.currentLogFileWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = this.currentLogFile.Close() }()
	defer this.currentLogFileWriter.Reset(nil)
}
