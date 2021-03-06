// Code generated by MockGen. DO NOT EDIT.
// Source: sub.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAlgSub is a mock of AlgSub interface
type MockAlgSub struct {
	ctrl     *gomock.Controller
	recorder *MockAlgSubMockRecorder
}

// MockAlgSubMockRecorder is the mock recorder for MockAlgSub
type MockAlgSubMockRecorder struct {
	mock *MockAlgSub
}

// NewMockAlgSub creates a new mock instance
func NewMockAlgSub(ctrl *gomock.Controller) *MockAlgSub {
	mock := &MockAlgSub{ctrl: ctrl}
	mock.recorder = &MockAlgSubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlgSub) EXPECT() *MockAlgSubMockRecorder {
	return m.recorder
}

// Sub mocks base method
func (m *MockAlgSub) Sub(lhs, rhs int) int {
	ret := m.ctrl.Call(m, "Sub", lhs, rhs)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sub indicates an expected call of Sub
func (mr *MockAlgSubMockRecorder) Sub(lhs, rhs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockAlgSub)(nil).Sub), lhs, rhs)
}

// Sub1 mocks base method
func (m *MockAlgSub) Sub1(lhs, mid, rhs int) int {
	ret := m.ctrl.Call(m, "Sub1", lhs, mid, rhs)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sub1 indicates an expected call of Sub1
func (mr *MockAlgSubMockRecorder) Sub1(lhs, mid, rhs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub1", reflect.TypeOf((*MockAlgSub)(nil).Sub1), lhs, mid, rhs)
}
