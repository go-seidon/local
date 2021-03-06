// Code generated by MockGen. DO NOT EDIT.
// Source: internal/deleting/deleter.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	deleting "github.com/go-seidon/local/internal/deleting"
	gomock "github.com/golang/mock/gomock"
)

// MockDeleter is a mock of Deleter interface.
type MockDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDeleterMockRecorder
}

// MockDeleterMockRecorder is the mock recorder for MockDeleter.
type MockDeleterMockRecorder struct {
	mock *MockDeleter
}

// NewMockDeleter creates a new mock instance.
func NewMockDeleter(ctrl *gomock.Controller) *MockDeleter {
	mock := &MockDeleter{ctrl: ctrl}
	mock.recorder = &MockDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleter) EXPECT() *MockDeleterMockRecorder {
	return m.recorder
}

// DeleteFile mocks base method.
func (m *MockDeleter) DeleteFile(ctx context.Context, p deleting.DeleteFileParam) (*deleting.DeleteFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, p)
	ret0, _ := ret[0].(*deleting.DeleteFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockDeleterMockRecorder) DeleteFile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockDeleter)(nil).DeleteFile), ctx, p)
}
