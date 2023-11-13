// Code generated by MockGen. DO NOT EDIT.
// Source: systemd.go

// Package mock_driver is a generated GoMock package.
package mock_driver

import (
	context "context"
	io "io"
	reflect "reflect"

	driver "github.com/awslabs/aws-s3-csi-driver/pkg/driver"
	dbus "github.com/coreos/go-systemd/v22/dbus"
	gomock "github.com/golang/mock/gomock"
)

// MockSystemdConnection is a mock of SystemdConnection interface.
type MockSystemdConnection struct {
	ctrl     *gomock.Controller
	recorder *MockSystemdConnectionMockRecorder
}

// MockSystemdConnectionMockRecorder is the mock recorder for MockSystemdConnection.
type MockSystemdConnectionMockRecorder struct {
	mock *MockSystemdConnection
}

// NewMockSystemdConnection creates a new mock instance.
func NewMockSystemdConnection(ctrl *gomock.Controller) *MockSystemdConnection {
	mock := &MockSystemdConnection{ctrl: ctrl}
	mock.recorder = &MockSystemdConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemdConnection) EXPECT() *MockSystemdConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSystemdConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSystemdConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSystemdConnection)(nil).Close))
}

// ListUnitsContext mocks base method.
func (m *MockSystemdConnection) ListUnitsContext(ctx context.Context) ([]dbus.UnitStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnitsContext", ctx)
	ret0, _ := ret[0].([]dbus.UnitStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnitsContext indicates an expected call of ListUnitsContext.
func (mr *MockSystemdConnectionMockRecorder) ListUnitsContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnitsContext", reflect.TypeOf((*MockSystemdConnection)(nil).ListUnitsContext), ctx)
}

// ResetFailedUnitContext mocks base method.
func (m *MockSystemdConnection) ResetFailedUnitContext(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetFailedUnitContext", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetFailedUnitContext indicates an expected call of ResetFailedUnitContext.
func (mr *MockSystemdConnectionMockRecorder) ResetFailedUnitContext(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFailedUnitContext", reflect.TypeOf((*MockSystemdConnection)(nil).ResetFailedUnitContext), ctx, name)
}

// StartTransientUnitContext mocks base method.
func (m *MockSystemdConnection) StartTransientUnitContext(ctx context.Context, name, mode string, properties []dbus.Property, ch chan<- string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTransientUnitContext", ctx, name, mode, properties, ch)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTransientUnitContext indicates an expected call of StartTransientUnitContext.
func (mr *MockSystemdConnectionMockRecorder) StartTransientUnitContext(ctx, name, mode, properties, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTransientUnitContext", reflect.TypeOf((*MockSystemdConnection)(nil).StartTransientUnitContext), ctx, name, mode, properties, ch)
}

// MockSystemdConnector is a mock of SystemdConnector interface.
type MockSystemdConnector struct {
	ctrl     *gomock.Controller
	recorder *MockSystemdConnectorMockRecorder
}

// MockSystemdConnectorMockRecorder is the mock recorder for MockSystemdConnector.
type MockSystemdConnectorMockRecorder struct {
	mock *MockSystemdConnector
}

// NewMockSystemdConnector creates a new mock instance.
func NewMockSystemdConnector(ctrl *gomock.Controller) *MockSystemdConnector {
	mock := &MockSystemdConnector{ctrl: ctrl}
	mock.recorder = &MockSystemdConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemdConnector) EXPECT() *MockSystemdConnectorMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockSystemdConnector) Connect(ctx context.Context) (driver.SystemdConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(driver.SystemdConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockSystemdConnectorMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockSystemdConnector)(nil).Connect), ctx)
}

// MockPts is a mock of Pts interface.
type MockPts struct {
	ctrl     *gomock.Controller
	recorder *MockPtsMockRecorder
}

// MockPtsMockRecorder is the mock recorder for MockPts.
type MockPtsMockRecorder struct {
	mock *MockPts
}

// NewMockPts creates a new mock instance.
func NewMockPts(ctrl *gomock.Controller) *MockPts {
	mock := &MockPts{ctrl: ctrl}
	mock.recorder = &MockPtsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPts) EXPECT() *MockPtsMockRecorder {
	return m.recorder
}

// NewPts mocks base method.
func (m *MockPts) NewPts() (io.ReadCloser, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPts")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewPts indicates an expected call of NewPts.
func (mr *MockPtsMockRecorder) NewPts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPts", reflect.TypeOf((*MockPts)(nil).NewPts))
}