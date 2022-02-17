package service

import (
	"SuperListsAPI/cmd/lists/models"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewListService(t *testing.T) {
	type args struct {
		repository IListRepository
	}
	tests := []struct {
		name string
		args args
		want ListService
	}{
		{
			name: "Service with nil repository should return a service with nil repo",
			args: args{repository: nil},
			want: NewListService(nil),
		},
		{
			name: "Service with no nil repository should return a service with not nil repo",
			args: args{repository: NewMockIListRepository(gomock.NewController(t))},
			want: NewListService(NewMockIListRepository(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListService_Create(t *testing.T) {

	validList := GetValidList()

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Create(gomock.Any()).Return(&validList, nil)
	listService := NewListService(mockedRepo)

	result, err := listService.Create(validList)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

}

func TestListService_Create_Error(t *testing.T) {

	validList := GetValidList()

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from list repository"))
	listService := NewListService(mockedRepo)

	result, err := listService.Create(validList)

	assert.Error(t, err)
	assert.Empty(t, result)

}

func TestListService_GetLists(t *testing.T) {
	lists := []models.List{GetValidList(), GetValidList()}

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetLists(gomock.Any()).Return(&lists, nil)
	listService := NewListService(mockedRepo)

	result, err := listService.GetLists("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestListService_GetLists_Error(t *testing.T) {

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().GetLists(gomock.Any()).Return(nil, errors.New("error from list repository"))
	listService := NewListService(mockedRepo)

	result, err := listService.GetLists("1")

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestListService_Get(t *testing.T) {

	list := GetValidList()
	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Get(gomock.Any()).Return(&list, nil)
	listService := NewListService(mockedRepo)

	result, err := listService.Get("1")

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestListService_Get_Error(t *testing.T) {

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from list repository"))
	listService := NewListService(mockedRepo)

	result, err := listService.Get("1")

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestListService_Update(t *testing.T) {

	list := GetValidList()
	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Update(gomock.Any()).Return(&list, nil)
	listService := NewListService(mockedRepo)

	result, err := listService.Update(list)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestListService_Update_Error(t *testing.T) {

	list := GetValidList()
	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("error from list repository"))
	listService := NewListService(mockedRepo)

	result, err := listService.Update(list)

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestListService_Delete(t *testing.T) {

	deletedId := 1

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Delete(gomock.Any()).Return(&deletedId, nil)
	listService := NewListService(mockedRepo)

	result, err := listService.Delete([]uint{1})

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestListService_Delete_Error(t *testing.T) {

	mockedRepo := NewMockIListRepository(gomock.NewController(t))
	mockedRepo.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from list repository"))
	listService := NewListService(mockedRepo)

	result, err := listService.Delete([]uint{1})

	assert.Error(t, err)
	assert.Empty(t, result)
}

func GetValidList() models.List {

	inviteCode, _ := uuid.NewV4()

	return models.List{
		Model:         gorm.Model{},
		Name:          "mocked list name",
		Description:   "mocked list description",
		InviteCode:    inviteCode.String(),
		UserCreatorID: 1,
	}
}
