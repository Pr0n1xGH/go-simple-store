package service

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
)

type CartItemService struct {
	Repo *repository.CartItemRepository
}

func NewCartItemService(repo *repository.CartItemRepository) *CartItemService {
	return &CartItemService{Repo: repo}
}

func (s *CartItemService) AddItem(cartID uint, item *model.CartItem) error {
	item.CartID = cartID
	return s.Repo.Create(item)
}

func (s *CartItemService) RemoveItem(itemID uint) error {
	return s.Repo.Delete(itemID)
}

func (s *CartItemService) GetItems(cartID uint) ([]model.CartItem, error) {
	return s.Repo.GetByCartID(cartID)
}

func (s *CartItemService) UpdateItemQuantity(itemID uint, quantity int) error {
	return s.Repo.UpdateQuantity(itemID, quantity)
}
