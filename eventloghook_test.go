//go:build windows
// +build windows

package eventloghook

import (
	"testing"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

func TestDebugLogAndPrint(t *testing.T) {
	elog := debug.New("ServiceName")
	defer elog.Close()

	testTheInterface(t, elog)
}

func TestEventLogAndPrint(t *testing.T) {
	elog, err := eventlog.Open("ServiceName")
	if err != nil {
		t.Errorf("Unexpected error %q while opening eventlog", err)
	}

	testTheInterface(t, elog)
}

func testTheInterface(t *testing.T, l logger) {
	log := logrus.New()
	log.Hooks.Add(NewHook(l))

	for _, level := range logrus.AllLevels {
		if len(log.Hooks[level]) != 1 {
			t.Errorf("Log was not added. The length of log.Hooks[%v]: %v", level, len(log.Hooks[level]))
		}
	}

	log.Infof("Congratulations (%T)!", l)
}
