package healthcheck

import (
	"fmt"

	gLog "github.com/InVisionApp/go-logger"
	"github.com/go-seidon/local/internal/logging"
)

type goHealthLog struct {
	client logging.Logger
}

func (l *goHealthLog) Info(args ...interface{}) {
	l.client.Info(args...)
}

func (l *goHealthLog) Debug(args ...interface{}) {
	l.client.Debug(args...)
}

func (l *goHealthLog) Error(args ...interface{}) {
	l.client.Error(args...)
}

func (l *goHealthLog) Warn(args ...interface{}) {
	l.client.Warn(args...)
}

func (l *goHealthLog) Infof(format string, args ...interface{}) {
	l.client.Infof(format, args...)
}

func (l *goHealthLog) Debugf(format string, args ...interface{}) {
	l.client.Debugf(format, args...)
}

func (l *goHealthLog) Errorf(format string, args ...interface{}) {
	l.client.Errorf(format, args...)
}

func (l *goHealthLog) Warnf(format string, args ...interface{}) {
	l.client.Warnf(format, args...)
}

func (l *goHealthLog) Infoln(args ...interface{}) {
	l.client.Infoln(args...)
}

func (l *goHealthLog) Debugln(args ...interface{}) {
	l.client.Debugln(args...)
}

func (l *goHealthLog) Errorln(args ...interface{}) {
	l.client.Errorln(args...)
}

func (l *goHealthLog) Warnln(args ...interface{}) {
	l.client.Warnln(args...)
}

func (l *goHealthLog) WithFields(fs gLog.Fields) gLog.Logger {
	l.client.WithFields(fs)
	return l
}

func NewGoHealthLog(logger logging.Logger) (*goHealthLog, error) {
	if logger == nil {
		return nil, fmt.Errorf("invalid logger")
	}

	l := &goHealthLog{
		client: logger,
	}
	return l, nil
}
