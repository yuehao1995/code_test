/**
 * Created by g7tianyi on 09/18/2018
 */

package logger

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/op/go-logging"
)

const (
	logLevelFile     = logging.DEBUG
	logLevelConsole  = logging.DEBUG
	logFormatFile    = "%{time:0102 15:04:05.999999} %{shortfile}:%{shortfunc} ▶ %{level:.4s} %{message}"
	logFormatConsole = "%{color}%{time:15:04:05.000} %{shortfile}:%{shortfunc} ▶ " +
		"%{level:.4s} %{id:03x} %{message}%{color:reset}"
	logBackupTimeFormat = "2006-01-02_15-04-05.000"
)

var logger *logging.Logger

var logName string
var logFileName string
var logFileWriter *os.File
var logFileLock sync.Mutex
var logRotateChan chan os.Signal

func init() {
	initLogger("scaring")
}

func createLogFileWriter(logFile string) (*os.File, error) {
	const flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	const mask = 0666

	var err error

	dir, _ := path.Split(logFile)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	var file *os.File
	if _, err = os.Stat(logFile); os.IsNotExist(err) {
		file, err = os.Create(logFile)
	} else {
		file, err = os.OpenFile(logFile, flag, mask)
	}
	if err != nil {
		return nil, err
	}

	return file, nil
}

// backupName creates a new filename from the given name, inserting a timestamp
// between the filename and the extension, using the local time if requested
// (otherwise UTC).
func backupName(name string, local bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := time.Now()
	if !local {
		t = t.UTC()
	}

	timestamp := t.Format(logBackupTimeFormat)
	return filepath.Join(dir, fmt.Sprintf("%s.%s%s", prefix, timestamp, ext))
}

func closeLogFile() error {
	if logFileWriter == nil {
		return nil
	}
	logFileLock.Lock()
	defer logFileLock.Unlock()
	return logFileWriter.Close()
}

func rotateLogFile() error {
	if err := closeLogFile(); err != nil {
		return err
	}

	if logFileName != "" {
		logFileLock.Lock()
		defer logFileLock.Unlock()

		name := logFileName
		next := backupName(logFileName, true)
		if err := os.Rename(name, next); err != nil {
			return fmt.Errorf("can't rename log file: %s", err)
		}
		logger.Infof("rotate the log from [%s] to [%s]", name, next)

		initLogger(name)
	}
	return nil
}

func initLogger(loggerName string) *logging.Logger {

	logger = nil

	logName = loggerName
	logFileName = filepath.Join("/", "opt", "eechains", "scaring", "log", loggerName+".log")

	var backEnds []logging.Backend
	logger = logging.MustGetLogger(logName)

	stdLogBackground := logging.NewLogBackend(os.Stderr, "", 0)
	stdLogFormatter := logging.NewBackendFormatter(stdLogBackground,
		logging.MustStringFormatter(logFormatConsole))
	stdLogLeveled := logging.AddModuleLevel(stdLogFormatter)
	stdLogLeveled.SetLevel(logLevelConsole, "")

	backEnds = append(backEnds, stdLogLeveled)

	if logName != "" {

		var err error

		// if the log file name has been specified. then try to open it
		if logFileName != "" {
			logFileWriter, err = createLogFileWriter(logFileName)
			if err != nil {
				logger.Warningf("%s", err.Error())
			}
		}

		// if cannot open the specified log file, then open the default log file
		if logFileWriter == nil {
			logFileName = fmt.Sprintf("logs/%s.log", loggerName)
			logFileWriter, err = createLogFileWriter(logFileName)
			if err != nil {
				logger.Warningf("%s", err.Error())
			}
		}

		// if the log file can be opened, then added it into backend
		// otherwise, just log an warning message
		if logFileWriter == nil {
			logger.Warningf("%s", err.Error())
		} else {
			logger.Infof("opened the log file [%s]", logFileName)

			fileLogBackground := logging.NewLogBackend(logFileWriter, "", 0)
			fileLogFormatter := logging.NewBackendFormatter(fileLogBackground,
				logging.MustStringFormatter(logFormatFile))
			fileLogLeveled := logging.AddModuleLevel(fileLogFormatter)
			fileLogLeveled.SetLevel(logLevelFile, "")
			backEnds = append(backEnds, fileLogLeveled)

			// try to trap the HUP signal for log rotate
			if logRotateChan == nil {
				logRotateChan = make(chan os.Signal, 1)
				signal.Notify(logRotateChan, syscall.SIGHUP)
				go func() {
					for {
						<-logRotateChan
						logger.Infof("received the HUP signal.")
						rotateLogFile()
					}
				}()
			}
		}
	}

	// Could not directly use logger instance, Must use follow wrapped function
	logger.ExtraCalldepth = 1

	logging.SetBackend(backEnds...)

	return logger
}

func InitLogger() *logging.Logger {
	return logger
}
