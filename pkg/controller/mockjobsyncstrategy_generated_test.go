// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./jobsyncstrategy.go

package controller

import (
	gomock "github.com/golang/mock/gomock"
	v10 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Mock of JobSyncStrategy interface
type MockJobSyncStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockJobSyncStrategyRecorder
}

// Recorder for MockJobSyncStrategy (not exported)
type _MockJobSyncStrategyRecorder struct {
	mock *MockJobSyncStrategy
}

func NewMockJobSyncStrategy(ctrl *gomock.Controller) *MockJobSyncStrategy {
	mock := &MockJobSyncStrategy{ctrl: ctrl}
	mock.recorder = &_MockJobSyncStrategyRecorder{mock}
	return mock
}

func (_m *MockJobSyncStrategy) EXPECT() *_MockJobSyncStrategyRecorder {
	return _m.recorder
}

func (_m *MockJobSyncStrategy) GetOwner(key string) (v1.Object, error) {
	ret := _m.ctrl.Call(_m, "GetOwner", key)
	ret0, _ := ret[0].(v1.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobSyncStrategyRecorder) GetOwner(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOwner", arg0)
}

func (_m *MockJobSyncStrategy) DoesOwnerNeedProcessing(owner v1.Object) bool {
	ret := _m.ctrl.Call(_m, "DoesOwnerNeedProcessing", owner)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockJobSyncStrategyRecorder) DoesOwnerNeedProcessing(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoesOwnerNeedProcessing", arg0)
}

func (_m *MockJobSyncStrategy) GetJobFactory(owner v1.Object) (JobFactory, error) {
	ret := _m.ctrl.Call(_m, "GetJobFactory", owner)
	ret0, _ := ret[0].(JobFactory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobSyncStrategyRecorder) GetJobFactory(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobFactory", arg0)
}

func (_m *MockJobSyncStrategy) GetOwnerCurrentJob(owner v1.Object) string {
	ret := _m.ctrl.Call(_m, "GetOwnerCurrentJob", owner)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockJobSyncStrategyRecorder) GetOwnerCurrentJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOwnerCurrentJob", arg0)
}

func (_m *MockJobSyncStrategy) SetOwnerCurrentJob(owner v1.Object, jobName string) {
	_m.ctrl.Call(_m, "SetOwnerCurrentJob", owner, jobName)
}

func (_mr *_MockJobSyncStrategyRecorder) SetOwnerCurrentJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetOwnerCurrentJob", arg0, arg1)
}

func (_m *MockJobSyncStrategy) DeepCopyOwner(owner v1.Object) v1.Object {
	ret := _m.ctrl.Call(_m, "DeepCopyOwner", owner)
	ret0, _ := ret[0].(v1.Object)
	return ret0
}

func (_mr *_MockJobSyncStrategyRecorder) DeepCopyOwner(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeepCopyOwner", arg0)
}

func (_m *MockJobSyncStrategy) SetOwnerJobSyncCondition(owner v1.Object, conditionType JobSyncConditionType, status v10.ConditionStatus, reason string, message string, updateConditionCheck UpdateConditionCheck) {
	_m.ctrl.Call(_m, "SetOwnerJobSyncCondition", owner, conditionType, status, reason, message, updateConditionCheck)
}

func (_mr *_MockJobSyncStrategyRecorder) SetOwnerJobSyncCondition(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetOwnerJobSyncCondition", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockJobSyncStrategy) OnJobCompletion(owner v1.Object) {
	_m.ctrl.Call(_m, "OnJobCompletion", owner)
}

func (_mr *_MockJobSyncStrategyRecorder) OnJobCompletion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnJobCompletion", arg0)
}

func (_m *MockJobSyncStrategy) OnJobFailure(owner v1.Object) {
	_m.ctrl.Call(_m, "OnJobFailure", owner)
}

func (_mr *_MockJobSyncStrategyRecorder) OnJobFailure(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnJobFailure", arg0)
}

func (_m *MockJobSyncStrategy) UpdateOwnerStatus(original v1.Object, owner v1.Object) error {
	ret := _m.ctrl.Call(_m, "UpdateOwnerStatus", original, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobSyncStrategyRecorder) UpdateOwnerStatus(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateOwnerStatus", arg0, arg1)
}

func (_m *MockJobSyncStrategy) ProcessDeletedOwner(owner v1.Object) error {
	ret := _m.ctrl.Call(_m, "ProcessDeletedOwner", owner)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockJobSyncStrategyRecorder) ProcessDeletedOwner(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ProcessDeletedOwner", arg0)
}
