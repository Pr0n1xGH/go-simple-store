package repository

import (
	"go-start/pkg/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) Create(cart *model.Cart) error {
	return r.DB.Create(cart).Error
}

func (r *CartRepository) UpdateCartUser(id uint, userID uint) error {
	return r.DB.Model(&model.Cart{}).Where("id = ?", id).Update("user_id", userID).Error
}
