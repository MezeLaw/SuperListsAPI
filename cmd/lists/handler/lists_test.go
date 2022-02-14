package handler

import (
	"SuperListsAPI/cmd/lists/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewListHandler(t *testing.T) {
	type args struct {
		service IListService
	}
	tests := []struct {
		name string
		args args
		want ListHandler
	}{
		{
			name: "Test with nil service should pass",
			args: args{service: nil},
			want: NewListHandler(nil),
		},
		{
			name: "Test with no nil service should pass",
			args: args{service: NewMockIListService(gomock.NewController(t))},
			want: NewListHandler(NewMockIListService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListHandler_Create(t *testing.T) {

	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()

	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Create(gomock.Any()).Return(&validList, nil)
	listHandler := NewListHandler(listService)

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
	listService.EXPECT().Create(gomock.Any()).Return(&validList, errors.New("error from list service"))
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(&validList, nil)
	listHandler := NewListHandler(listService)

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

func TestListHandler_Get_Returns_Not_Found(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Get(gomock.Any()).Return(nil, nil)
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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
	listHandler := NewListHandler(listService)

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

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	deletedID := 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Delete(gomock.Any()).Return(&deletedID, nil)
	listHandler := NewListHandler(listService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListHandler_Delete_Returns_Service_Error(t *testing.T) {
	validList := GetValidList()

	inviteCode, _ := uuid.NewV4()
	validList.InviteCode = inviteCode.String()
	validList.UserCreatorID = 1
	validList.ID = 1

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from list service"))
	listHandler := NewListHandler(listService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListHandler_Delete_Invalid_ID(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	listHandler := NewListHandler(listService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/invalidID", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListHandler_Delete_Returns_Not_Found(t *testing.T) {

	listService := NewMockIListService(gomock.NewController(t))
	listService.EXPECT().Delete(gomock.Any()).Return(nil, nil)
	listHandler := NewListHandler(listService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/lists")
	{
		v1.DELETE("/:id", listHandler.Delete)
	}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/v1/lists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func GetValidList() models.List {
	return models.List{
		Name:        "Mocked list name",
		Description: "Mocked list description",
	}
}
