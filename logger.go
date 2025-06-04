//go:build windows
// +build windows

package eventloghook

// Logger interface to cover the two common windows loggers
// IE:
// * golang.org/x/sys/windows/svc/eventlog.Log
// * golang.org/x/sys/windows/svc/debug.Log
type Logger interface {
	Error(eid uint32, msg string) error
	Info(eid uint32, msg string) error
	Warning(eid uint32, msg string) error
}
