package handler

import (
	"SuperListsAPI/cmd/auth/models"
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

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		authService IAuthService
	}
	tests := []struct {
		name string
		args args
		want AuthHandler
	}{
		{
			name: "Test with nil service should create a handler",
			args: args{authService: nil},
			want: NewAuthHandler(nil),
		},
		{
			name: "Test with no nil service should create a handler with no nil service",
			args: args{authService: NewMockIAuthService(gomock.NewController(t))},
			want: NewAuthHandler(NewMockIAuthService(gomock.NewController(t))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHandler(tt.args.authService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthHandler_Login_Ok(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	validToken := "unTokenValido"

	authService.EXPECT().Login(gomock.Any()).Return(&validToken, nil)
	resp := httptest.NewRecorder()

	validLoginPayload := GetValidLoginPayload()

	jsonDto, _ := json.Marshal(&validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

}

func TestAuthHandler_Login_Bad_Request(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	resp := httptest.NewRecorder()

	validLoginPayload := 1

	jsonDto, _ := json.Marshal(validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)

}

func TestAuthHandler_Login_Invalid_Required_JsonValues(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	resp := httptest.NewRecorder()

	validLoginPayload := map[string]interface{}{
		"email": "emailvalid@gmail.com",
	}

	jsonDto, _ := json.Marshal(validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)

}

func TestAuthHandler_Login_Invalid_Credentials(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	authService.EXPECT().Login(gomock.Any()).Return(nil, errors.New("invalid credentials"))
	resp := httptest.NewRecorder()

	validLoginPayload := GetValidLoginPayload()

	jsonDto, _ := json.Marshal(validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusUnauthorized)

}

func TestAuthHandler_Login_Email_Not_Found(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	authService.EXPECT().Login(gomock.Any()).Return(nil, errors.New("email not found"))
	resp := httptest.NewRecorder()

	validLoginPayload := GetValidLoginPayload()

	jsonDto, _ := json.Marshal(validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusNotFound)

}

func TestAuthHandler_Login_Default_Error(t *testing.T) {

	authService := NewMockIAuthService(gomock.NewController(t))
	authHandler := NewAuthHandler(authService)

	authService.EXPECT().Login(gomock.Any()).Return(nil, errors.New("other error should return 500"))
	resp := httptest.NewRecorder()

	validLoginPayload := GetValidLoginPayload()

	jsonDto, _ := json.Marshal(validLoginPayload)

	req, _ := http.NewRequest(http.MethodGet, "/v1/auth/login", strings.NewReader(string(jsonDto)))

	gin.SetMode(gin.TestMode)

	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.GET("/login", authHandler.Login)
	}

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusInternalServerError)

}

func TestAuthHandler_SignUp(t *testing.T) {

	validUser := GetValidUser()

	authService := NewMockIAuthService(gomock.NewController(t))
	authService.EXPECT().SignUp(gomock.Any()).Return(validUser, nil)

	authHandler := NewAuthHandler(authService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.POST("/signup", authHandler.SignUp)
	}

	jsonDto, _ := json.Marshal(validUser)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/signup", strings.NewReader(string(jsonDto)))

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusCreated)

}

func TestAuthHandler_SignUp_Internal_Server_Error(t *testing.T) {

	validUser := GetValidUser()

	authService := NewMockIAuthService(gomock.NewController(t))
	authService.EXPECT().SignUp(gomock.Any()).Return(nil, errors.New("internal sv error"))

	authHandler := NewAuthHandler(authService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.POST("/signup", authHandler.SignUp)
	}

	jsonDto, _ := json.Marshal(validUser)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/signup", strings.NewReader(string(jsonDto)))

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusInternalServerError)

}

func TestAuthHandler_SignUp_Invalid_Json(t *testing.T) {

	invalidUser := 1

	authService := NewMockIAuthService(gomock.NewController(t))

	authHandler := NewAuthHandler(authService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.POST("/signup", authHandler.SignUp)
	}

	jsonDto, _ := json.Marshal(invalidUser)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/signup", strings.NewReader(string(jsonDto)))

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)

}

func TestAuthHandler_SignUp_Invalid_Json_On_Required_Attr(t *testing.T) {

	invalidUser := map[string]interface{}{
		"email": "emailvalido@gmail.com",
	}

	authService := NewMockIAuthService(gomock.NewController(t))

	authHandler := NewAuthHandler(authService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	v1 := router.Group("/v1/auth/")
	{
		v1.POST("/signup", authHandler.SignUp)
	}

	jsonDto, _ := json.Marshal(invalidUser)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/signup", strings.NewReader(string(jsonDto)))

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusBadRequest)

}

func GetValidLoginPayload() *models.LoginPayload {
	return &models.LoginPayload{
		Email:    "meze@meze.com",
		Password: "password",
	}
}

func GetValidUser() *models.User {
	return &models.User{
		Name:     "Meze",
		Email:    "meze@meze.com",
		Password: "password",
		Role:     "ADMIN",
	}
}
