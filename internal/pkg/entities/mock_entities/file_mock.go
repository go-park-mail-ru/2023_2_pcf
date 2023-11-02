// Code generated by MockGen. DO NOT EDIT.
// Source: file.go

// Package mock_entities is a generated GoMock package.
package mock_entities

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileRepoInterface is a mock of FileRepoInterface interface.
type MockFileRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFileRepoInterfaceMockRecorder
}

// MockFileRepoInterfaceMockRecorder is the mock recorder for MockFileRepoInterface.
type MockFileRepoInterfaceMockRecorder struct {
	mock *MockFileRepoInterface
}

// NewMockFileRepoInterface creates a new mock instance.
func NewMockFileRepoInterface(ctrl *gomock.Controller) *MockFileRepoInterface {
	mock := &MockFileRepoInterface{ctrl: ctrl}
	mock.recorder = &MockFileRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileRepoInterface) EXPECT() *MockFileRepoInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFileRepoInterface) Delete(fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFileRepoInterfaceMockRecorder) Delete(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFileRepoInterface)(nil).Delete), fileName)
}

// Get mocks base method.
func (m *MockFileRepoInterface) Get(fileName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", fileName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFileRepoInterfaceMockRecorder) Get(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFileRepoInterface)(nil).Get), fileName)
}

// Save mocks base method.
func (m *MockFileRepoInterface) Save(fileData []byte, originalName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", fileData, originalName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockFileRepoInterfaceMockRecorder) Save(fileData, originalName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileRepoInterface)(nil).Save), fileData, originalName)
}

// MockFileUseCaseInterface is a mock of FileUseCaseInterface interface.
type MockFileUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFileUseCaseInterfaceMockRecorder
}

// MockFileUseCaseInterfaceMockRecorder is the mock recorder for MockFileUseCaseInterface.
type MockFileUseCaseInterfaceMockRecorder struct {
	mock *MockFileUseCaseInterface
}

// NewMockFileUseCaseInterface creates a new mock instance.
func NewMockFileUseCaseInterface(ctrl *gomock.Controller) *MockFileUseCaseInterface {
	mock := &MockFileUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockFileUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileUseCaseInterface) EXPECT() *MockFileUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockFileUseCaseInterface) Delete(fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFileUseCaseInterfaceMockRecorder) Delete(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFileUseCaseInterface)(nil).Delete), fileName)
}

// Get mocks base method.
func (m *MockFileUseCaseInterface) Get(fileName string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", fileName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFileUseCaseInterfaceMockRecorder) Get(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFileUseCaseInterface)(nil).Get), fileName)
}

// Save mocks base method.
func (m *MockFileUseCaseInterface) Save(fileData []byte, originalName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", fileData, originalName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockFileUseCaseInterfaceMockRecorder) Save(fileData, originalName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockFileUseCaseInterface)(nil).Save), fileData, originalName)
}