// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: Input)

package pipeline

import (
	sync "sync"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Input interface
type MockInput struct {
	ctrl     *gomock.Controller
	recorder *_MockInputRecorder
}

// Recorder for MockInput (not exported)
type _MockInputRecorder struct {
	mock *MockInput
}

func NewMockInput(ctrl *gomock.Controller) *MockInput {
	mock := &MockInput{ctrl: ctrl}
	mock.recorder = &_MockInputRecorder{mock}
	return mock
}

func (_m *MockInput) EXPECT() *_MockInputRecorder {
	return _m.recorder
}

func (_m *MockInput) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockInputRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockInput) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockInputRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}

func (_m *MockInput) Start(_param0 PluginHelper, _param1 *sync.WaitGroup) error {
	ret := _m.ctrl.Call(_m, "Start", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInputRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}
