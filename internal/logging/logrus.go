package logging

import (
	"fmt"
	"os"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
	client LogClient
}

func (l *logrusLog) Info(format string, args ...interface{}) error {
	l.client.Infof(format, args...)
	return nil
}

func (l *logrusLog) Debug(format string, args ...interface{}) error {
	l.client.Debugf(format, args...)
	return nil
}

func (l *logrusLog) Error(format string, args ...interface{}) error {
	l.client.Errorf(format, args...)
	return nil
}

type LogClient interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type NewLogrusLogOption struct {
	AppName    string
	AppVersion string
}

func NewLogrusLog(opt *NewLogrusLogOption) (*logrusLog, error) {
	if opt == nil {
		return nil, fmt.Errorf("logrus option is invalid")
	}

	c := logrus.New()
	c.SetFormatter(stackdriver.NewFormatter(
		stackdriver.WithService(opt.AppName),
		stackdriver.WithVersion(opt.AppVersion),
	))
	c.SetOutput(os.Stdout)
	c.SetLevel(logrus.DebugLevel)

	l := &logrusLog{
		client: c,
	}
	return l, nil
}
