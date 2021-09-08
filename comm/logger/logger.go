package logger

import (
	"io"
	"os"

	"github.com/micro/go-micro/v2/logger"
)

var (
	// Default logger
	l logger.Logger = logger.DefaultLogger
)

func Out() io.Writer {
	return l.Options().Out
}

func Info(args ...interface{}) {
	l.Log(logger.InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	l.Logf(logger.InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	l.Log(logger.TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	l.Logf(logger.TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	l.Log(logger.DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	l.Logf(logger.DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	l.Log(logger.WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	l.Logf(logger.WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	l.Log(logger.ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	l.Logf(logger.ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	l.Log(logger.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	l.Logf(logger.FatalLevel, template, args...)
	os.Exit(1)
}
