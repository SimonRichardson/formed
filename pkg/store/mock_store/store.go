// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SimonRichardson/formed/pkg/store (interfaces: Store)

package mock_store

import (
	models "github.com/SimonRichardson/formed/pkg/models"
	gomock "github.com/golang/mock/gomock"
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
func (_m *MockStore) EXPECT() *MockStoreMockRecorder {
	return _m.recorder
}

// Read mocks base method
func (_m *MockStore) Read() ([]models.User, error) {
	ret := _m.ctrl.Call(_m, "Read")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (_mr *MockStoreMockRecorder) Read() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read")
}

// Write mocks base method
func (_m *MockStore) Write(_param0 []models.User) error {
	ret := _m.ctrl.Call(_m, "Write", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (_mr *MockStoreMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}
