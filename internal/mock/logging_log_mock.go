// Code generated by MockGen. DO NOT EDIT.
// Source: internal/logging/log.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	logging "github.com/go-seidon/local/internal/logging"
	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Debugln mocks base method.
func (m *MockLogger) Debugln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugln", varargs...)
}

// Debugln indicates an expected call of Debugln.
func (mr *MockLoggerMockRecorder) Debugln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugln", reflect.TypeOf((*MockLogger)(nil).Debugln), msg...)
}

// Error mocks base method.
func (m *MockLogger) Error(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Errorln mocks base method.
func (m *MockLogger) Errorln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorln", varargs...)
}

// Errorln indicates an expected call of Errorln.
func (mr *MockLoggerMockRecorder) Errorln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorln", reflect.TypeOf((*MockLogger)(nil).Errorln), msg...)
}

// Info mocks base method.
func (m *MockLogger) Info(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Infoln mocks base method.
func (m *MockLogger) Infoln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infoln", varargs...)
}

// Infoln indicates an expected call of Infoln.
func (mr *MockLoggerMockRecorder) Infoln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infoln", reflect.TypeOf((*MockLogger)(nil).Infoln), msg...)
}

// Warn mocks base method.
func (m *MockLogger) Warn(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), args...)
}

// Warnf mocks base method.
func (m *MockLogger) Warnf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockLoggerMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}

// Warnln mocks base method.
func (m *MockLogger) Warnln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnln", varargs...)
}

// Warnln indicates an expected call of Warnln.
func (mr *MockLoggerMockRecorder) Warnln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnln", reflect.TypeOf((*MockLogger)(nil).Warnln), msg...)
}

// WithFields mocks base method.
func (m *MockLogger) WithFields(fs map[string]interface{}) logging.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithFields", fs)
	ret0, _ := ret[0].(logging.Logger)
	return ret0
}

// WithFields indicates an expected call of WithFields.
func (mr *MockLoggerMockRecorder) WithFields(fs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithFields", reflect.TypeOf((*MockLogger)(nil).WithFields), fs)
}

// MockSimpleLog is a mock of SimpleLog interface.
type MockSimpleLog struct {
	ctrl     *gomock.Controller
	recorder *MockSimpleLogMockRecorder
}

// MockSimpleLogMockRecorder is the mock recorder for MockSimpleLog.
type MockSimpleLogMockRecorder struct {
	mock *MockSimpleLog
}

// NewMockSimpleLog creates a new mock instance.
func NewMockSimpleLog(ctrl *gomock.Controller) *MockSimpleLog {
	mock := &MockSimpleLog{ctrl: ctrl}
	mock.recorder = &MockSimpleLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimpleLog) EXPECT() *MockSimpleLogMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockSimpleLog) Debug(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockSimpleLogMockRecorder) Debug(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockSimpleLog)(nil).Debug), args...)
}

// Error mocks base method.
func (m *MockSimpleLog) Error(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockSimpleLogMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSimpleLog)(nil).Error), args...)
}

// Info mocks base method.
func (m *MockSimpleLog) Info(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockSimpleLogMockRecorder) Info(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockSimpleLog)(nil).Info), args...)
}

// Warn mocks base method.
func (m *MockSimpleLog) Warn(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockSimpleLogMockRecorder) Warn(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockSimpleLog)(nil).Warn), args...)
}

// MockFormatedLog is a mock of FormatedLog interface.
type MockFormatedLog struct {
	ctrl     *gomock.Controller
	recorder *MockFormatedLogMockRecorder
}

// MockFormatedLogMockRecorder is the mock recorder for MockFormatedLog.
type MockFormatedLogMockRecorder struct {
	mock *MockFormatedLog
}

// NewMockFormatedLog creates a new mock instance.
func NewMockFormatedLog(ctrl *gomock.Controller) *MockFormatedLog {
	mock := &MockFormatedLog{ctrl: ctrl}
	mock.recorder = &MockFormatedLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFormatedLog) EXPECT() *MockFormatedLogMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockFormatedLog) Debugf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockFormatedLogMockRecorder) Debugf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockFormatedLog)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockFormatedLog) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockFormatedLogMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockFormatedLog)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *MockFormatedLog) Infof(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockFormatedLogMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockFormatedLog)(nil).Infof), varargs...)
}

// Warnf mocks base method.
func (m *MockFormatedLog) Warnf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockFormatedLogMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockFormatedLog)(nil).Warnf), varargs...)
}

// MockLineLog is a mock of LineLog interface.
type MockLineLog struct {
	ctrl     *gomock.Controller
	recorder *MockLineLogMockRecorder
}

// MockLineLogMockRecorder is the mock recorder for MockLineLog.
type MockLineLogMockRecorder struct {
	mock *MockLineLog
}

// NewMockLineLog creates a new mock instance.
func NewMockLineLog(ctrl *gomock.Controller) *MockLineLog {
	mock := &MockLineLog{ctrl: ctrl}
	mock.recorder = &MockLineLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLineLog) EXPECT() *MockLineLogMockRecorder {
	return m.recorder
}

// Debugln mocks base method.
func (m *MockLineLog) Debugln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugln", varargs...)
}

// Debugln indicates an expected call of Debugln.
func (mr *MockLineLogMockRecorder) Debugln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugln", reflect.TypeOf((*MockLineLog)(nil).Debugln), msg...)
}

// Errorln mocks base method.
func (m *MockLineLog) Errorln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorln", varargs...)
}

// Errorln indicates an expected call of Errorln.
func (mr *MockLineLogMockRecorder) Errorln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorln", reflect.TypeOf((*MockLineLog)(nil).Errorln), msg...)
}

// Infoln mocks base method.
func (m *MockLineLog) Infoln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infoln", varargs...)
}

// Infoln indicates an expected call of Infoln.
func (mr *MockLineLogMockRecorder) Infoln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infoln", reflect.TypeOf((*MockLineLog)(nil).Infoln), msg...)
}

// Warnln mocks base method.
func (m *MockLineLog) Warnln(msg ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnln", varargs...)
}

// Warnln indicates an expected call of Warnln.
func (mr *MockLineLogMockRecorder) Warnln(msg ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnln", reflect.TypeOf((*MockLineLog)(nil).Warnln), msg...)
}
