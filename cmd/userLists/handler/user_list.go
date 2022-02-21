package handler

import (
	"SuperListsAPI/cmd/userLists/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=user_list.go -destination user_lists_mock.go -package handler

type IUserListService interface {
	Create(list models.UserList) (*models.UserList, error)
	Get(userListID string) (*models.UserList, error)
	Delete(userListIDToDelete *[]uint) (*int, error)
	GetUserListsByUserID(userId string) (*[]models.UserList, error)
	GetUserListsByListID(listID string) (*[]models.UserList, error)
}

type UserListHandler struct {
	userListService IUserListService
}

func NewUserListHandler(userListService IUserListService) UserListHandler {
	return UserListHandler{userListService: userListService}
}

func (ulh *UserListHandler) Create(c *gin.Context) {
	var userList models.UserList

	validate := validator.New()
	err := c.ShouldBindJSON(&userList)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}
	err = validate.Struct(userList)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result, err := ulh.userListService.Create(userList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}

func (ulh *UserListHandler) Get(c *gin.Context) {
	userListID := c.Param("id")

	if _, err := strconv.Atoi(userListID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid user list id",
		})
		c.Abort()
		return
	}

	list, err := ulh.userListService.Get(userListID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if list == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("User list with id %s not found", userListID))
		return
	}

	c.JSON(http.StatusOK, list)
	return
}

func (ulh *UserListHandler) Delete(c *gin.Context) {
	userListID := c.Param("id")

	parsedUserListID, err := strconv.Atoi(userListID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
		})
		c.Abort()
		return
	}

	userListToDelete := []uint{uint(parsedUserListID)}

	deletedUserListID, err := ulh.userListService.Delete(&userListToDelete)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if deletedUserListID == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("User list with id %d not found", deletedUserListID))
		return
	}

	c.JSON(http.StatusOK, deletedUserListID)
	return
}

func (ulh *UserListHandler) GetUserListsByUserID(c *gin.Context) {
	userID := c.Request.Header.Get("user_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing user id on request header",
		})
		c.Abort()
		return
	}

	if _, err := strconv.Atoi(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid user id",
		})
		c.Abort()
		return
	}

	userLists, err := ulh.userListService.GetUserListsByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if userLists == nil || len(*userLists) < 1 {
		c.JSON(http.StatusNoContent, userLists)
		return
	}

	c.JSON(http.StatusOK, userLists)
	return

}
