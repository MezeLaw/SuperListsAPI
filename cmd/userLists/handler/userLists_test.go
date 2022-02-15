package handler

import (
	"SuperListsAPI/cmd/userLists/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewUserListHandler(t *testing.T) {
	type args struct {
		userListService IUserListService
	}
	tests := []struct {
		name string
		args args
		want UserListHandler
	}{
		{
			name: "Test with nil service should pass",
			args: args{userListService: nil},
			want: NewUserListHandler(nil),
		},
		{
			name: "Testi with no nil service should pass",
			args: args{userListService: NewMockIUserListService(gomock.NewController(t))},
			want: NewUserListHandler(NewMockIUserListService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserListHandler(tt.args.userListService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserListHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserListHandler_Create(t *testing.T) {

	validUserList := GetValidUserList()

	jsonDto, _ := json.Marshal(validUserList)

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Create(gomock.Any()).Return(&validUserList, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.POST("/", userListHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/userLists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)

}

func TestUserListHandler_Create_Service_Error(t *testing.T) {

	validUserList := GetValidUserList()

	jsonDto, _ := json.Marshal(validUserList)

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("error from user list service"))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.POST("/", userListHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/userLists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

}

func TestUserListHandler_Create_Invalid_JSON(t *testing.T) {

	validUserList := map[string]interface{}{
		"list_id": make(chan string),
	}

	jsonDto, _ := json.Marshal(validUserList)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.POST("/", userListHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/userLists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestUserListHandler_Create_Missing_Mandatory_Values(t *testing.T) {

	validUserList := map[string]interface{}{
		"list_id": 1,
	}

	jsonDto, _ := json.Marshal(validUserList)

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.POST("/", userListHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/userLists/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestUserListHandler_Get(t *testing.T) {
	validUserList := GetValidUserList()

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Get(gomock.Any()).Return(&validUserList, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/:id", userListHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUserListHandler_Get_Service_Error(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Get(gomock.Any()).Return(nil, errors.New("error from user list service"))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/:id", userListHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestUserListHandler_Get_Not_Found(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Get(gomock.Any()).Return(nil, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/:id", userListHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestUserListHandler_Get_Invalid_ID(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/:id", userListHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/invalidId", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestUserListHandler_Delete(t *testing.T) {

	deletedUserListID := 1

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Delete(gomock.Any()).Return(&deletedUserListID, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.DELETE("/:id", userListHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUserListHandler_Delete_Not_Found(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Delete(gomock.Any()).Return(nil, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.DELETE("/:id", userListHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestUserListHandler_Delete_Service_Error(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from user lists service trying to delete"))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.DELETE("/:id", userListHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/userLists/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestUserListHandler_Delete_Invalid_ID(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.DELETE("/:id", userListHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/userLists/invalidID", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestUserListHandler_GetUserListsByUserID(t *testing.T) {
	validUserList := GetValidUserList()

	userLists := []models.UserList{validUserList, validUserList}
	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().GetUserListsByUserID(gomock.Any()).Return(&userLists, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/", userListHandler.GetUserListsByUserID)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUserListHandler_GetUserListsByUserID_No_Content(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().GetUserListsByUserID(gomock.Any()).Return(nil, nil)

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/", userListHandler.GetUserListsByUserID)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNoContent)
}

func TestUserListHandler_GetUserListsByUserID_Service_Error(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListService.EXPECT().GetUserListsByUserID(gomock.Any()).Return(nil, errors.New("errors retrieving userListsByUserID"))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/", userListHandler.GetUserListsByUserID)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/", nil)
	req.Header.Add("USER_ID", "1")

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestUserListHandler_GetUserListsByUserID_No_UserID_Header_Present(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))

	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/", userListHandler.GetUserListsByUserID)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestUserListHandler_GetUserListsByUserID_Invalid_User_ID(t *testing.T) {

	userListService := NewMockIUserListService(gomock.NewController(t))
	userListHandler := NewUserListHandler(userListService)

	gin.SetMode(gin.TestMode)
	c := gin.Default()

	v1 := c.Group("/v1/userLists")
	{
		v1.GET("/", userListHandler.GetUserListsByUserID)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/userLists/", nil)
	req.Header.Add("USER_ID", "invalidID")

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func GetValidUserList() models.UserList {
	return models.UserList{
		Model:  gorm.Model{},
		ListID: 1,
		UserID: 1,
	}
}
