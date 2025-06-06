//go:build windows
// +build windows

package eventloghook

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// EventLogHook to send logs via windows log.
type EventLogHook struct {
	upstream Logger
}

// NewHook creates and returns a new EventLogHook wrapped around anything that implements the Logger interface
// for example:
// * golang.org/x/sys/windows/svc/eventlog.Log
// * golang.org/x/sys/windows/svc/debug.Log
func NewHook(Logger Logger) *EventLogHook {
	return &EventLogHook{upstream: Logger}
}

func (hook *EventLogHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.upstream.Error(3, line)
	case logrus.FatalLevel:
		return hook.upstream.Error(2, line)
	case logrus.ErrorLevel:
		return hook.upstream.Error(1, line)
	case logrus.WarnLevel:
		return hook.upstream.Warning(1, line)
	case logrus.InfoLevel:
		return hook.upstream.Info(3, line)
	case logrus.DebugLevel:
		return hook.upstream.Info(2, line)
	case logrus.TraceLevel:
		return hook.upstream.Info(1, line)
	default:
		return nil
	}
}

func (hook *EventLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
