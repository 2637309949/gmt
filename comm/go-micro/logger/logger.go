package logger

import (
	"github.com/micro/go-micro/v2/logger"
)

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Tracef(template string, args ...interface{}) {
	logger.Tracef(template, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Returns true if the given level is at or lower the current logger level
func V(lvl logger.Level, log logger.Logger) bool {
	return logger.V(lvl, log)
}
