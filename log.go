package restgo

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
}
