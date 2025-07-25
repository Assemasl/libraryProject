package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtKey = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "secret" // default fallback
	}
	return secret
}

func GenerateJWT(userID uint, isAuthor bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"is_author": isAuthor,
		"exp":       time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
