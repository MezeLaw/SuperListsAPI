package handler

import (
	"SuperListsAPI/cmd/auth/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		userService IUserService
	}
	tests := []struct {
		name string
		args args
		want UserHandler
	}{
		{
			name: "Test with nil service should pass",
			args: args{nil},
			want: NewUserHandler(nil),
		},
		{
			name: "Test with no nil service should pass",
			args: args{NewMockIUserService(gomock.NewController(t))},
			want: NewUserHandler(NewMockIUserService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_Get(t *testing.T) {
	user := GetValidUser()
	service := NewMockIUserService(gomock.NewController(t))
	service.EXPECT().GetUser(gomock.Any()).Return(&user, nil)

	handler := NewUserHandler(service)

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users/:email", handler.Get)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1/users/mezequielabogado@gmail.com", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestUserHandler_Get_Error_From_Service(t *testing.T) {
	service := NewMockIUserService(gomock.NewController(t))
	service.EXPECT().GetUser(gomock.Any()).Return(nil, errors.New("error from service"))

	handler := NewUserHandler(service)

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users/:email", handler.Get)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1/users/mezequielabogado@gmail.com", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)

}

func TestUserHandler_Get_Missing_Email_On_Path(t *testing.T) {
	service := NewMockIUserService(gomock.NewController(t))

	handler := NewUserHandler(service)

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/users/:email", handler.Get)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1/users/1", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)

}

func GetValidUser() models.User {
	return models.User{
		Name:     "Meze Test",
		Email:    "mezetest@gmail.com",
		Password: "password",
		Role:     "CAPO DI TUTTI CAPI",
	}
}
