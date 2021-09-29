// Package log provides a log interface
package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is a generic logging interface
type Logger interface {
	// Init initialises options
	Init(options ...Option) error
	// The Logger options
	Options() Options
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger
	// Log writes a log entry
	Log(level logrus.Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level logrus.Level, format string, v ...interface{})
	// String returns the name of logger
	String() string
}

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level logrus.Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level logrus.Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}

func Out() io.Writer {
	return DefaultLogger.Options().Out
}

func Info(args ...interface{}) {
	Log(logrus.InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	Logf(logrus.InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	Log(logrus.TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	Logf(logrus.TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	Log(logrus.DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	Logf(logrus.DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	Log(logrus.WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	Logf(logrus.WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	Log(logrus.ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	Logf(logrus.ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	Log(logrus.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	Logf(logrus.FatalLevel, template, args...)
	os.Exit(1)
}
