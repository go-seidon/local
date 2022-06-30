package logging

type Logger interface {
	Info(format string, args ...interface{}) error
	Debug(format string, args ...interface{}) error
	Error(format string, args ...interface{}) error
}

type LogOption struct {
	AppName    string
	AppVersion string

	DebuggingEnabled bool
}

type Option func(*LogOption)

func WithAppContext(name, version string) Option {
	return func(lo *LogOption) {
		lo.AppName = name
		lo.AppVersion = version
	}
}

func EnableDebugging() Option {
	return func(lo *LogOption) {
		lo.DebuggingEnabled = true
	}
}
