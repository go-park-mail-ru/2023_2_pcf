package mock_entities

import (
	"AdHub/internal/pkg/entities"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockCsrfUseCaseInterface is a mock of CsrfUseCaseInterface interface.
type MockCsrfUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCsrfUseCaseInterfaceMockRecorder
}

// MockCsrfUseCaseInterfaceMockRecorder is the mock recorder for MockCsrfUseCaseInterface.
type MockCsrfUseCaseInterfaceMockRecorder struct {
	mock *MockCsrfUseCaseInterface
}

// NewMockCsrfUseCaseInterface creates a new mock instance.
func NewMockCsrfUseCaseInterface(ctrl *gomock.Controller) *MockCsrfUseCaseInterface {
	mock := &MockCsrfUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockCsrfUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCsrfUseCaseInterface) EXPECT() *MockCsrfUseCaseInterfaceMockRecorder {
	return m.recorder
}

// CsrfCreate mocks base method.
func (m *MockCsrfUseCaseInterface) CsrfCreate(userId int) (*entities.Csrf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CsrfCreate", userId)
	ret0, _ := ret[0].(*entities.Csrf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CsrfCreate indicates an expected call of CsrfCreate.
func (mr *MockCsrfUseCaseInterfaceMockRecorder) CsrfCreate(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CsrfCreate", reflect.TypeOf((*MockCsrfUseCaseInterface)(nil).CsrfCreate), userId)
}

// CsrfRemove mocks base method.
func (m *MockCsrfUseCaseInterface) CsrfRemove(sr *entities.Csrf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CsrfRemove", sr)
	ret0, _ := ret[0].(error)
	return ret0
}

// CsrfRemove indicates an expected call of CsrfRemove.
func (mr *MockCsrfUseCaseInterfaceMockRecorder) CsrfRemove(sr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CsrfRemove", reflect.TypeOf((*MockCsrfUseCaseInterface)(nil).CsrfRemove), sr)
}

// GetByUserId mocks base method.
func (m *MockCsrfUseCaseInterface) GetByUserId(id int) (*entities.Csrf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", id)
	ret0, _ := ret[0].(*entities.Csrf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockCsrfUseCaseInterfaceMockRecorder) GetByUserId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockCsrfUseCaseInterface)(nil).GetByUserId), id)
}
