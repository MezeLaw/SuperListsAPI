package service

import (
	"SuperListsAPI/cmd/auth/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	type args struct {
		authRepo IAuthRepository
	}
	tests := []struct {
		name string
		args args
		want AuthService
	}{
		{
			name: "Valid repo should return AuthAService",
			args: args{authRepo: NewMockIAuthRepository(gomock.NewController(t))},
			want: NewAuthService(NewMockIAuthRepository(gomock.NewController(t))),
		},
		{
			name: "Empty repo should return AuthAService",
			args: args{authRepo: nil},
			want: NewAuthService(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.authRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestAuthService_Login(t *testing.T) {

	authRepo := NewMockIAuthRepository(gomock.NewController(t))
	authService := NewAuthService(authRepo)
	tokenResponse := "unToken"
	authRepo.EXPECT().Login(gomock.Any()).Return(&tokenResponse, nil)

	loginPayload := GetValidLoginPayload()

	token, err := authService.Login(*loginPayload)

	assert.NoError(t, err)
	assert.Equal(t, &tokenResponse, token)

}

func TestAuthService_Login_Error(t *testing.T) {

	authRepo := NewMockIAuthRepository(gomock.NewController(t))
	authService := NewAuthService(authRepo)
	authRepo.EXPECT().Login(gomock.Any()).Return(nil, errors.New("error"))

	loginPayload := GetValidLoginPayload()

	token, err := authService.Login(*loginPayload)

	assert.Error(t, err)
	assert.Nil(t, token)

}

func TestAuthService_SignUp(t *testing.T) {
	authRepo := NewMockIAuthRepository(gomock.NewController(t))
	authService := NewAuthService(authRepo)

	validUser := GetValidUser()

	authRepo.EXPECT().SignUp(gomock.Any()).Return(validUser, nil)

	user, err := authService.SignUp(validUser)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Email, "meze@meze.com")
}

func TestAuthService_SignUp_Error(t *testing.T) {
	authRepo := NewMockIAuthRepository(gomock.NewController(t))
	authService := NewAuthService(authRepo)

	authRepo.EXPECT().SignUp(gomock.Any()).Return(nil, errors.New("error"))

	user, err := authService.SignUp(nil)

	assert.Error(t, err)
	assert.Nil(t, user, nil)
}

func GetValidLoginPayload() *models.LoginPayload {
	return &models.LoginPayload{
		Email:    "meze@meze.com",
		Password: "password",
	}
}

func GetValidUser() *models.User {
	return &models.User{
		Name:     "Meze",
		Email:    "meze@meze.com",
		Password: "password",
		Role:     "ADMIN",
	}
}
