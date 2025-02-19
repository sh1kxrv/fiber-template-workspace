package utils

import (
	"shared/enum"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Role enum.Role `json:"role"`
	ID   string    `json:"id"`
	jwt.RegisteredClaims
}

func ParseJwtToken(token string, secret string) (JwtClaims, error) {
	claims := JwtClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return claims, err
}
