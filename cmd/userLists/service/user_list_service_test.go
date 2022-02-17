package service

import (
	"SuperListsAPI/cmd/userLists/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewUserListService(t *testing.T) {
	type args struct {
		repository IUserListRepository
	}
	tests := []struct {
		name string
		args args
		want UserListService
	}{
		{
			name: "Service with nil repo should pass",
			args: args{repository: nil},
			want: NewUserListService(nil),
		},
		{
			name: "Service with no nil repo should pass",
			args: args{repository: NewMockIUserListRepository(gomock.NewController(t))},
			want: NewUserListService(NewMockIUserListRepository(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserListService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserListService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserListService_Create(t *testing.T) {

	validUserList := GetValidUserList()

	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Create(gomock.Any()).Return(&validUserList, nil)

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Create(validUserList)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.ListID, uint(1))

}

func TestUserListService_Create_Error(t *testing.T) {

	validUserList := GetValidUserList()

	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from userList repository"))

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Create(validUserList)

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestUserListService_Get(t *testing.T) {

	validUserList := GetValidUserList()

	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Get(gomock.Any()).Return(&validUserList, nil)

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Get("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.ListID, uint(1))

}

func TestUserListService_Get_Error(t *testing.T) {

	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from userList repository"))

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Get("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestUserListService_Delete(t *testing.T) {
	deletedID := 1
	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Delete("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, *result, 1)

}

func TestUserListService_Delete_Error(t *testing.T) {

	mockedRepository := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepository.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from userList repository"))

	userListService := NewUserListService(mockedRepository)

	result, err := userListService.Delete("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestUserListService_GetUserListsByUserID(t *testing.T) {
	mockedRepo := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetUserListsByUserID(gomock.Any()).Return(&[]models.UserList{GetValidUserList()}, nil)

	userListService := NewUserListService(mockedRepo)

	result, err := userListService.GetUserListsByUserID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestUserListService_GetUserListsByUserID_Error(t *testing.T) {
	mockedRepo := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetUserListsByUserID(gomock.Any()).Return(nil, errors.New("error from userLists repository"))

	userListService := NewUserListService(mockedRepo)

	result, err := userListService.GetUserListsByUserID("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestUserListService_GetUserListsByListID(t *testing.T) {
	mockedRepo := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetUserListsByListID(gomock.Any()).Return(&[]models.UserList{GetValidUserList()}, nil)

	userListService := NewUserListService(mockedRepo)

	result, err := userListService.GetUserListsByListID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestUserListService_GetUserListsByListID_Error(t *testing.T) {
	mockedRepo := NewMockIUserListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetUserListsByListID(gomock.Any()).Return(nil, errors.New("error from userLists repository"))

	userListService := NewUserListService(mockedRepo)

	result, err := userListService.GetUserListsByListID("1")

	assert.Error(t, err)
	assert.Nil(t, result)

}

func GetValidUserList() models.UserList {
	return models.UserList{
		Model:  gorm.Model{},
		ListID: 1,
		UserID: 1,
	}
}
