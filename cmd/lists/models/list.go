package models

import (
	"SuperListsAPI/cmd/listItems/models"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Name          string            `json:"name" validate:"required"`
	Description   string            `json:"description" validate:"required"`
	InviteCode    string            `json:"invite_code"`
	UserCreatorID uint              `json:"user_creator_id"`
	ListItems     []models.ListItem `json:"list_items" gorm:"-"`
}

type ListJoinRequest struct {
	InviteCode    string `json:"invite_code" validate:"required"`
	UserInvitedId uint   `json:"user_invited_id" validate:"required"`
}
