package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJwt(name string) string {
	secret := []byte("key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
	})
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func ParseJwt(token string) string {
	fmt.Println("解析Token", token)
	secret := []byte("key")
	if token == "" {
		return ""
	}
	realToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	fmt.Println("真实的Token", realToken)
	claims, _ := realToken.Claims.(jwt.Claims)

	fmt.Println("claims", claims)
	if claims != nil {
		return "解析成功"
	} else {
		return ""
	}
}
