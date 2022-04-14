package repository

import (
	"SuperListsAPI/cmd/lists/models"
	userListsModel "SuperListsAPI/cmd/userLists/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"log"
)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) ListRepository {
	return ListRepository{db: db}
}

func (lr *ListRepository) Create(list models.List) (*models.List, error) {

	inviteCode, _ := uuid.NewV4()

	list.InviteCode = inviteCode.String()

	if result := lr.db.Create(&list); result.Error != nil {
		return nil, result.Error
	}

	return &list, nil

}

func (lr *ListRepository) GetLists(userId string) (*[]models.List, error) {

	var lists []models.List
	var userLists []userListsModel.UserList
	listsIDs := []int{}

	if result := lr.db.Where("user_id = ?", userId).Find(&userLists); result.Error != nil {
		return nil, result.Error
	}
	//TODO Add tests for this if statement
	if len(userLists) < 1 {
		log.Print("No user-lists for user with ID: ", userId)
		return &[]models.List{}, nil
	}

	for _, userList := range userLists {
		listsIDs = append(listsIDs, int(userList.ListID))
	}

	if result := lr.db.Find(&lists, listsIDs); result.Error != nil {
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

func (lr *ListRepository) Delete(idToDelete string) (*string, error) {
	//db.Delete(&users, []int{1,2,3})
	if result := lr.db.Delete(&models.List{}, idToDelete); result.Error != nil || result.RowsAffected < 1 {
		return nil, result.Error
	}

	return &idToDelete, nil

}

func (lr *ListRepository) GetListByInvitationCode(invitationCode string) (*models.List, error) {

	var list models.List

	if result := lr.db.Where("invite_code = ?", invitationCode).Find(&list); result.Error != nil {
		return nil, result.Error
	}

	return &list, nil
}

func (lr *ListRepository) BulkDelete(listsToDelete []models.List) (*int, error) {

	idsToDelete := extractIdsFromListsToDelete(listsToDelete)

	result := lr.db.Delete(&models.List{}, idsToDelete)

	if result.Error != nil {
		return nil, result.Error
	}

	rowsDeleted := int(result.RowsAffected)

	return &rowsDeleted, nil
}

func extractIdsFromListsToDelete(listsToDelete []models.List) *[]uint {

	var idsToDelete []uint

	for _, list := range listsToDelete {
		idsToDelete = append(idsToDelete, list.ID)
	}

	return &idsToDelete
}