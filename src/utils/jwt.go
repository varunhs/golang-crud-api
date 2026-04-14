package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("secret")

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email" : email,
		"exp" : time.Now().Add(time.Minute * 2).Unix(),
	}

	token  := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}