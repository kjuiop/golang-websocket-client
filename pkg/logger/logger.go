package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"

	"log"
	"os"
	"time"
)

const (
	fileField = "file"
)

type Logger struct {
	log *logrus.Logger
}

func LogInitialize(logLevel, logPath string) (*Logger, error) {

	l := new(Logger)
	l.log = logrus.New()

	lv := l.getLevel(logLevel)

	l.log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	l.log.SetLevel(lv)

	if len(logPath) > 0 {
		logger, err := l.setRollingLogFile(logPath)
		if err != nil {
			log.Printf(logPath+" : %s", err.Error())
			return nil, err
		}

		l.log.SetOutput(logger)
	} else {
		l.log.SetOutput(os.Stdout)
	}

	l.log.SetLevel(lv)

	return l, nil
}

func (l *Logger) Info(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.InfoLevel {
		prefix.Data[fileField] = l.fileInfo(2)
		prefix.Infof(format, args...)
	}
}

func (l *Logger) Trace(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.TraceLevel {
		prefix.Data[fileField] = l.fileInfo(2)
		prefix.Debugf(format, args...)
	}
}

func (l *Logger) Debug(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.DebugLevel {
		prefix.Data[fileField] = l.fileInfo(2)
		prefix.Debugf(format, args...)
	}
}

func (l *Logger) Warn(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.WarnLevel {
		prefix.Data[fileField] = l.fileInfo(2)
		prefix.Debugf(format, args...)
	}
}

func (l *Logger) Error(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.ErrorLevel {
		prefix.Data[fileField] = l.fileInfo(2)
		prefix.Errorf(format, args...)
	}
}

func (l *Logger) TestSocketConnectionProcessPrefix(start time.Time) *logrus.Entry {
	prefix := l.WithFields(logrus.Fields{})
	prefix.Data["start"] = start
	return prefix
}

func (l *Logger) InitPrefixData() *logrus.Entry {
	prefix := l.WithFields(logrus.Fields{})
	return prefix
}

func (l *Logger) getLevel(level string) (lv logrus.Level) {
	lv = logrus.InfoLevel
	switch strings.ToLower(level) {
	case "debug":
		lv = logrus.DebugLevel
	case "info":
		lv = logrus.InfoLevel
	case "warn":
		lv = logrus.WarnLevel
	case "error":
		lv = logrus.ErrorLevel
	default:
		logrus.Info("Unknown level string.")
	}
	return
}

// SetRollingLogFile periodically changes the log file.
func (l *Logger) setRollingLogFile(path string) (*rotatelogs.RotateLogs, error) {
	rotateLogger, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(path),
	)

	if err != nil {
		return nil, err
	}

	return rotateLogger, nil
}

func (l *Logger) fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return (*logrus.Entry)(l.log.WithFields(logrus.Fields(fields)))
}
