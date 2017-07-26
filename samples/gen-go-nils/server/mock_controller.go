// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

package server

import (
	context "context"
	models "github.com/Clever/wag/samples/gen-go-nils/models"
	gomock "github.com/golang/mock/gomock"
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

// NilCheck mocks base method
func (_m *MockController) NilCheck(ctx context.Context, i *models.NilCheckInput) error {
	ret := _m.ctrl.Call(_m, "NilCheck", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// NilCheck indicates an expected call of NilCheck
func (_mr *MockControllerMockRecorder) NilCheck(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NilCheck", arg0, arg1)
}
