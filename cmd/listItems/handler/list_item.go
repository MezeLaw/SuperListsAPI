package handler

import (
	"SuperListsAPI/cmd/listItems/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=list_item.go -destination list_item_mock.go -package handler

type IListItemService interface {
	Create(item models.ListItem) (*models.ListItem, error)
	Get(listItemID string) (*models.ListItem, error)
	Update(item models.ListItem) (*models.ListItem, error)
	Delete(listItemID string) (*int, error)
	GetItemsListByListID(listId string) (*[]models.ListItem, error)
	DeleteListItemsByListID(listId string) (*int, error)
	BulkDelete(tasksToDelete []models.ListItem) (*int, error)
	MarkAsCompleted(tasksToDelete []models.ListItem) (*int, error)
	MarkAsPending(tasksToDelete []models.ListItem) (*int, error)
}

type ListItemHandler struct {
	listItemService IListItemService
}

func NewListItemHandler(service IListItemService) ListItemHandler {
	return ListItemHandler{service}
}

func (lih *ListItemHandler) Create(c *gin.Context) {
	listItem := models.ListItem{}

	validate := validator.New()
	err := c.ShouldBindJSON(&listItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}
	err = validate.Struct(listItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.Create(listItem)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}

func (lih *ListItemHandler) Get(c *gin.Context) {
	listItemID := c.Param("id")

	if _, err := strconv.Atoi(listItemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid listItem id",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.Get(listItemID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	if result == nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("ListItem with id %s not found", listItemID))
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (lih *ListItemHandler) Update(c *gin.Context) {
	listItemUpdateRequest := models.ListItem{}

	listItemID := c.Param("id")

	err := c.ShouldBindJSON(&listItemUpdateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	if _, err := strconv.Atoi(listItemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list item id",
		})
		c.Abort()
		return
	}

	if listItemID == "" || fmt.Sprint(listItemUpdateRequest.ID) != listItemID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing list item id on request path or list item id mistmatch",
		})
		c.Abort()
		return
	}

	validate := validator.New()

	err = validate.Struct(listItemUpdateRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.Update(listItemUpdateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (lih *ListItemHandler) Delete(c *gin.Context) {
	listItemID := c.Param("id")

	if _, err := strconv.Atoi(listItemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid list item id",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.Delete(listItemID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return

}

func (lih *ListItemHandler) BulkDelete(c *gin.Context) {
	var listItemsToDelete []models.ListItem

	err := c.ShouldBindJSON(&listItemsToDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	if len(listItemsToDelete) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "no tasks received for delete",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.BulkDelete(listItemsToDelete)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return

}

func (lih *ListItemHandler) MarkAsCompleted(c *gin.Context) {
	var listItemsToUpdate []models.ListItem

	err := c.ShouldBindJSON(&listItemsToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	if len(listItemsToUpdate) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "no tasks received for delete",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.MarkAsCompleted(listItemsToUpdate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func (lih *ListItemHandler) MarkAsPending(c *gin.Context) {
	var listItemsToUpdate []models.ListItem

	err := c.ShouldBindJSON(&listItemsToUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	if len(listItemsToUpdate) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "no tasks received for delete",
		})
		c.Abort()
		return
	}

	result, err := lih.listItemService.MarkAsPending(listItemsToUpdate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}
