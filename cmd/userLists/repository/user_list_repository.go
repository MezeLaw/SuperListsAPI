package repository

import (
	"SuperListsAPI/cmd/userLists/models"
	"gorm.io/gorm"
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

func (ulr *UserListRepository) Delete(userListIDs *[]uint) (*int, error) {

	result := ulr.db.Delete(&models.UserList{}, userListIDs)

	if result.Error != nil {
		return nil, result.Error
	}

	rowsQtyDeleted := int(result.RowsAffected)

	return &rowsQtyDeleted, nil

}

func (ulr *UserListRepository) GetUserListsByUserID(userId string) (*[]models.UserList, error) {

	var userLists []models.UserList
	if result := ulr.db.Where("user_id = ?", userId).Find(&userLists); result.Error != nil {
		return nil, result.Error
	}

	return &userLists, nil
}

func (ulr *UserListRepository) GetUserListsByListID(listID string) (*[]models.UserList, error) {
	var userLists []models.UserList
	if result := ulr.db.Where("list_id = ?", listID).Find(&userLists); result.Error != nil {
		return nil, result.Error
	}

	return &userLists, nil
}
