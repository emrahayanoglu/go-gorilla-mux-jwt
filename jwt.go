package main

import "github.com/dgrijalva/jwt-go"

func CreateToken() (string, error) {
	signingKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "hello",
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
