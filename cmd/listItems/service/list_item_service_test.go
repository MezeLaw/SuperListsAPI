package service

import (
	"SuperListsAPI/cmd/listItems/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewListItemService(t *testing.T) {
	type args struct {
		repository IListItemRepository
	}
	tests := []struct {
		name string
		args args
		want ListItemService
	}{
		{
			name: "Service with nil repo should pass",
			args: args{nil},
			want: NewListItemService(nil),
		},
		{
			name: "Service with no nil repo should pass",
			args: args{NewMockIListItemRepository(gomock.NewController(t))},
			want: NewListItemService(NewMockIListItemRepository(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListItemService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListItemService_Create(t *testing.T) {

	validListItem := GetValidListItem()

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Create(gomock.Any()).Return(&validListItem, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Create(validListItem)

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestListItemService_Create_Error(t *testing.T) {

	validListItem := GetValidListItem()

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from list item repo"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Create(validListItem)

	assert.Error(t, err)
	assert.Nil(t, result)

}

func TestListItemService_Get(t *testing.T) {

	validListItem := GetValidListItem()

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Get(gomock.Any()).Return(&validListItem, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Get("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestListItemService_Get_Error(t *testing.T) {

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from list item repo"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Get("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListItemService_Delete(t *testing.T) {
	deletedListItemID := 1

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Delete(gomock.Any()).Return(&deletedListItemID, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Delete("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListItemService_Delete_Error(t *testing.T) {

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from list item repo"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Delete("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListItemService_Update(t *testing.T) {
	validListItem := GetValidListItem()

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Update(gomock.Any()).Return(&validListItem, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Update(validListItem)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListItemService_Update_Error(t *testing.T) {
	validListItem := GetValidListItem()

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("error from list item repo"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.Update(validListItem)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListItemService_GetItemsListByListID(t *testing.T) {

	validListItem := GetValidListItem()
	items := []models.ListItem{validListItem}
	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetItemsListByListID(gomock.Any()).Return(&items, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.GetItemsListByListID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListItemService_GetItemsListByListID_Error(t *testing.T) {

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetItemsListByListID(gomock.Any()).Return(nil, errors.New("error from list item repo"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.GetItemsListByListID("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestListItemService_DeleteListItemsByListID(t *testing.T) {
	deletedListItemsQty := 1

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(&deletedListItemsQty, nil)

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.DeleteListItemsByListID("1")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestListItemService_DeleteListItemsByListID_Error(t *testing.T) {

	mockedRepo := NewMockIListItemRepository(gomock.NewController(t))
	mockedRepo.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(nil, errors.New("error from list item repository"))

	listItemService := NewListItemService(mockedRepo)

	result, err := listItemService.DeleteListItemsByListID("1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func GetValidListItem() models.ListItem {
	return models.ListItem{
		ListID:      1,
		UserID:      1,
		Title:       "Hacer la tarea",
		Description: "Completar las funciones cuadraticas, graficarlas y derivarlas",
		IsDone:      false,
	}
}
