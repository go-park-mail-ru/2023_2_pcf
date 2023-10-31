// Code generated by MockGen. DO NOT EDIT.
// Source: logger.go

// Package mock_logger is a generated GoMock package.
package mock_logger

import (
	http "net/http"
	reflect "reflect"
	time "time"

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

// Error mocks base method.
func (m *MockLogger) Error(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", message)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), message)
}

// Fatal mocks base method.
func (m *MockLogger) Fatal(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Fatal", message)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerMockRecorder) Fatal(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLogger)(nil).Fatal), message)
}

// Info mocks base method.
func (m *MockLogger) Info(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", message)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), message)
}

// MW mocks base method.
func (m *MockLogger) MW(message string, r *http.Request, duration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MW", message, r, duration)
}

// MW indicates an expected call of MW.
func (mr *MockLoggerMockRecorder) MW(message, r, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MW", reflect.TypeOf((*MockLogger)(nil).MW), message, r, duration)
}

// SetLogLevel mocks base method.
func (m *MockLogger) SetLogLevel(level string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogLevel", level)
}

// SetLogLevel indicates an expected call of SetLogLevel.
func (mr *MockLoggerMockRecorder) SetLogLevel(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockLogger)(nil).SetLogLevel), level)
}
