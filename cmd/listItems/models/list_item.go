package models

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	ListID      int    `json:"list_id" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDone      bool   `json:"is_done"`
}
