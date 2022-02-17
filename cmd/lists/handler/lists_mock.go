// Code generated by MockGen. DO NOT EDIT.
// Source: lists.go

// Package handler is a generated GoMock package.
package handler

import (
	models "SuperListsAPI/cmd/lists/models"
	models0 "SuperListsAPI/cmd/userLists/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIListService is a mock of IListService interface.
type MockIListService struct {
	ctrl     *gomock.Controller
	recorder *MockIListServiceMockRecorder
}

// MockIListServiceMockRecorder is the mock recorder for MockIListService.
type MockIListServiceMockRecorder struct {
	mock *MockIListService
}

// NewMockIListService creates a new mock instance.
func NewMockIListService(ctrl *gomock.Controller) *MockIListService {
	mock := &MockIListService{ctrl: ctrl}
	mock.recorder = &MockIListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIListService) EXPECT() *MockIListServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIListService) Create(list models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", list)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIListServiceMockRecorder) Create(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIListService)(nil).Create), list)
}

// Delete mocks base method.
func (m *MockIListService) Delete(idsToDelete []uint) (*[]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", idsToDelete)
	ret0, _ := ret[0].(*[]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIListServiceMockRecorder) Delete(idsToDelete interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIListService)(nil).Delete), idsToDelete)
}

// Get mocks base method.
func (m *MockIListService) Get(listId string) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", listId)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIListServiceMockRecorder) Get(listId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIListService)(nil).Get), listId)
}

// GetLists mocks base method.
func (m *MockIListService) GetLists(userId string) (*[]models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLists", userId)
	ret0, _ := ret[0].(*[]models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLists indicates an expected call of GetLists.
func (mr *MockIListServiceMockRecorder) GetLists(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLists", reflect.TypeOf((*MockIListService)(nil).GetLists), userId)
}

// Update mocks base method.
func (m *MockIListService) Update(list models.List) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", list)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIListServiceMockRecorder) Update(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIListService)(nil).Update), list)
}

// MockIUserListService is a mock of IUserListService interface.
type MockIUserListService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserListServiceMockRecorder
}

// MockIUserListServiceMockRecorder is the mock recorder for MockIUserListService.
type MockIUserListServiceMockRecorder struct {
	mock *MockIUserListService
}

// NewMockIUserListService creates a new mock instance.
func NewMockIUserListService(ctrl *gomock.Controller) *MockIUserListService {
	mock := &MockIUserListService{ctrl: ctrl}
	mock.recorder = &MockIUserListServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserListService) EXPECT() *MockIUserListServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserListService) Create(list models0.UserList) (*models0.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", list)
	ret0, _ := ret[0].(*models0.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUserListServiceMockRecorder) Create(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserListService)(nil).Create), list)
}

// Delete mocks base method.
func (m *MockIUserListService) Delete(userListID string) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userListID)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserListServiceMockRecorder) Delete(userListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserListService)(nil).Delete), userListID)
}

// Get mocks base method.
func (m *MockIUserListService) Get(userListID string) (*models0.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userListID)
	ret0, _ := ret[0].(*models0.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIUserListServiceMockRecorder) Get(userListID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIUserListService)(nil).Get), userListID)
}

// GetUserListsByListID mocks base method.
func (m *MockIUserListService) GetUserListsByListID(listID string) (*[]models0.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserListsByListID", listID)
	ret0, _ := ret[0].(*[]models0.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserListsByListID indicates an expected call of GetUserListsByListID.
func (mr *MockIUserListServiceMockRecorder) GetUserListsByListID(listID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserListsByListID", reflect.TypeOf((*MockIUserListService)(nil).GetUserListsByListID), listID)
}

// GetUserListsByUserID mocks base method.
func (m *MockIUserListService) GetUserListsByUserID(userId string) (*[]models0.UserList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserListsByUserID", userId)
	ret0, _ := ret[0].(*[]models0.UserList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserListsByUserID indicates an expected call of GetUserListsByUserID.
func (mr *MockIUserListServiceMockRecorder) GetUserListsByUserID(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserListsByUserID", reflect.TypeOf((*MockIUserListService)(nil).GetUserListsByUserID), userId)
}
