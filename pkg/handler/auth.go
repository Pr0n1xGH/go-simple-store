package handler

import (
	"encoding/json"
	"go-start/pkg/auth"
	"go-start/pkg/service"
	"net/http"
)

type AuthHandler struct {
	Service *service.UserService
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func NewAuthHandler(s *service.UserService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, "Неверный JSON", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		WriteJSONError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokens, err := auth.CreateTokens(user.ID)
	if err != nil {
		WriteJSONError(w, "Ошибка генерации токенов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.RefreshToken == "" {
		WriteJSONError(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	token, err := auth.VerifyToken(req.RefreshToken, true) // true = refresh
	if err != nil || !token.Valid {
		WriteJSONError(w, "Недействительный refresh токен", http.StatusUnauthorized)
		return
	}

	userID, err := auth.ExtractUserID(token)
	if err != nil {
		WriteJSONError(w, "Не удалось извлечь user_id", http.StatusUnauthorized)
		return
	}

	tokens, err := auth.CreateTokens(userID)
	if err != nil {
		WriteJSONError(w, "Ошибка при создании токенов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
