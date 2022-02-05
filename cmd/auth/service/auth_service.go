package service

import (
	"SuperListsAPI/cmd/auth/models"
)

type IAuthRepository interface {
	Login(payload models.LoginPayload) (*string, error)
	SignUp(user *models.User) (*models.User, error)
}

type AuthService struct {
	authRepo IAuthRepository
}

func NewAuthService(authRepo IAuthRepository) AuthService {
	return AuthService{authRepo: authRepo}
}

func (authService *AuthService) Login(payload models.LoginPayload) (*string, error) {

	token, err := authService.authRepo.Login(payload)
	return token, err

}

func (authService *AuthService) SignUp(userRequest *models.User) (*models.User, error) {

	result, err := authService.authRepo.SignUp(userRequest)

	return result, err
}
