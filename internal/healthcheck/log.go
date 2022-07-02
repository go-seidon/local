package healthcheck

import (
	gLog "github.com/InVisionApp/go-logger"
	"github.com/go-seidon/local/internal/logging"
)

type GoHealthLog struct {
	Client logging.Logger
}

func (l *GoHealthLog) Info(args ...interface{}) {
	l.Client.Info(args...)
}

func (l *GoHealthLog) Debug(args ...interface{}) {
	l.Client.Debug(args...)
}

func (l *GoHealthLog) Error(args ...interface{}) {
	l.Client.Error(args...)
}

func (l *GoHealthLog) Warn(args ...interface{}) {
	l.Client.Warn(args...)
}

func (l *GoHealthLog) Infof(format string, args ...interface{}) {
	l.Client.Infof(format, args...)
}

func (l *GoHealthLog) Debugf(format string, args ...interface{}) {
	l.Client.Debugf(format, args...)
}

func (l *GoHealthLog) Errorf(format string, args ...interface{}) {
	l.Client.Errorf(format, args...)
}

func (l *GoHealthLog) Warnf(format string, args ...interface{}) {
	l.Client.Warnf(format, args...)
}

func (l *GoHealthLog) Infoln(args ...interface{}) {
	l.Client.Infoln(args...)
}

func (l *GoHealthLog) Debugln(args ...interface{}) {
	l.Client.Debugln(args...)
}

func (l *GoHealthLog) Errorln(args ...interface{}) {
	l.Client.Errorln(args...)
}

func (l *GoHealthLog) Warnln(args ...interface{}) {
	l.Client.Warnln(args...)
}

func (l *GoHealthLog) WithFields(fs gLog.Fields) gLog.Logger {
	l.Client = l.Client.WithFields(fs)
	return l
}
