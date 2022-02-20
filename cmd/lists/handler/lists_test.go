package handler

import (
	listItemModels "SuperListsAPI/cmd/listItems/models"
	"SuperListsAPI/cmd/lists/models"
	userListsModel "SuperListsAPI/cmd/userLists/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListHandler_Create(t *testing.T) {

	validList := GetValidList()
	validUserList := GetValidUserList()
	inviteCode, _ := uuid.NewV4()

	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Create(gomock.Any()).Return(&validList, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Create(gomock.Any()).Return(&validUserList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.POST("/", listHandler.Create)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/v1/lists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestListHandler_Create_Returns_Service_Error(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()

	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from list service"))
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.POST("/", listHandler.Create)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/v1/lists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestListHandler_Create_Error_On_User_List_Service(t *testing.T) {

	validList := GetValidList()
	inviteCode, _ := uuid.NewV4()

	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Create(gomock.Any()).Return(&validList, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error on userList service"))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.POST("/", listHandler.Create)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/v1/lists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestListHandler_Create_Missing_Required_Value(t *testing.T) {

	invalidList := models.List{
		Description: "invalid description",
	}

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.POST("/", listHandler.Create)
	}

	jsonDto, _ := json.Marshal(invalidList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/v1/lists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestListHandler_Create_Invalid_Request_Body(t *testing.T) {

	invalidList := map[string]interface{}{
		"name": 1,
	}

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.POST("/", listHandler.Create)
	}

	jsonDto, _ := json.Marshal(invalidList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/v1/lists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestListHandler_GetLists(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	lists := []models.List{validList, validList}

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().GetLists(gomock.Any()).Return(&lists, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_GetLists_Returns_Service_Error(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	lists := []models.List{validList, validList}

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().GetLists(gomock.Any()).Return(&lists, errors.New("error from list service"))
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_GetLists_Returns_No_Header_ID_Error(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))
	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_GetLists_Returns_Invalid_ID_Error(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)
	req.Header.Add("USER_ID", "invalidID")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_GetLists_Returns_No_Content(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().GetLists(gomock.Any()).Return(nil, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestListHandler_GetLists_Returns_No_Content_With_Lists_Not_Nil(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().GetLists(gomock.Any()).Return(&[]models.List{}, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/")
	{
		v1.GET("/", listHandler.GetLists)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestListHandler_Get(t *testing.T) {
	validList := GetValidList()

	validListItem := GetValidListItem()

	listItemsReturned := []listItemModels.ListItem{validListItem}

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().GetItemsListByListID(gomock.Any()).Return(&listItemsReturned, nil)
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.GET("/:id", listHandler.Get)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_Get_Error_On_ListItem_Service(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().GetItemsListByListID(gomock.Any()).Return(nil, errors.New("error from list item service"))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.GET("/:id", listHandler.Get)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Get_Returns_Not_Found(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(nil, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.GET("/:id", listHandler.Get)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListHandler_Get_Returns_Service_Error(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(&validList, errors.New("list service error"))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.GET("/:id", listHandler.Get)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Get_Missing_ID_On_URL(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.GET("/:id", listHandler.Get)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1/lists/invalidID", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Update(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Update(gomock.Any()).Return(&validList, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_Update_Returns_Service_Error(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Update(gomock.Any()).Return(&validList, errors.New("error from list service"))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Update_Invalid_ID(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/invalidID", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Update_Mismatch_ID(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/2", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Update_Invalid_Request_Body(t *testing.T) {

	validList := map[string]interface{}{
		"name": 1,
	}

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/2", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Update_Missing_Request_Body_Mandatory_Values(t *testing.T) {

	validList := map[string]interface{}{
		"name": "Lista de tareas",
		"id":   2,
	}

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/2", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Update_Returns_Not_Found_List_To_Update(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Update(gomock.Any()).Return(nil, nil)
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.PUT("/:id", listHandler.Update)
	}

	jsonDto, _ := json.Marshal(validList)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPut, "/v1/lists/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListHandler_Delete(t *testing.T) {
	validList := GetValidList()

	deletedUserListID := 1

	userList := userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}

	userList.ID = 1

	userLists := []userListsModel.UserList{userList}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	deletedID := "1"
	deletedListItemQty := 1

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	userListService.EXPECT().Delete(gomock.Any()).Return(&deletedUserListID, nil)

	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(&deletedListItemQty, nil)
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_Delete_List_Item_Service_Error(t *testing.T) {
	validList := GetValidList()

	deletedUserListID := 1

	userList := userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}

	userList.ID = 1

	userLists := []userListsModel.UserList{userList}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	deletedID := "1"

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	userListService.EXPECT().Delete(gomock.Any()).Return(&deletedUserListID, nil)

	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(nil, errors.New("Error on listItemService"))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Delete_Error_On_User_List_Service(t *testing.T) {
	validList := GetValidList()

	userList := userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}

	userList.ID = 1

	userLists := []userListsModel.UserList{userList}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	deletedID := "1"

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	userListService.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from user list service"))

	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Delete_Error_On_User_List_Service_Nil_Results(t *testing.T) {
	validList := GetValidList()

	userList := userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}

	userList.ID = 1

	userLists := []userListsModel.UserList{userList}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	deletedID := "1"

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	userListService.EXPECT().Delete(gomock.Any()).Return(nil, nil)

	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestListHandler_Delete_Not_Owner_List(t *testing.T) {
	validList := GetValidList()
	userLists := []userListsModel.UserList{userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}, userListsModel.UserList{
		ListID: 1,
		UserID: 2,
	}}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 2
	validList.ID = 1

	deletedUserListQty := 1
	listItemDeletedQty := 1

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	userListService.EXPECT().Delete(gomock.Any()).Return(&deletedUserListQty, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(&listItemDeletedQty, nil)
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_Delete_Error_Getting_List(t *testing.T) {

	userLists := []userListsModel.UserList{userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}}

	listService := NewMockIListService(gomock.NewController(t))

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	listService.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from list service executing get"))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Delete_Returns_Service_Error(t *testing.T) {
	validList := GetValidList()
	userLists := []userListsModel.UserList{{UserID: 1, ListID: 1}}
	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))

	listService.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from list service"))
	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Delete_Invalid_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/invalidID", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Delete_Invalid_User_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "a")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Delete_Missing_User_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Delete_Returns_Error_On_User_Lists_Delete(t *testing.T) {
	rowsQtyDeleted := 1
	validList := GetValidList()
	userLists := []userListsModel.UserList{GetValidUserList()}

	listService := NewMockIListService(gomock.NewController(t))

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListService.EXPECT().GetUserListsByListID(gomock.Any()).Return(&userLists, nil)
	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)

	userListService.EXPECT().Delete(gomock.Any()).Return(&rowsQtyDeleted, nil)

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listItemService.EXPECT().DeleteListItemsByListID(gomock.Any()).Return(nil, errors.New("error from list items service"))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func GetValidList() models.List {
	return models.List{
		Name:        "Mocked list name",
		Description: "Mocked list description",
	}
}

func GetValidUserList() userListsModel.UserList {
	return userListsModel.UserList{
		ListID: 1,
		UserID: 1,
	}
}

func GetValidListItem() listItemModels.ListItem {
	return listItemModels.ListItem{
		ListID:      1,
		UserID:      1,
		Title:       "Trabajar",
		Description: "Programar 9 horas como mono",
		IsDone:      false,
	}
}

func TestListHandler_JoinList(t *testing.T) {

	getValidUserList := GetValidUserList()

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	userListService.EXPECT().Create(gomock.Any()).Return(&getValidUserList, nil)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/joinList")
	{
		v1.POST("/:listID", listHandler.JoinList)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/lists/joinList/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)

}

func TestListHandler_JoinList_No_Cant_Join(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	userListService.EXPECT().Create(gomock.Any()).Return(nil, nil)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/joinList")
	{
		v1.POST("/:listID", listHandler.JoinList)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/lists/joinList/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)

}

func TestListHandler_JoinList_Error(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	userListService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from userList service"))

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/joinList")
	{
		v1.POST("/:listID", listHandler.JoinList)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/lists/joinList/1", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

}

func TestListHandler_JoinList_Invalid_List_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/joinList")
	{
		v1.POST("/:listID", listHandler.JoinList)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/lists/joinList/a", nil)
	req.Header.Add("USER_ID", "1")
	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestListHandler_JoinList_Missing_User_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	userListService := NewMockIUserListService(gomock.NewController(t))

	listItemService := NewMockIListItemService(gomock.NewController(t))
	listHandler := NewListHandler(listService, userListService, listItemService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists/joinList")
	{
		v1.POST("/:listID", listHandler.JoinList)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/lists/joinList/1", nil)
	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestNewListHandler(t *testing.T) {
	type args struct {
		service          IListService
		userListService  IUserListService
		listItemsService IListItemService
	}
	tests := []struct {
		name string
		args args
		want ListHandler
	}{
		{
			name: "Test with nil services should pass",
			args: args{
				service:          nil,
				userListService:  nil,
				listItemsService: nil,
			},
			want: NewListHandler(nil, nil, nil),
		},
		{
			name: "Test with no nil services should pass",
			args: args{
				service:          NewMockIListService(gomock.NewController(t)),
				userListService:  NewMockIUserListService(gomock.NewController(t)),
				listItemsService: NewMockIListItemService(gomock.NewController(t)),
			},
			want: NewListHandler(NewMockIListService(gomock.NewController(t)), NewMockIUserListService(gomock.NewController(t)), NewMockIListItemService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewListHandler(tt.args.service, tt.args.userListService, tt.args.listItemsService), "NewListHandler(%v, %v, %v)", tt.args.service, tt.args.userListService, tt.args.listItemsService)
		})
	}
}
