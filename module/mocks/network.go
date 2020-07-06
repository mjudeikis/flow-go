// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapperlabs/flow-go/module (interfaces: Network,Local,Requester)

// Package mocks is a generated GoMock package.
package mocks

import (
	crypto "github.com/dapperlabs/flow-go/crypto"
	hash "github.com/dapperlabs/flow-go/crypto/hash"
	flow "github.com/dapperlabs/flow-go/model/flow"
	module "github.com/dapperlabs/flow-go/module"
	network "github.com/dapperlabs/flow-go/network"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNetwork is a mock of Network interface
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockNetwork) Register(arg0 byte, arg1 network.Engine) (network.Conduit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(network.Conduit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockNetworkMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockNetwork)(nil).Register), arg0, arg1)
}

// MockLocal is a mock of Local interface
type MockLocal struct {
	ctrl     *gomock.Controller
	recorder *MockLocalMockRecorder
}

// MockLocalMockRecorder is the mock recorder for MockLocal
type MockLocalMockRecorder struct {
	mock *MockLocal
}

// NewMockLocal creates a new mock instance
func NewMockLocal(ctrl *gomock.Controller) *MockLocal {
	mock := &MockLocal{ctrl: ctrl}
	mock.recorder = &MockLocalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocal) EXPECT() *MockLocalMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockLocal) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockLocalMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockLocal)(nil).Address))
}

// NodeID mocks base method
func (m *MockLocal) NodeID() flow.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(flow.Identifier)
	return ret0
}

// NodeID indicates an expected call of NodeID
func (mr *MockLocalMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockLocal)(nil).NodeID))
}

// NotMeFilter mocks base method
func (m *MockLocal) NotMeFilter() flow.IdentityFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotMeFilter")
	ret0, _ := ret[0].(flow.IdentityFilter)
	return ret0
}

// NotMeFilter indicates an expected call of NotMeFilter
func (mr *MockLocalMockRecorder) NotMeFilter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotMeFilter", reflect.TypeOf((*MockLocal)(nil).NotMeFilter))
}

// Sign mocks base method
func (m *MockLocal) Sign(arg0 []byte, arg1 hash.Hasher) (crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockLocalMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockLocal)(nil).Sign), arg0, arg1)
}

// MockRequester is a mock of Requester interface
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterMockRecorder
}

// MockRequesterMockRecorder is the mock recorder for MockRequester
type MockRequesterMockRecorder struct {
	mock *MockRequester
}

// NewMockRequester creates a new mock instance
func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequester) EXPECT() *MockRequesterMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockRequester) Request(arg0 flow.Identifier, arg1 module.HandleFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockRequesterMockRecorder) Request(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockRequester)(nil).Request), arg0, arg1)
}
