package models

import "gorm.io/gorm"

type UserList struct {
	gorm.Model
	ListID uint `json:"list_id" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
}
