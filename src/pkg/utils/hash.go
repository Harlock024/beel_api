package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashToken(token string) string {
	hashedToken := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hashedToken[:])
}

func CompareToken(hashedToken, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(token))
	return err == nil
}
