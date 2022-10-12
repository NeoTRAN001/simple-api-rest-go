package utils

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"platzi.com/go/rest-ws/models"
	"platzi.com/go/rest-ws/server"
)

func Validate(s server.Server, a string) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(a)

	return jwt.ParseWithClaims(
		tokenString, &models.AppClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		},
	)
}
