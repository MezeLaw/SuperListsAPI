// Code generated by MockGen. DO NOT EDIT.
// Source: user_list_service.go

// Package service is a generated GoMock package.
package service

import (
	models "SuperListsAPI/cmd/userLists/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserListRepository is a mock of IUserListRepository interface.
type MockIUserListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserListRepositoryMockRecorder
}

// MockIUserListRepositoryMockRecorder is the mock recorder for MockIUserListRepository.
type MockIUserListRepositoryMockRecorder struct {
	mock *MockIUserListRepository
}

// NewMockIUserListRepository creates a new mock instance.
func NewMockIUserListRepository(ctrl *gomock.Controller) *MockIUserListRepository {
	mock := &MockIUserListRepository{ctrl: ctrl}
	mock.recorder = &MockIUserListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserListRepository) EXPECT() *MockIUserListRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserListRepository) Create(list models.UserList) (*models.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", list)
	ret0, _ := ret[0].(*models.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserListRepositoryMockRecorder) Create(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserListRepository)(nil).Create), list)
}

// Delete mocks base method.
func (m *MockIUserListRepository) Delete(userListID string) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userListID)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserListRepositoryMockRecorder) Delete(userListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserListRepository)(nil).Delete), userListID)
}

// Get mocks base method.
func (m *MockIUserListRepository) Get(userListID string) (*models.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userListID)
	ret0, _ := ret[0].(*models.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserListRepositoryMockRecorder) Get(userListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserListRepository)(nil).Get), userListID)
}

// GetUserListsByUserID mocks base method.
func (m *MockIUserListRepository) GetUserListsByUserID(userId string) (*[]models.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserListsByUserID", userId)
	ret0, _ := ret[0].(*[]models.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserListsByUserID indicates an expected call of GetUserListsByUserID.
func (mr *MockIUserListRepositoryMockRecorder) GetUserListsByUserID(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserListsByUserID", reflect.TypeOf((*MockIUserListRepository)(nil).GetUserListsByUserID), userId)
}