// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dynatrace-oss/dynatrace-monitoring-as-code/pkg/api (interfaces: Api)

// Package mock_api is a generated GoMock package.
package api

import (
	environment "github.com/dynatrace-oss/dynatrace-monitoring-as-code/pkg/environment"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApi is a mock of Api interface
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// GetApiPath mocks base method
func (m *MockApi) GetApiPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetApiPath indicates an expected call of GetApiPath
func (mr *MockApiMockRecorder) GetApiPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiPath", reflect.TypeOf((*MockApi)(nil).GetApiPath))
}

// GetId mocks base method
func (m *MockApi) GetId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId
func (mr *MockApiMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockApi)(nil).GetId))
}

// GetUrl mocks base method
func (m *MockApi) GetUrl(arg0 environment.Environment) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrl", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrl indicates an expected call of GetUrl
func (mr *MockApiMockRecorder) GetUrl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockApi)(nil).GetUrl), arg0)
}

// GetUrlFromEnvironmentUrl mocks base method
func (m *MockApi) GetUrlFromEnvironmentUrl(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrlFromEnvironmentUrl", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUrlFromEnvironmentUrl indicates an expected call of GetUrlFromEnvironmentUrl
func (mr *MockApiMockRecorder) GetUrlFromEnvironmentUrl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrlFromEnvironmentUrl", reflect.TypeOf((*MockApi)(nil).GetUrlFromEnvironmentUrl), arg0)
}