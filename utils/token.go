package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtKey = []byte("ice_sparkhire")

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(expireDuration time.Duration, id int64) (string, error) {
	expireTime := time.Now().Add(expireDuration)

	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ice_sparkhire",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token signing error: %s", err.Error())
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token, please login again")
	}
	return claims, nil
}
