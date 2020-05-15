// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/database/sql_handler.go

// Package database is a generated GoMock package.
package database

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSqlHandler is a mock of SqlHandler interface.
type MockSqlHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSqlHandlerMockRecorder
}

// MockSqlHandlerMockRecorder is the mock recorder for MockSqlHandler.
type MockSqlHandlerMockRecorder struct {
	mock *MockSqlHandler
}

// NewMockSqlHandler creates a new mock instance.
func NewMockSqlHandler(ctrl *gomock.Controller) *MockSqlHandler {
	mock := &MockSqlHandler{ctrl: ctrl}
	mock.recorder = &MockSqlHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqlHandler) EXPECT() *MockSqlHandlerMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockSqlHandler) Execute(arg0 string, arg1 ...interface{}) (SqlResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(SqlResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockSqlHandlerMockRecorder) Execute(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSqlHandler)(nil).Execute), varargs...)
}

// Query mocks base method.
func (m *MockSqlHandler) Query(arg0 string, arg1 ...interface{}) (SqlRow, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(SqlRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockSqlHandlerMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSqlHandler)(nil).Query), varargs...)
}

// MockSqlResult is a mock of SqlResult interface.
type MockSqlResult struct {
	ctrl     *gomock.Controller
	recorder *MockSqlResultMockRecorder
}

// MockSqlResultMockRecorder is the mock recorder for MockSqlResult.
type MockSqlResultMockRecorder struct {
	mock *MockSqlResult
}

// NewMockSqlResult creates a new mock instance.
func NewMockSqlResult(ctrl *gomock.Controller) *MockSqlResult {
	mock := &MockSqlResult{ctrl: ctrl}
	mock.recorder = &MockSqlResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqlResult) EXPECT() *MockSqlResultMockRecorder {
	return m.recorder
}

// LastInsertedId mocks base method.
func (m *MockSqlResult) LastInsertedId() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastInsertedId")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastInsertedId indicates an expected call of LastInsertedId.
func (mr *MockSqlResultMockRecorder) LastInsertedId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastInsertedId", reflect.TypeOf((*MockSqlResult)(nil).LastInsertedId))
}

// RowsAffected mocks base method.
func (m *MockSqlResult) RowsAffected() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RowsAffected indicates an expected call of RowsAffected.
func (mr *MockSqlResultMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockSqlResult)(nil).RowsAffected))
}

// MockSqlRow is a mock of SqlRow interface.
type MockSqlRow struct {
	ctrl     *gomock.Controller
	recorder *MockSqlRowMockRecorder
}

// MockSqlRowMockRecorder is the mock recorder for MockSqlRow.
type MockSqlRowMockRecorder struct {
	mock *MockSqlRow
}

// NewMockSqlRow creates a new mock instance.
func NewMockSqlRow(ctrl *gomock.Controller) *MockSqlRow {
	mock := &MockSqlRow{ctrl: ctrl}
	mock.recorder = &MockSqlRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqlRow) EXPECT() *MockSqlRowMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockSqlRow) Scan(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockSqlRowMockRecorder) Scan(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockSqlRow)(nil).Scan), arg0...)
}

// Next mocks base method.
func (m *MockSqlRow) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockSqlRowMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSqlRow)(nil).Next))
}

// Close mocks base method.
func (m *MockSqlRow) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSqlRowMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSqlRow)(nil).Close))
}
