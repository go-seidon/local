package logging

type Logger interface {
	Info(format string, args ...interface{}) error
	Debug(format string, args ...interface{}) error
	Error(format string, args ...interface{}) error
}
