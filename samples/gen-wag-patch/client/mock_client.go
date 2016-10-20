// Automatically generated by MockGen. DO NOT EDIT!
// Source: interface.go

package client

import (
	context "context"
	models "github.com/Clever/wag/samples/gen-wag-patch/models"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Wagpatch(ctx context.Context, i *models.PatchData) (*models.Data, error) {
	ret := _m.ctrl.Call(_m, "Wagpatch", ctx, i)
	ret0, _ := ret[0].(*models.Data)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Wagpatch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Wagpatch", arg0, arg1)
}
