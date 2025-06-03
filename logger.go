package eventloghook

type logger interface {
	Close() error
	Error(eid uint32, msg string) error
	Info(eid uint32, msg string) error
	Warning(eid uint32, msg string) error
}
