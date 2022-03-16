package handler

//go:generate mockgen -source=users.go -destination users_mock.go -package handler
import (
	"SuperListsAPI/cmd/auth/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/mail"
	_ "net/mail"
)

type IUserService interface {
	GetUser(email string) (*models.User, error)
}

type UserHandler struct {
	userService IUserService
}

func NewUserHandler(userService IUserService) UserHandler {
	return UserHandler{userService: userService}
}

func (uh *UserHandler) Get(c *gin.Context) {
	userEmail := c.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing user email",
		})
		c.Abort()
		return
	}

	user, err := uh.userService.GetUser(userEmail)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
	return

}
