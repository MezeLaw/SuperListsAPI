package service

import "SuperListsAPI/cmd/userLists/models"

//go:generate mockgen -source=user_list_service.go -destination user_list_service_mock.go -package service

type IUserListRepository interface {
	Create(list models.UserList) (*models.UserList, error)
	Get(userListID string) (*models.UserList, error)
	Delete(userListID string) (*int, error)
	GetUserListsByUserID(userId string) (*[]models.UserList, error)
	GetUserListsByListID(listID string) (*[]models.UserList, error)
}

type UserListService struct {
	userListRepository IUserListRepository
}

func NewUserListService(repository IUserListRepository) UserListService {
	return UserListService{userListRepository: repository}
}

func (uls *UserListService) Create(list models.UserList) (*models.UserList, error) {
	return uls.userListRepository.Create(list)
}

func (uls *UserListService) Get(userListID string) (*models.UserList, error) {
	return uls.userListRepository.Get(userListID)
}

func (uls *UserListService) Delete(userListID string) (*int, error) {
	return uls.userListRepository.Delete(userListID)
}

func (uls *UserListService) GetUserListsByUserID(userId string) (*[]models.UserList, error) {
	return uls.userListRepository.GetUserListsByUserID(userId)
}

func (uls *UserListService) GetUserListsByListID(listID string) (*[]models.UserList, error) {
	return uls.userListRepository.GetUserListsByListID(listID)
}
