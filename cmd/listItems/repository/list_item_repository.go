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

func (lir *ListItemRepository) BulkDelete(tasksToDelete []models.ListItem) (*int, error) {

	//db.Delete(&users, []int{1,2,3})
	//// DELETE FROM users WHERE id IN (1,2,3);
	idsToDelete := extractIdsFromTasksToDelete(tasksToDelete)

	result := lir.db.Delete(&models.ListItem{}, idsToDelete)

	if result.Error != nil {
		return nil, result.Error
	}

	rowsDeleted := int(result.RowsAffected)

	return &rowsDeleted, nil

}

func (lir *ListItemRepository) MarkAsCompleted(tasksToDelete []models.ListItem) (*int, error) {

	//db.Delete(&users, []int{1,2,3})
	//// DELETE FROM users WHERE id IN (1,2,3);
	idsToUpdate := extractIdsFromTasksToUpdate(tasksToDelete)

	result := lir.db.Table("list_items").Where("id IN ?", idsToUpdate).Updates(map[string]interface{}{"is_done": true})
	//result := lir.db.Model(models.ListItem{}).Where("id in ?", idsToDelete).Updates(map[string]interface{}{"is_done": true})
	if result.Error != nil {
		return nil, result.Error
	}

	rowsDeleted := int(result.RowsAffected)

	return &rowsDeleted, nil

}

func (lir *ListItemRepository) MarkAsPending(tasksToDelete []models.ListItem) (*int, error) {

	//db.Delete(&users, []int{1,2,3})
	//// DELETE FROM users WHERE id IN (1,2,3);
	idsToUpdate := extractIdsFromTasksToUpdate(tasksToDelete)

	result := lir.db.Table("list_items").Where("id IN ?", idsToUpdate).Updates(map[string]interface{}{"is_done": false})

	if result.Error != nil {
		return nil, result.Error
	}

	rowsDeleted := int(result.RowsAffected)

	return &rowsDeleted, nil

}

func extractIdsFromTasksToDelete(tasksToDelete []models.ListItem) *[]uint {

	var idsToDelete []uint

	for _, task := range tasksToDelete {
		idsToDelete = append(idsToDelete, task.ID)
	}

	return &idsToDelete
}

func extractIdsFromTasksToUpdate(tasksToDelete []models.ListItem) []int {

	var idsToUpdate []int

	for _, task := range tasksToDelete {
		idsToUpdate = append(idsToUpdate, int(task.ID))
	}

	return idsToUpdate
}
