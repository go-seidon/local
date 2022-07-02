package logging

import (
	"os"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
	client *logrus.Logger
}

func (l *logrusLog) Info(args ...interface{}) {
	l.client.Info(args...)
}

func (l *logrusLog) Debug(args ...interface{}) {
	l.client.Debug(args...)
}

func (l *logrusLog) Error(args ...interface{}) {
	l.client.Error(args...)
}

func (l *logrusLog) Warn(args ...interface{}) {
	l.client.Warn(args...)
}

func (l *logrusLog) Infof(format string, args ...interface{}) {
	l.client.Infof(format, args...)
}

func (l *logrusLog) Debugf(format string, args ...interface{}) {
	l.client.Debugf(format, args...)
}

func (l *logrusLog) Errorf(format string, args ...interface{}) {
	l.client.Errorf(format, args...)
}

func (l *logrusLog) Warnf(format string, args ...interface{}) {
	l.client.Warnf(format, args...)
}

func (l *logrusLog) Infoln(args ...interface{}) {
	l.client.Infoln(args...)
}

func (l *logrusLog) Debugln(args ...interface{}) {
	l.client.Debugln(args...)
}

func (l *logrusLog) Errorln(args ...interface{}) {
	l.client.Errorln(args...)
}

func (l *logrusLog) Warnln(args ...interface{}) {
	l.client.Warnln(args...)
}

func (l *logrusLog) WithFields(fs map[string]interface{}) Logger {
	l.client = l.client.WithFields(fs).Logger
	return l
}

func NewLogrusLog(opts ...Option) *logrusLog {
	option := LogOption{}
	for _, opt := range opts {
		opt(&option)
	}

	c := logrus.New()
	c.SetFormatter(stackdriver.NewFormatter(
		stackdriver.WithService(option.AppName),
		stackdriver.WithVersion(option.AppVersion),
	))
	c.SetOutput(os.Stdout)

	if option.DebuggingEnabled {
		c.SetLevel(logrus.DebugLevel)
	}

	l := &logrusLog{
		client: c,
	}
	return l
}
