package service

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
)

type CartService struct {
	Repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{
		Repo: repo,
	}
}

func (s *CartService) CreateCart(cart *model.Cart) error {
	return s.Repo.Create(cart)
}
