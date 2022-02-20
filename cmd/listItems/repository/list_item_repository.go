package repository

import (
	"SuperListsAPI/cmd/listItems/models"
	"gorm.io/gorm"
	"strconv"
)

type ListItemRepository struct {
	db *gorm.DB
}

func NewListItemRepository(db *gorm.DB) ListItemRepository {
	return ListItemRepository{db: db}
}

func (lir *ListItemRepository) Create(item models.ListItem) (*models.ListItem, error) {

	if result := lir.db.Create(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (lir *ListItemRepository) Get(listItemID string) (*models.ListItem, error) {

	var listItem models.ListItem

	if result := lir.db.First(&listItem, listItemID); result.Error != nil {
		return nil, result.Error
	}

	return &listItem, nil

}

func (lir *ListItemRepository) Update(item models.ListItem) (*models.ListItem, error) {

	if result := lir.db.Save(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (lir *ListItemRepository) Delete(listItemID string) (*int, error) {

	parsedID, _ := strconv.Atoi(listItemID)

	if result := lir.db.Delete(&models.ListItem{}, listItemID); result.Error != nil {
		return nil, result.Error
	}

	return &parsedID, nil

}

func (lir *ListItemRepository) GetItemsListByListID(listId string) (*[]models.ListItem, error) {

	var listItems []models.ListItem

	if result := lir.db.Where("list_id = ?", listId).Find(&listItems); result.Error != nil {
		return nil, result.Error
	}

	return &listItems, nil

}

func (lir *ListItemRepository) DeleteListItemsByListID(listId string) (*int, error) {
	//TODO Probar esto funcionalmente

	var listItem models.ListItem

	result := lir.db.Where("list_id = ? ", &listId).Delete(&listItem)

	if result.Error != nil {
		return nil, result.Error
	}

	rowsDeleted := int(result.RowsAffected)

	return &rowsDeleted, nil
}
