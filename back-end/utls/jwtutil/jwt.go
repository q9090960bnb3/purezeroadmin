package jwtutil

import (
	"github.com/golang-jwt/jwt/v4"
)

func GetToken(secretKey string, iat, expired int64, payload map[string]any) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = expired
	claims["iat"] = iat

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
