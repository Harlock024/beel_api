package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetAccessSecretKey() []byte {
	return []byte(os.Getenv("ACCESS_SECRET_KEY"))
}

func GetRefreshSecretKey() []byte {
	return []byte(os.Getenv("REFRESH_SECRET_KEY"))
}

func GenerateTokens(username string, userId uuid.UUID) (string, string, error) {
	accessToken, err := CreateAccessToken(username, userId)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := CreateRefreshToken(username, userId)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func CreateRefreshToken(username string, userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  userId,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	tokenString, err := token.SignedString(GetRefreshSecretKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateAccessToken(username string, userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  userId,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(GetAccessSecretKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
