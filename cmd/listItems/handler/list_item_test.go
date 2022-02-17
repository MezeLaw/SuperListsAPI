package handler

import (
	"SuperListsAPI/cmd/listItems/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewLisItemHandler(t *testing.T) {
	type args struct {
		service IListItemService
	}
	tests := []struct {
		name string
		args args
		want ListItemHandler
	}{
		{
			name: "Test with nil service should pass",
			args: args{nil},
			want: NewListItemHandler(nil),
		},
		{
			name: "Test with no nil service should pass",
			args: args{NewMockIListItemService(gomock.NewController(t))},
			want: NewListItemHandler(NewMockIListItemService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListItemHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLisItemHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListItemHandler_Create(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Create(gomock.Any()).Return(&listItem, nil)

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.POST("/", listItemHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/listItems/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)

}

func TestListItemHandler_Create_Error(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Create(gomock.Any()).Return(nil, errors.New("Error from itemListService "))

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.POST("/", listItemHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/listItems/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

}

func TestListItemHandler_Create_Invalid_Struct(t *testing.T) {
	listItem := map[string]interface{}{
		"title": 1,
	}
	mockedService := NewMockIListItemService(gomock.NewController(t))
	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.POST("/", listItemHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/listItems/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestListItemHandler_Create_Missing_Mandatory_Values(t *testing.T) {
	listItem := map[string]interface{}{
		"title": "titulo",
	}
	mockedService := NewMockIListItemService(gomock.NewController(t))
	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.POST("/", listItemHandler.Create)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/listItems/", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestListItemHandler_Get(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Get(gomock.Any()).Return(&listItem, nil)

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.GET("/:id", listItemHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestListItemHandler_Get_No_Results(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Get(gomock.Any()).Return(nil, nil)

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.GET("/:id", listItemHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)
}

func TestListItemHandler_Get_Error(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Get(gomock.Any()).Return(&listItem, errors.New("error from list item service"))

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.GET("/:id", listItemHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestListItemHandler_Get_Invalid_ID(t *testing.T) {
	listItem := GetValidListItem()
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	jsonDto, _ := json.Marshal(listItem)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.GET("/:id", listItemHandler.Get)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/listItems/a", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestListItemHandler_Delete(t *testing.T) {
	idDeleted := 1
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Delete(gomock.Any()).Return(&idDeleted, nil)

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.DELETE("/:id", listItemHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/listItems/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestListItemHandler_Delete_Invalid_ID(t *testing.T) {
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.DELETE("/:id", listItemHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/listItems/a", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestListItemHandler_Delete_Error(t *testing.T) {

	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Delete(gomock.Any()).Return(nil, errors.New("error from item list service"))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.DELETE("/:id", listItemHandler.Delete)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v1/listItems/1", nil)

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestListItemHandler_Update(t *testing.T) {
	validListItem := GetValidListItem()
	validListItem.ID = 1
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Update(gomock.Any()).Return(&validListItem, nil)

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestListItemHandler_Update_ID_Mismatch(t *testing.T) {
	validListItem := GetValidListItem()
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestListItemHandler_Update_Error(t *testing.T) {
	validListItem := GetValidListItem()
	validListItem.ID = 1
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))
	mockedService.EXPECT().Update(gomock.Any()).Return(nil, errors.New("error from list item service"))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

func TestListItemHandler_Update_Invalid_ID(t *testing.T) {
	validListItem := map[string]interface{}{}
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/a", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestListItemHandler_Update_Cant_Bind_JSON(t *testing.T) {
	validListItem := map[string]interface{}{
		"list_id": "a",
	}
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/a", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestListItemHandler_Update_Invalid_JSON(t *testing.T) {
	validListItem := map[string]interface{}{
		"id":      1,
		"list_id": 1,
	}
	jsonDto, _ := json.Marshal(validListItem)
	mockedService := NewMockIListItemService(gomock.NewController(t))

	listItemHandler := NewListItemHandler(mockedService)

	gin.SetMode(gin.TestMode)

	c := gin.Default()

	v1 := c.Group("/v1/listItems")
	{
		v1.PUT("/:id", listItemHandler.Update)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v1/listItems/1", strings.NewReader(string(jsonDto)))

	c.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func GetValidListItem() models.ListItem {
	return models.ListItem{
		ListID:      1,
		UserID:      1,
		Title:       "titulo",
		Description: "description",
		IsDone:      false,
	}
}
