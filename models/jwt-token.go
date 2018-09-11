package models

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
)

// JWTToken is the token used for authentication
type JWTToken struct {
	Token string `json:"Token"`
}

// GenerateToken generates a token based on username and password
func (tok JWTToken) GenerateToken(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
	})

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
		log.Error(error)
	}

	return tokenString
}
