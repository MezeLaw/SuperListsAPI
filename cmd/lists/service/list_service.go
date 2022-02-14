package service

import (
	"SuperListsAPI/cmd/lists/models"
)

//go:generate mockgen -source=list_service.go -destination lists_service_mock.go -package service

type IListRepository interface {
	Create(list models.List) (*models.List, error)
	GetLists(userId string) (*[]models.List, error)
	Get(listId string) (*models.List, error)
	Update(list models.List) (*models.List, error)
	Delete(listId string) (*int, error)
}

type ListService struct {
	listRepository IListRepository
}

func NewListService(repository IListRepository) ListService {
	return ListService{listRepository: repository}
}

func (ls *ListService) Create(list models.List) (*models.List, error) {
	return ls.listRepository.Create(list)
}

func (ls *ListService) GetLists(userId string) (*[]models.List, error) {
	return ls.listRepository.GetLists(userId)
}

func (ls *ListService) Get(listId string) (*models.List, error) {
	return ls.listRepository.Get(listId)
}

func (ls *ListService) Update(list models.List) (*models.List, error) {
	return ls.listRepository.Update(list)
}

func (ls *ListService) Delete(listId string) (*int, error) {
	return ls.listRepository.Delete(listId)
}
