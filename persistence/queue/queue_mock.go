// Code generated by MockGen. DO NOT EDIT.
// Source: persistence/queue/queue.go

// Package queue is a generated GoMock package.
package queue

import (
	gomock "github.com/golang/mock/gomock"
	packets "github.com/DrmagicE/gmqtt/pkg/packets"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// Init mocks base method
func (m *MockStore) Init(opts *InitOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockStoreMockRecorder) Init(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockStore)(nil).Init), opts)
}

// Clean mocks base method
func (m *MockStore) Clean() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clean indicates an expected call of Clean
func (mr *MockStoreMockRecorder) Clean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockStore)(nil).Clean))
}

// Add mocks base method
func (m *MockStore) Add(elem *Elem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", elem)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockStoreMockRecorder) Add(elem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStore)(nil).Add), elem)
}

// Replace mocks base method
func (m *MockStore) Replace(elem *Elem) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", elem)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Replace indicates an expected call of Replace
func (mr *MockStoreMockRecorder) Replace(elem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockStore)(nil).Replace), elem)
}

// Read mocks base method
func (m *MockStore) Read(pids []packets.PacketID) ([]*Elem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", pids)
	ret0, _ := ret[0].([]*Elem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockStoreMockRecorder) Read(pids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockStore)(nil).Read), pids)
}

// ReadInflight mocks base method
func (m *MockStore) ReadInflight(maxSize uint) ([]*Elem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInflight", maxSize)
	ret0, _ := ret[0].([]*Elem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInflight indicates an expected call of ReadInflight
func (mr *MockStoreMockRecorder) ReadInflight(maxSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInflight", reflect.TypeOf((*MockStore)(nil).ReadInflight), maxSize)
}

// Remove mocks base method
func (m *MockStore) Remove(pid packets.PacketID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", pid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockStoreMockRecorder) Remove(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStore)(nil).Remove), pid)
}

// GetPublishedMessage mocks base method
func (m *MockStore) GetPublishedMessage(pid packets.PacketID) (*Publish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishedMessage", pid)
	ret0, _ := ret[0].(*Publish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublishedMessage indicates an expected call of GetPublishedMessage
func (mr *MockStoreMockRecorder) GetPublishedMessage(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishedMessage", reflect.TypeOf((*MockStore)(nil).Remove), pid)
}
