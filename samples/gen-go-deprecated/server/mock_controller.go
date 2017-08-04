// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

package server

import (
	context "context"
	models "github.com/Clever/wag/samples/gen-go-deprecated/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockController) EXPECT() *MockControllerMockRecorder {
	return _m.recorder
}

// Health mocks base method
func (_m *MockController) Health(ctx context.Context, i *models.HealthInput) error {
	ret := _m.ctrl.Call(_m, "Health", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Health indicates an expected call of Health
func (_mr *MockControllerMockRecorder) Health(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Health", reflect.TypeOf((*MockController)(nil).Health), arg0, arg1)
}
