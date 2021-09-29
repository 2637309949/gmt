package logger

import "github.com/sirupsen/logrus"

// A hook to be fired when logging on the logging levels returned from
// `Levels()` on your implementation of the interface. Note that this is not
// fired in a goroutine or a channel with workers, you should handle such
// functionality yourself if your call is non-blocking and you don't wish for
// the logging calls for levels returned from `Levels()` to block.
type Hook interface {
	Levels() []logrus.Level
	Fire(l logrus.Level, t string, m interface{}, fields map[string]interface{}) error
}
