// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/oauth.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	repository "github.com/go-seidon/local/internal/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockOAuthRepository is a mock of OAuthRepository interface.
type MockOAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthRepositoryMockRecorder
}

// MockOAuthRepositoryMockRecorder is the mock recorder for MockOAuthRepository.
type MockOAuthRepositoryMockRecorder struct {
	mock *MockOAuthRepository
}

// NewMockOAuthRepository creates a new mock instance.
func NewMockOAuthRepository(ctrl *gomock.Controller) *MockOAuthRepository {
	mock := &MockOAuthRepository{ctrl: ctrl}
	mock.recorder = &MockOAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthRepository) EXPECT() *MockOAuthRepositoryMockRecorder {
	return m.recorder
}

// FindClient mocks base method.
func (m *MockOAuthRepository) FindClient(ctx context.Context, p repository.FindClientParam) (*repository.FindClientResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindClient", ctx, p)
	ret0, _ := ret[0].(*repository.FindClientResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindClient indicates an expected call of FindClient.
func (mr *MockOAuthRepositoryMockRecorder) FindClient(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClient", reflect.TypeOf((*MockOAuthRepository)(nil).FindClient), ctx, p)
}
