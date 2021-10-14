package graylog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const StackTraceKey = "_stacktrace"

// Set graylog.BufSize = <value> _before_ calling NewGraylogHook
// Once the buffer is full, logging will start blocking, waiting for slots to
// be available in the queue.
var BufSize uint = 8192

// GraylogHook to send logs to a logging service compatible with the Graylog API and the GELF format.
type GraylogHook struct {
	Extra       map[string]interface{}
	Host        string
	Level       logrus.Level
	gelfLogger  *Writer
	buf         chan graylogEntry
	wg          sync.WaitGroup
	mu          sync.RWMutex
	synchronous bool
	blacklist   map[string]bool
}

// Graylog needs file and line params
type graylogEntry struct {
	time     string
	level    int
	file     string
	function string
	line     int
	message  string
	fields   map[string]interface{}
}

// NewGraylogHook creates a hook to be added to an instance of logger.
func NewGraylogHook(addr string, extra map[string]interface{}) *GraylogHook {
	g, _ := NewWriter(addr)
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	hook := &GraylogHook{
		Host:        host,
		Extra:       extra,
		Level:       logrus.DebugLevel,
		gelfLogger:  g,
		synchronous: true,
		buf:         make(chan graylogEntry, BufSize),
	}
	return hook
}

// NewAsyncGraylogHook creates a hook to be added to an instance of logger.
// The hook created will be asynchronous, and it's the responsibility of the user to call the Flush method
// before exiting to empty the log queue.
func NewAsyncGraylogHook(addr string, extra map[string]interface{}) *GraylogHook {
	g, _ := NewWriter(addr)
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	hook := &GraylogHook{
		Host:       host,
		Extra:      extra,
		Level:      logrus.DebugLevel,
		gelfLogger: g,
		buf:        make(chan graylogEntry, BufSize),
	}
	go hook.fire() // Log in background
	return hook
}

// Fire is called when a log event is fired.
// We assume the entry will be altered by another hook,
// otherwise we might logging something wrong to Graylog
func (hook *GraylogHook) Fire(l logrus.Level, t string, m interface{}, fields map[string]interface{}) error {
	hook.mu.RLock() // Claim the mutex as a RLock - allowing multiple go routines to log simultaneously
	defer hook.mu.RUnlock()

	var file = fields["file"]
	var line = fields["line"]
	var function = fields["function"]

	newData := make(map[string]interface{})
	for k, v := range fields {
		newData[k] = v
	}

	le, _ := strconv.Atoi(fmt.Sprintf("%v", line))
	gEntry := graylogEntry{
		time:     t,
		level:    int(l),
		line:     le,
		file:     fmt.Sprintf("%v", file),
		function: fmt.Sprintf("%v", function),
		message:  fmt.Sprintf("%v", m),
		fields:   newData,
	}

	if hook.synchronous {
		hook.sendEntry(gEntry)
	} else {
		hook.wg.Add(1)
		hook.buf <- gEntry
	}

	return nil
}

// Flush waits for the log queue to be empty.
// This func is meant to be used when the hook was created with NewAsyncGraylogHook.
func (hook *GraylogHook) Flush() {
	hook.mu.Lock() // claim the mutex as a Lock - we want exclusive access to it
	defer hook.mu.Unlock()

	hook.wg.Wait()
}

// fire will loop on the 'buf' channel, and write entries to graylog
func (hook *GraylogHook) fire() {
	for {
		entry := <-hook.buf // receive new entry on channel
		hook.sendEntry(entry)
		hook.wg.Done()
	}
}

// sendEntry sends an entry to graylog synchronously
func (hook *GraylogHook) sendEntry(entry graylogEntry) {
	if hook.gelfLogger == nil {
		fmt.Println("Can't connect to Graylog")
		return
	}
	w := hook.gelfLogger

	// remove trailing and leading whitespace
	p := bytes.TrimSpace([]byte(entry.message))

	// If there are newlines in the message, use the first line
	// for the short message and set the full message to the
	// original input.  If the input has no newlines, stick the
	// whole thing in Short.
	short := p
	full := []byte("")
	if i := bytes.IndexRune(p, '\n'); i > 0 {
		short = p[:i]
		full = p
	}

	// Don't modify entry.Data directly, as the entry will used after this hook was fired
	extra := map[string]interface{}{}
	// Merge extra fields
	for k, v := range hook.Extra {
		k = fmt.Sprintf("_%s", k) // "[...] every field you send and prefix with a _ (underscore) will be treated as an additional field."
		extra[k] = v
	}

	extra["_file"] = entry.file
	extra["_line"] = entry.line
	extra["_function"] = entry.function

	for k, v := range entry.fields {
		if !hook.blacklist[k] {
			extraK := fmt.Sprintf("_%s", k) // "[...] every field you send and prefix with a _ (underscore) will be treated as an additional field."
			if k == logrus.ErrorKey {
				asError, isError := v.(error)
				_, isMarshaler := v.(json.Marshaler)
				if isError && !isMarshaler {
					extra[extraK] = newMarshalableError(asError)
				} else {
					extra[extraK] = v
				}
				if stackTrace := extractStackTrace(asError); stackTrace != nil {
					extra[StackTraceKey] = fmt.Sprintf("%+v", stackTrace)
				}
			} else {
				extra[extraK] = v
			}
		}
	}

	m := Message{
		Version:  "1.1",
		Host:     hook.Host,
		Short:    string(short),
		Full:     string(full),
		TimeUnix: float64(time.Now().UnixNano()/1000000) / 1000.,
		Level:    entry.level,
		File:     entry.file,
		Line:     entry.line,
		Extra:    extra,
	}

	if err := w.WriteMessage(&m); err != nil {
		fmt.Println(err)
	}
}

// Levels returns the available logging levels.
func (hook *GraylogHook) Levels() []logrus.Level {
	levels := []logrus.Level{}
	for _, level := range logrus.AllLevels {
		if level <= hook.Level {
			levels = append(levels, level)
		}
	}
	return levels
}

// Blacklist create a blacklist map to filter some message keys.
// This useful when you want your application to log extra fields locally
// but don't want graylog to store them.
func (hook *GraylogHook) Blacklist(b []string) {
	hook.blacklist = make(map[string]bool)
	for _, elem := range b {
		hook.blacklist[elem] = true
	}
}

// SetWriter sets the hook Gelf Writer
func (hook *GraylogHook) SetWriter(w *Writer) error {
	if w == nil {
		return errors.New("writer can't be nil")
	}
	hook.gelfLogger = w
	return nil
}

// Writer returns the logger Gelf Writer
func (hook *GraylogHook) Writer() *Writer {
	return hook.gelfLogger
}
