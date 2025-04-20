package repository

import (
	"go-start/pkg/model"

	"gorm.io/gorm"
)

type CartItemRepository struct {
	DB *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{DB: db}
}

func (r *CartItemRepository) Create(cartItem *model.CartItem) error {
	return r.DB.Create(cartItem).Error
}

func (r *CartItemRepository) Delete(id uint) error {
	return r.DB.Delete(&model.CartItem{}, id).Error
}

func (r *CartItemRepository) GetByCartID(cartID uint) ([]model.CartItem, error) {
	var items []model.CartItem
	err := r.DB.Where("cart_id = ?", cartID).Find(&items).Error
	return items, err
}

func (r *CartItemRepository) UpdateQuantity(id uint, quantity int) error {
	return r.DB.Model(&model.CartItem{}).Where("id = ?", id).Update("quantity", quantity).Error
}
