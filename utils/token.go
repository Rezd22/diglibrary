package utils

import (
	"fmt"
	"paa/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret")

func GenerateToken(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	// parse token
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		//check signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid tokennn")
	}

	return claims, nil
}
