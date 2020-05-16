// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/paper_repository.go

// Package usecase is a generated GoMock package.
package usecase

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.comm/wakatakeru/user-auth-jwt-api/domain"
	reflect "reflect"
)

// MockPaperRepository is a mock of PaperRepository interface.
type MockPaperRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaperRepositoryMockRecorder
}

// MockPaperRepositoryMockRecorder is the mock recorder for MockPaperRepository.
type MockPaperRepositoryMockRecorder struct {
	mock *MockPaperRepository
}

// NewMockPaperRepository creates a new mock instance.
func NewMockPaperRepository(ctrl *gomock.Controller) *MockPaperRepository {
	mock := &MockPaperRepository{ctrl: ctrl}
	mock.recorder = &MockPaperRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaperRepository) EXPECT() *MockPaperRepositoryMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockPaperRepository) Store(arg0 domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockPaperRepositoryMockRecorder) Store(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockPaperRepository)(nil).Store), arg0)
}

// Update mocks base method.
func (m *MockPaperRepository) Update(arg0 domain.User) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPaperRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPaperRepository)(nil).Update), arg0)
}

// FindByName mocks base method.
func (m *MockPaperRepository) FindByName(arg0 string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockPaperRepositoryMockRecorder) FindByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockPaperRepository)(nil).FindByName), arg0)
}

// FindByID mocks base method.
func (m *MockPaperRepository) FindByID(arg0 int) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockPaperRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPaperRepository)(nil).FindByID), arg0)
}
