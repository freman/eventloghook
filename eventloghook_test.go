//go:build windows
// +build windows

package eventloghook

import (
	"testing"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/debug"
)

func TestDebugLogAndPrint(t *testing.T) {
	log := logrus.New()
	elog = debug.New("ServiceName")
	defer elog.Close()
	log.Hooks.Add(NewEventLogHook(elog))

	for _, level := range hook.Levels() {
		if len(log.Hooks[level]) != 1 {
			t.Errorf("EventLogHook was not added. The length of log.Hooks[%v]: %v", level, len(log.Hooks[level]))
		}
	}

	log.Info("Congratulations!")
}
