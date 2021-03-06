// Copyright (c) 2016-2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/lib/healthcheck (interfaces: Checker,Filter,PassiveFilter)

// Package mockhealthcheck is a generated GoMock package.
package mockhealthcheck

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	stringset "github.com/uber/kraken/utils/stringset"
	reflect "reflect"
)

// MockChecker is a mock of Checker interface
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockChecker) Check(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockCheckerMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockChecker)(nil).Check), arg0, arg1)
}

// MockFilter is a mock of Filter interface
type MockFilter struct {
	ctrl     *gomock.Controller
	recorder *MockFilterMockRecorder
}

// MockFilterMockRecorder is the mock recorder for MockFilter
type MockFilterMockRecorder struct {
	mock *MockFilter
}

// NewMockFilter creates a new mock instance
func NewMockFilter(ctrl *gomock.Controller) *MockFilter {
	mock := &MockFilter{ctrl: ctrl}
	mock.recorder = &MockFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilter) EXPECT() *MockFilterMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockFilter) Run(arg0 stringset.Set) stringset.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(stringset.Set)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockFilterMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFilter)(nil).Run), arg0)
}

// MockPassiveFilter is a mock of PassiveFilter interface
type MockPassiveFilter struct {
	ctrl     *gomock.Controller
	recorder *MockPassiveFilterMockRecorder
}

// MockPassiveFilterMockRecorder is the mock recorder for MockPassiveFilter
type MockPassiveFilterMockRecorder struct {
	mock *MockPassiveFilter
}

// NewMockPassiveFilter creates a new mock instance
func NewMockPassiveFilter(ctrl *gomock.Controller) *MockPassiveFilter {
	mock := &MockPassiveFilter{ctrl: ctrl}
	mock.recorder = &MockPassiveFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPassiveFilter) EXPECT() *MockPassiveFilterMockRecorder {
	return m.recorder
}

// Failed mocks base method
func (m *MockPassiveFilter) Failed(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Failed", arg0)
}

// Failed indicates an expected call of Failed
func (mr *MockPassiveFilterMockRecorder) Failed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MockPassiveFilter)(nil).Failed), arg0)
}

// Run mocks base method
func (m *MockPassiveFilter) Run(arg0 stringset.Set) stringset.Set {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(stringset.Set)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockPassiveFilterMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPassiveFilter)(nil).Run), arg0)
}
