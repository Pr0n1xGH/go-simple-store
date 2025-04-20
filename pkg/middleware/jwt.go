package middleware

import (
	"context"
	"go-start/pkg/auth"
	"go-start/pkg/handler"
	"net/http"
	"strings"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handler.WriteJSONError(w, "Отсутствует access токен", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := auth.VerifyToken(tokenStr, false)
		if err != nil {
			refreshToken := r.Header.Get("X-Refresh-Token")
			if refreshToken == "" {
				handler.WriteJSONError(w, "Истёк access токен, и нет refresh токена", http.StatusUnauthorized)
				return
			}

			refreshParsed, err := auth.VerifyToken(refreshToken, true)
			if err != nil || !refreshParsed.Valid {
				handler.WriteJSONError(w, "Refresh токен недействителен", http.StatusUnauthorized)
				return
			}

			userID, err := auth.ExtractUserID(refreshParsed)
			if err != nil {
				handler.WriteJSONError(w, "Не удалось извлечь user_id из refresh токена", http.StatusUnauthorized)
				return
			}

			newTokens, err := auth.CreateTokens(userID)
			if err != nil {
				handler.WriteJSONError(w, "Ошибка при обновлении токенов", http.StatusInternalServerError)
				return
			}

			w.Header().Set("X-New-Access-Token", newTokens.AccessToken)

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		userID, err := auth.ExtractUserID(token)
		if err != nil {
			handler.WriteJSONError(w, "Неверный access токен", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (uint, bool) {
	userID, ok := r.Context().Value(userIDKey).(uint)
	return userID, ok
}
