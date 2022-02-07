package repository

import (
	"SuperListsAPI/cmd/auth/models"
	"errors"
	"gorm.io/gorm"
	"strings"
)

//go:generate mockgen -source=auth_repository.go -destination auth_repository_mock.go -package repository
const (
	ADMIN            = "ADMIN"
	USER             = "USER"
	SECRET_KEY       = "rumpelstiltskin"
	ISSUER           = "MezeTheKing"
	EXPIRATION_HOURS = 7
	INVALID_PASSWORD = "invalid credentials"
	EMAIL_NOT_FOUND  = "email not found"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return AuthRepository{database: db}
}

func (authRepo *AuthRepository) SignUp(user *models.User) (*models.User, error) {
	//TODO probar de alguna forma el err del hash
	if err := user.HashPassword(user.Password); err != nil {
		return nil, err
	}

	if user.Role == "" {
		user.Role = USER
	}

	user.Email = strings.ToLower(user.Email)

	if result := authRepo.database.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (authRepo *AuthRepository) Login(payload models.LoginPayload) (*string, error) {

	user := models.User{}

	if result := authRepo.database.Where("email = ?", strings.ToLower(payload.Email)).First(&user); result.Error != nil || result.RowsAffected < 1 {
		if result.Error.Error() == "record not found" {
			return nil, errors.New(EMAIL_NOT_FOUND)
		}
		return nil, result.Error
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		return nil, errors.New(INVALID_PASSWORD)
	}

	jwtWrapper := models.JwtWrapper{
		SecretKey:       SECRET_KEY,
		Issuer:          ISSUER,
		ExpirationHours: EXPIRATION_HOURS,
	}

	token, err := jwtWrapper.GenerateToken(user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	return &token, nil

}
