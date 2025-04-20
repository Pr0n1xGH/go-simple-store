package handler

import (
	"encoding/json"
	"go-start/pkg/model"
	"go-start/pkg/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CartItemHandler struct {
	Service *service.CartItemService
}

func NewCartItemHandler(s *service.CartItemService) *CartItemHandler {
	return &CartItemHandler{Service: s}
}

func (h *CartItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item model.CartItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	cartIDStr := chi.URLParam(r, "cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		http.Error(w, "Неверный cartID", http.StatusBadRequest)
		return
	}

	if err := h.Service.AddItem(uint(cartID), &item); err != nil {
		http.Error(w, "Ошибка при добавлении", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *CartItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	cartIDStr := chi.URLParam(r, "cartID")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		http.Error(w, "Неверный cartID", http.StatusBadRequest)
		return
	}

	items, err := h.Service.GetItems(uint(cartID))
	if err != nil {
		http.Error(w, "Ошибка при получении", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func (h *CartItemHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, "Неверный itemID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Quantity int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateItemQuantity(uint(itemID), payload.Quantity); err != nil {
		http.Error(w, "Ошибка при обновлении", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Количество обновлено"})
}

func (h *CartItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, "Неверный itemID", http.StatusBadRequest)
		return
	}

	if err := h.Service.RemoveItem(uint(itemID)); err != nil {
		http.Error(w, "Ошибка при удалении", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Товар удалён из корзины"})
}
