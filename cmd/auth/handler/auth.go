package handler

import (
	"SuperListsAPI/cmd/auth/models"
	"SuperListsAPI/cmd/auth/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//go:generate mockgen -source=auth.go -destination auth_mock.go -package handler

type IAuthService interface {
	Login(payload models.LoginPayload) (*string, error)
	SignUp(userRequest *models.User) (*models.User, error)
}

type AuthHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) AuthHandler {
	return AuthHandler{authService: authService}
}

func (authHandler *AuthHandler) Login(c *gin.Context) {
	var payload models.LoginPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	token, err := authHandler.authService.Login(payload)

	if err != nil {
		switch err.Error() {
		case repository.INVALID_PASSWORD:
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		case repository.EMAIL_NOT_FOUND:
			c.JSON(http.StatusNotFound, err.Error())
			return
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.Header("token", *token)
	c.JSON(http.StatusOK, token)
	return
}

func (authHandler *AuthHandler) SignUp(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	newUser, err := authHandler.authService.SignUp(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, newUser)
	return

}
