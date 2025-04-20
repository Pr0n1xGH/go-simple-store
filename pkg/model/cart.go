package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
}
