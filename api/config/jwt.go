package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenTTL = time.Hour * 24 * 30

type Claims struct {
	UserID uint `json:"uid"`
	jwt.RegisteredClaims
}

func jwtSecret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "dev-insecure-jwt-secret"
	}

	return []byte(s)
}

func GenerateJWT(userId uint) (string, error) {
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret())
}
