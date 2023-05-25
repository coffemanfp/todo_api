package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, lifespan int, secretKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_id"] = id
	if lifespan != 0 {
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
