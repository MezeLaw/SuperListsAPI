package handler

import (
	"SuperListsAPI/cmd/lists/models"
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
	Delete(listId string) (*int, error)
}

type ListHandler struct {
	listService IListService
}

func NewListHandler(service IListService) ListHandler {
	return ListHandler{listService: service}
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

	if _, err := strconv.Atoi(listID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list id",
		})
		c.Abort()
		return
	}

	deletedListID, err := lh.listService.Delete(listID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if deletedListID == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("List with id %s not found", listID))
		return
	}

	c.JSON(http.StatusOK, deletedListID)
	return
}
