package repository

import (
	"SuperListsAPI/cmd/userLists/models"
	"gorm.io/gorm"
	"strconv"
)

type UserListRepository struct {
	db *gorm.DB
}

func NewUserListRepository(gormDB *gorm.DB) UserListRepository {
	return UserListRepository{db: gormDB}
}

func (ulr *UserListRepository) Create(list models.UserList) (*models.UserList, error) {

	if result := ulr.db.Create(&list); result.Error != nil {
		return nil, result.Error
	}
	return &list, nil
}

func (ulr *UserListRepository) Get(userListID string) (*models.UserList, error) {

	userList := models.UserList{}

	if result := ulr.db.Find(&userList, userListID); result.Error != nil {
		return nil, result.Error
	}
	return &userList, nil
}

func (ulr *UserListRepository) Delete(userListID string) (*int, error) {

	if result := ulr.db.Delete(&models.UserList{}, userListID); result.Error != nil {
		return nil, result.Error
	}

	deletedID, _ := strconv.Atoi(userListID)

	return &deletedID, nil

}

func (ulr *UserListRepository) GetUserListsByUserID(userId string) (*[]models.UserList, error) {

	var userLists []models.UserList

	if result := ulr.db.Find(&userLists, userId); result.Error != nil {
		return nil, result.Error
	}

	return &userLists, nil
}
