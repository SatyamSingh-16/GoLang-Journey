package utils

import (
	"student-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, email string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"email":  email,
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(config.JWTSecret)
}
