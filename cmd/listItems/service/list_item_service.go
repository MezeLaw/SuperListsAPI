package service

import "SuperListsAPI/cmd/listItems/models"

//go:generate mockgen -source=list_item_service.go -destination list_item_service_mock.go -package service

type IListItemRepository interface {
	Create(item models.ListItem) (*models.ListItem, error)
	Get(listItemID string) (*models.ListItem, error)
	Update(item models.ListItem) (*models.ListItem, error)
	Delete(listItemID string) (*int, error)
	GetItemsListByListID(listId string) (*[]models.ListItem, error)
	DeleteListItemsByListID(listId string) (*int, error)
	BulkDelete(tasksToDelete []models.ListItem) (*int, error)
}

type ListItemService struct {
	repository IListItemRepository
}

func NewListItemService(repository IListItemRepository) ListItemService {
	return ListItemService{repository: repository}
}

func (lis *ListItemService) Create(item models.ListItem) (*models.ListItem, error) {

	result, err := lis.repository.Create(item)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lis *ListItemService) Get(listItemID string) (*models.ListItem, error) {

	result, err := lis.repository.Get(listItemID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lis *ListItemService) Update(item models.ListItem) (*models.ListItem, error) {
	result, err := lis.repository.Update(item)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lis *ListItemService) Delete(listItemID string) (*int, error) {

	result, err := lis.repository.Delete(listItemID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lis *ListItemService) GetItemsListByListID(listId string) (*[]models.ListItem, error) {

	result, err := lis.repository.GetItemsListByListID(listId)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (lis *ListItemService) DeleteListItemsByListID(listId string) (*int, error) {

	result, err := lis.repository.DeleteListItemsByListID(listId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (lis *ListItemService) BulkDelete(tasksToDelete []models.ListItem) (*int, error) {

	result, err := lis.repository.BulkDelete(tasksToDelete)

	if err != nil {
		return nil, err
	}

	return result, nil
}
