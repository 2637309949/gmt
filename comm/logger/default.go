package logger

import (
	"comm/conf"
	"comm/graylog"
	"context"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	dlog "github.com/micro/go-micro/v2/debug/log"
	"github.com/micro/go-micro/v2/logger"
	"github.com/sirupsen/logrus"
)

// Default logger
var DefaultLogger Logger

func init() {
	graylogAddress := conf.Load("comm", "graylog_address").String()
	logger.Infof("graylog_address:%v", graylogAddress)

	options := Options{
		Level:           logrus.InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		CallerSkipCount: 2,
		Context:         context.Background(),
	}
	l := &defaultLogger{opts: options, hooks: make(LevelHooks)}
	l.AddHook(graylog.NewAsyncGraylogHook(graylogAddress, map[string]interface{}{}))
	DefaultLogger = NewHelper(l)
}

// Internal type for storing the hooks on a logger instance.
type LevelHooks map[logrus.Level][]Hook

// Add a hook to an instance of logger. This is called with
// `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.
func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// Fire all the hooks for the passed level. Used by `entry.log` to fire
// appropriate hooks for a log entry.
func (hooks LevelHooks) Fire(level logrus.Level, t string, m interface{}, fields map[string]interface{}) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(level, t, m, fields); err != nil {
			return err
		}
	}

	return nil
}

type graylogEntry struct {
	l      logrus.Level
	t      string
	m      interface{}
	fields map[string]interface{}
}

type defaultLogger struct {
	sync.RWMutex
	opts  Options
	hooks LevelHooks
}

// Init(opts...) should only overwrite provided options
func (l *defaultLogger) Init(opts ...Option) error {
	for _, o := range opts {
		o(&l.opts)
	}
	return nil
}

func (l *defaultLogger) String() string {
	return "default"
}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.Lock()
	l.opts.Fields = copyFields(fields)
	l.Unlock()
	return l
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// logCallerfilePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func logCallerfilePath(loggingFilePath string) string {
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func (l *defaultLogger) Log(level logrus.Level, v ...interface{}) {
	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = logCallerfilePath(file)
		fields["line"] = line
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprint(v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	dlog.DefaultLog.Write(rec)

	t := rec.Timestamp.Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %v\n", t, metadata, rec.Message)
	l.hooks.Fire(level, t, rec.Message, fields)
}

func (l *defaultLogger) Logf(level logrus.Level, format string, v ...interface{}) {
	if level < l.opts.Level {
		return
	}

	l.RLock()
	fields := copyFields(l.opts.Fields)
	l.RUnlock()

	fields["level"] = level.String()

	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["file"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	rec := dlog.Record{
		Timestamp: time.Now(),
		Message:   fmt.Sprintf(format, v...),
		Metadata:  make(map[string]string, len(fields)),
	}

	keys := make([]string, 0, len(fields))
	for k, v := range fields {
		keys = append(keys, k)
		rec.Metadata[k] = fmt.Sprintf("%v", v)
	}

	sort.Strings(keys)
	metadata := ""

	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	dlog.DefaultLog.Write(rec)

	t := rec.Timestamp.Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %v\n", t, metadata, rec.Message)
	l.hooks.Fire(level, t, rec.Message, fields)
}

func (n *defaultLogger) Options() Options {
	n.RLock()
	opts := n.opts
	opts.Fields = copyFields(n.opts.Fields)
	n.RUnlock()
	return opts
}

// AddHook adds a hook to the logger hooks.
func (n *defaultLogger) AddHook(hook Hook) {
	n.Lock()
	defer n.Unlock()
	n.hooks.Add(hook)
}
