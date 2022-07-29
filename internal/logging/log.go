package logging

import "context"

type Logger interface {
	SimpleLog
	FormatedLog
	LineLog
	CustomLog
}

type SimpleLog interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}

type FormatedLog interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

type LineLog interface {
	Infoln(msg ...interface{})
	Debugln(msg ...interface{})
	Errorln(msg ...interface{})
	Warnln(msg ...interface{})
}

type CustomLog interface {
	WithFields(fs map[string]interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger
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
