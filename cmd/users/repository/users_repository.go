package repository

//go:generate mockgen -source=users_repository.go -destination users_repository_mock.go -package repository
import (
	"SuperListsAPI/cmd/auth/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	database *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return UsersRepository{db}
}

func (ur *UsersRepository) GetUser(email string) (*models.User, error) {
	var user models.User

	if result := ur.database.Where("email = ?", email).Find(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
