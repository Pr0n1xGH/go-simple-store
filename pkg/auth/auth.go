package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte(os.Getenv("ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

func CreateTokens(userID uint) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.AccessExp = time.Now().Add(15 * time.Minute).Unix()
	td.RefreshExp = time.Now().Add(7 * 24 * time.Hour).Unix()

	accessClaims := jwt.MapClaims{}
	accessClaims["user_id"] = userID
	accessClaims["exp"] = td.AccessExp

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	at, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return nil, err
	}
	td.AccessToken = at

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userID
	refreshClaims["exp"] = td.RefreshExp

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}
	td.RefreshToken = rt

	return td, nil
}

func VerifyToken(tokenStr string, isRefresh bool) (*jwt.Token, error) {
	var secret []byte
	if isRefresh {
		secret = refreshSecret
	} else {
		secret = accessSecret
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("некорректный метод подписи")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractUserID(token *jwt.Token) (uint, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("недействительный токен")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id отсутствует")
	}

	return uint(userIDFloat), nil
}
