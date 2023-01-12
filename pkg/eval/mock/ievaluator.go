// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/eval/ievaluator.go

// Package evalmock is a generated GoMock package.
package evalmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// MockIEvaluator is a mock of IEvaluator interface.
type MockIEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockIEvaluatorMockRecorder
}

// MockIEvaluatorMockRecorder is the mock recorder for MockIEvaluator.
type MockIEvaluatorMockRecorder struct {
	mock *MockIEvaluator
}

// NewMockIEvaluator creates a new mock instance.
func NewMockIEvaluator(ctrl *gomock.Controller) *MockIEvaluator {
	mock := &MockIEvaluator{ctrl: ctrl}
	mock.recorder = &MockIEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEvaluator) EXPECT() *MockIEvaluatorMockRecorder {
	return m.recorder
}

// GetState mocks base method.
func (m *MockIEvaluator) GetState() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockIEvaluatorMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockIEvaluator)(nil).GetState))
}

// ResolveBooleanValue mocks base method.
func (m *MockIEvaluator) ResolveBooleanValue(reqID, flagKey string, context *structpb.Struct) (bool, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBooleanValue", reqID, flagKey, context)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveBooleanValue indicates an expected call of ResolveBooleanValue.
func (mr *MockIEvaluatorMockRecorder) ResolveBooleanValue(reqID, flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBooleanValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveBooleanValue), reqID, flagKey, context)
}

// ResolveFloatValue mocks base method.
func (m *MockIEvaluator) ResolveFloatValue(reqID, flagKey string, context *structpb.Struct) (float64, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveFloatValue", reqID, flagKey, context)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveFloatValue indicates an expected call of ResolveFloatValue.
func (mr *MockIEvaluatorMockRecorder) ResolveFloatValue(reqID, flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveFloatValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveFloatValue), reqID, flagKey, context)
}

// ResolveIntValue mocks base method.
func (m *MockIEvaluator) ResolveIntValue(reqID, flagKey string, context *structpb.Struct) (int64, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveIntValue", reqID, flagKey, context)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveIntValue indicates an expected call of ResolveIntValue.
func (mr *MockIEvaluatorMockRecorder) ResolveIntValue(reqID, flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveIntValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveIntValue), reqID, flagKey, context)
}

// ResolveObjectValue mocks base method.
func (m *MockIEvaluator) ResolveObjectValue(reqID, flagKey string, context *structpb.Struct) (map[string]any, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveObjectValue", reqID, flagKey, context)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveObjectValue indicates an expected call of ResolveObjectValue.
func (mr *MockIEvaluatorMockRecorder) ResolveObjectValue(reqID, flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveObjectValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveObjectValue), reqID, flagKey, context)
}

// ResolveStringValue mocks base method.
func (m *MockIEvaluator) ResolveStringValue(reqID, flagKey string, context *structpb.Struct) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveStringValue", reqID, flagKey, context)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveStringValue indicates an expected call of ResolveStringValue.
func (mr *MockIEvaluatorMockRecorder) ResolveStringValue(reqID, flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveStringValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveStringValue), reqID, flagKey, context)
}

// SetState mocks base method.
func (m *MockIEvaluator) SetState(source, state string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", source, state)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetState indicates an expected call of SetState.
func (mr *MockIEvaluatorMockRecorder) SetState(source, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockIEvaluator)(nil).SetState), source, state)
}