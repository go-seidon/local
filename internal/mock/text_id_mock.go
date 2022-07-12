// Code generated by MockGen. DO NOT EDIT.
// Source: internal/text/id.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIdentifier is a mock of Identifier interface.
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier.
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance.
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// GenerateId mocks base method.
func (m *MockIdentifier) GenerateId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateId indicates an expected call of GenerateId.
func (mr *MockIdentifierMockRecorder) GenerateId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateId", reflect.TypeOf((*MockIdentifier)(nil).GenerateId))
}
