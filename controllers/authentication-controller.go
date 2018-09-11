package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	model "github.com/mccoins/ToDo-List/models"
)

// GetToken obtains a jwt token for authentication
func GetToken(w http.ResponseWriter, req *http.Request) {
	creds := model.Credentials{}
	t := model.JWTToken{}

	vars := mux.Vars(req)

	creds.Username = vars["username"]
	creds.Password = vars["password"]

	if len(vars) < 1 {
		json.NewEncoder(w).Encode(model.ResponseCode{ID: "400", Description: "Invalid Request"})
	} else {
		_ = json.NewDecoder(req.Body).Decode(&creds)
		t.Token = t.GenerateToken(creds.Username, creds.Password)

		json.NewEncoder(w).Encode(t)
	}
}

// Authenticated parses the header for the auth token and does simple validation (nothing fancy) and returns true or false
func Authenticated(req *http.Request) bool {
	authorizationHeader := req.Header.Get("authorization")
	isValid := false // assume false
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil {
				return false
			}
			if token.Valid {
				return true
			}
		}
	}
	return isValid
}
