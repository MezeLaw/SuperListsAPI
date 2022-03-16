package service

import (
	"SuperListsAPI/cmd/auth/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepo IUserRepository
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		{
			name: "Test with nil repo should pass",
			args: args{userRepo: nil},
			want: NewUserService(nil),
		},
		{
			name: "Test with no nil repo should pass",
			args: args{userRepo: NewMockIUserRepository(gomock.NewController(t))},
			want: NewUserService(NewMockIUserRepository(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUser(t *testing.T) {

	user := GetValidUser()

	repo := NewMockIUserRepository(gomock.NewController(t))
	repo.EXPECT().GetUser(gomock.Any()).Return(&user, nil)

	service := NewUserService(repo)

	result, err := service.GetUser("mezetest@gmail.com")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserService_GetUser_Error(t *testing.T) {

	repo := NewMockIUserRepository(gomock.NewController(t))
	repo.EXPECT().GetUser(gomock.Any()).Return(nil, errors.New("error from repo"))

	service := NewUserService(repo)

	result, err := service.GetUser("mezetest@gmail.com")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func GetValidUser() models.User {
	return models.User{
		Name:     "Meze Test",
		Email:    "mezetest@gmail.com",
		Password: "password",
		Role:     "CAPO DI TUTTI CAPI",
	}
}
