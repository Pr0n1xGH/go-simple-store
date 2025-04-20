package handler

import (
	"encoding/json"
	"go-start/pkg/model"
	"go-start/pkg/service"
	"net/http"
)

type CartHandler struct {
	Service *service.CartService
}

func NewCartHandler(s *service.CartService) *CartHandler {
	return &CartHandler{
		Service: s,
	}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart model.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		WriteJSONError(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateCart(&cart); err != nil {
		WriteJSONError(w, "Ошибка при создании", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
}
