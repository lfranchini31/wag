// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package server

import (
	models "github.com/Clever/wag/generated/models"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *_MockControllerRecorder
}

// Recorder for MockController (not exported)
type _MockControllerRecorder struct {
	mock *MockController
}

func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &_MockControllerRecorder{mock}
	return mock
}

func (_m *MockController) EXPECT() *_MockControllerRecorder {
	return _m.recorder
}

func (_m *MockController) GetBooks(ctx context.Context, i *models.GetBooksInput) (*[]models.Book, error) {
	ret := _m.ctrl.Call(_m, "GetBooks", ctx, i)
	ret0, _ := ret[0].(*[]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerRecorder) GetBooks(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBooks", arg0, arg1)
}

func (_m *MockController) GetBookByID(ctx context.Context, i *models.GetBookByIDInput) (models.GetBookByIDOutput, error) {
	ret := _m.ctrl.Call(_m, "GetBookByID", ctx, i)
	ret0, _ := ret[0].(models.GetBookByIDOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerRecorder) GetBookByID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBookByID", arg0, arg1)
}

func (_m *MockController) CreateBook(ctx context.Context, i *models.CreateBookInput) (*models.Book, error) {
	ret := _m.ctrl.Call(_m, "CreateBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockControllerRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateBook", arg0, arg1)
}
