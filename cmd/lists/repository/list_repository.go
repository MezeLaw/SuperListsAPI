package repository

import (
	"SuperListsAPI/cmd/lists/models"
	"gorm.io/gorm"
	"strconv"
)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) ListRepository {
	return ListRepository{db: db}
}

func (lr *ListRepository) Create(list models.List) (*models.List, error) {

	if result := lr.db.Create(&list); result.Error != nil {
		return nil, result.Error
	}

	return &list, nil

}

func (lr *ListRepository) GetLists(userId string) (*[]models.List, error) {

	var lists []models.List
	//TODO definir si desde el controller se debe hacer el cambio para appendear las userLists
	if result := lr.db.Where("user_creator_id = ?", userId).Find(&lists); result.Error != nil {
		return nil, result.Error
	}

	return &lists, nil

}

func (lr *ListRepository) Get(listId string) (*models.List, error) {
	var list models.List

	if result := lr.db.First(&list, listId); result.Error != nil {
		return nil, result.Error
	}
	return &list, nil
}

func (lr *ListRepository) Update(list models.List) (*models.List, error) {

	if result := lr.db.Save(&list); result.Error != nil {
		return nil, result.Error
	}

	return &list, nil
}

func (lr *ListRepository) Delete(listId string) (*int, error) {

	if result := lr.db.Delete(&models.List{}, listId); result.Error != nil || result.RowsAffected < 1 {
		return nil, result.Error
	}
	deletedID, _ := strconv.Atoi(listId)

	return &deletedID, nil

}
