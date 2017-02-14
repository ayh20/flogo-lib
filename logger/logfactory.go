package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"strings"
)

var loggerMap = make(map[string]interface{})

type DefaultLoggerFactory struct {
}

func init() {
	RegisterLoggerFactory(&DefaultLoggerFactory{})
}

type DefaultLogger struct {
	loggerName string
	loggerImpl *logrus.Logger
}

type LogFormatter struct {
	loggerName string
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logEntry := fmt.Sprintf("%s %-6s [%s] - %s\n", entry.Time.Format("2006-01-02 15:04:05.000000"), getLevel(entry.Level), f.loggerName, strings.TrimPrefix(strings.TrimSuffix(entry.Message, "]"), "["))
	return []byte(logEntry), nil
}

func getLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.PanicLevel:
		return "PANIC"
	case logrus.FatalLevel:
		return "FATAL"
	}

	return "UNKNOWN"
}

// Debug logs message at Debug level.
func (logger *DefaultLogger) Debug(args ...interface{}) {
	logger.loggerImpl.Debug(args)
}

// DebugEnabled checks if Debug level is enabled.
func (logger *DefaultLogger) DebugEnabled() bool {
	return logger.loggerImpl.Level >= logrus.DebugLevel
}

// Info logs message at Info level.
func (logger *DefaultLogger) Info(args ...interface{}) {
	logger.loggerImpl.Info(args)
}

// InfoEnabled checks if Info level is enabled.
func (logger *DefaultLogger) InfoEnabled() bool {
	return logger.loggerImpl.Level >= logrus.InfoLevel
}

// Warn logs message at Warning level.
func (logger *DefaultLogger) Warn(args ...interface{}) {
	logger.loggerImpl.Warn(args)
}

// WarnEnabled checks if Warning level is enabled.
func (logger *DefaultLogger) WarnEnabled() bool {
	return logger.loggerImpl.Level >= logrus.WarnLevel
}

// Error logs message at Error level.
func (logger *DefaultLogger) Error(args ...interface{}) {
	logger.loggerImpl.Error(args)
}

// ErrorEnabled checks if Error level is enabled.
func (logger *DefaultLogger) ErrorEnabled() bool {
	return logger.loggerImpl.Level >= logrus.ErrorLevel
}

//SetLog Level
func (logger *DefaultLogger) SetLogLevel(logLevel Level) {
	switch logLevel {
	case Debug:
		logger.loggerImpl.Level = logrus.DebugLevel
	case Info:
		logger.loggerImpl.Level = logrus.InfoLevel
	case Error:
		logger.loggerImpl.Level = logrus.ErrorLevel
	case Warn:
		logger.loggerImpl.Level = logrus.WarnLevel
	default:
		logger.loggerImpl.Level = logrus.ErrorLevel
	}
}

func (logfactory *DefaultLoggerFactory) GetLogger(name string) (Logger, error) {
	logger := loggerMap[name]
	if logger == nil {
		logImpl := logrus.New()
		logImpl.Formatter = &LogFormatter{
			loggerName: name,
		}
		logger = &DefaultLogger{
			loggerName: name,
			loggerImpl: logImpl,
		}
		loggerMap[name] = logger
	}
	return logger.(Logger), nil
}
