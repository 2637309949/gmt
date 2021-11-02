// Package log provides a log interface
package logger

import (
	"context"
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

func Log(ctx context.Context, level logrus.Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(ctx context.Context, level logrus.Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}

func Out() io.Writer {
	return DefaultLogger.Options().Out
}

func Info(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.InfoLevel, args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.InfoLevel, template, args...)
}

func Trace(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.TraceLevel, args...)
}

func Tracef(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.TraceLevel, template, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.DebugLevel, args...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.DebugLevel, template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.WarnLevel, args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.WarnLevel, template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.ErrorLevel, args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.ErrorLevel, template, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	Logf(ctx, logrus.FatalLevel, template, args...)
	os.Exit(1)
}
