package handler

import (
	"SuperListsAPI/cmd/lists/models"
	userListsHandler "SuperListsAPI/cmd/userLists/handler"
	userListsModel "SuperListsAPI/cmd/userLists/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=lists.go -destination lists_mock.go -package handler

type IListService interface {
	Create(list models.List) (*models.List, error)
	GetLists(userId string) (*[]models.List, error)
	Get(listId string) (*models.List, error)
	Update(list models.List) (*models.List, error)
	Delete(idsToDelete []uint) (*[]uint, error)
}

type IUserListService interface {
	Create(list userListsModel.UserList) (*userListsModel.UserList, error)
	Get(userListID string) (*userListsModel.UserList, error)
	Delete(userListID string) (*int, error)
	GetUserListsByUserID(userId string) (*[]userListsModel.UserList, error)
	GetUserListsByListID(listID string) (*[]userListsModel.UserList, error)
}

type ListHandler struct {
	listService      IListService
	userListsService userListsHandler.IUserListService
}

func NewListHandler(service IListService, userListService userListsHandler.IUserListService) ListHandler {
	return ListHandler{listService: service, userListsService: userListService}
}

func (lh *ListHandler) Create(c *gin.Context) {

	list := models.List{}

	validate := validator.New()
	err := c.ShouldBindJSON(&list)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}
	err = validate.Struct(list)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result, err := lh.listService.Create(list)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	userList := userListsModel.UserList{
		ListID: result.ID,
		UserID: result.UserCreatorID,
	}

	_, err = lh.userListsService.Create(userList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
	return

}

func (lh *ListHandler) GetLists(c *gin.Context) {

	userID := c.Request.Header.Get("USER_ID")

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

	lists, err := lh.listService.GetLists(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if lists == nil || len(*lists) < 1 {
		c.JSON(http.StatusNoContent, lists)
		return
	}

	c.JSON(http.StatusOK, lists)
	return

}

func (lh *ListHandler) Get(c *gin.Context) {
	listID := c.Param("id")

	if _, err := strconv.Atoi(listID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
		})
		c.Abort()
		return
	}

	list, err := lh.listService.Get(listID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if list == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("List with id %s not found", listID))
		return
	}

	c.JSON(http.StatusOK, list)
	return
}

func (lh *ListHandler) Update(c *gin.Context) {

	listUpdateRequest := models.List{}

	listID := c.Param("id")

	err := c.ShouldBindJSON(&listUpdateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	if _, err := strconv.Atoi(listID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
		})
		c.Abort()
		return
	}

	if listID == "" || fmt.Sprint(listUpdateRequest.ID) != listID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing list id on request path or list id mistmatch",
		})
		c.Abort()
		return
	}

	validate := validator.New()

	err = validate.Struct(listUpdateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	list, err := lh.listService.Update(listUpdateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if list == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("List with id %d not found", listUpdateRequest.ID))
		return
	}

	c.JSON(http.StatusOK, list)
	return
}

func (lh *ListHandler) Delete(c *gin.Context) {
	listID := c.Param("id")
	userID := c.Request.Header.Get("USER_ID")
	var idsToDelete []uint

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing user id on request header",
		})
		c.Abort()
		return
	}

	if _, err := strconv.Atoi(listID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
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
	parsedUserID, _ := strconv.Atoi(userID)
	userListsByListID, err := lh.userListsService.GetUserListsByListID(listID)

	list, err := lh.listService.Get(listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	//TODO mejorar esto, pasar a la logica de servicio
	//Si es el dueño, delete all the userLists sino a la unica que tiene
	if list.UserCreatorID == uint(parsedUserID) {
		idsToDelete = UserListsToDelete(*userListsByListID, parsedUserID, true)
	} else {
		idsToDelete = UserListsToDelete(*userListsByListID, parsedUserID, false)
	}

	deletedIDs, err := lh.listService.Delete(idsToDelete)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if deletedIDs == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("List with id %s not found", listID))
		return
	}

	c.JSON(http.StatusOK, deletedIDs)
	return
}

func (lh *ListHandler) JoinList(c *gin.Context) {

	listID := c.Param("listID")
	userID := c.Request.Header.Get("USER_ID")

	if _, err := strconv.Atoi(listID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
		})
		c.Abort()
		return
	}

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing user id on request header",
		})
		c.Abort()
		return
	}

	parsedListID, _ := strconv.Atoi(listID)
	parsedUserID, _ := strconv.Atoi(userID)

	userList := userListsModel.UserList{
		ListID: uint(parsedListID),
		UserID: uint(parsedUserID),
	}
	ul, err := lh.userListsService.Create(userList)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if ul == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("UserList with id %s could not be created", listID))
		return
	}

	c.JSON(http.StatusOK, ul)
	return
}

func UserListsToDelete(userListsRecovered []userListsModel.UserList, userID int, isOwner bool) []uint {

	var idListToDelete []uint

	if !isOwner {
		for _, ul := range userListsRecovered {
			if ul.UserID == uint(userID) {
				idListToDelete = append(idListToDelete, ul.ID)
			}
		}
		return idListToDelete
	}

	for _, ul := range userListsRecovered {
		idListToDelete = append(idListToDelete, ul.ID)
	}
	return idListToDelete
}
