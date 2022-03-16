package service

//go:generate mockgen -source=users_service.go -destination users_service_mock.go -package service
import (
	"SuperListsAPI/cmd/auth/models"
	"log"
)

type IUserRepository interface {
	GetUser(email string) (*models.User, error)
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (us *UserService) GetUser(email string) (*models.User, error) {

	result, err := us.userRepo.GetUser(email)

	if err != nil {
		log.Println("Error retrieving user")
		return nil, err
	}

	return result, nil
}
