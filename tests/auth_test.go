package tests

import (
	"testing"

	m "github.com/mccoins/ToDo-List/models"
)

// TestJWTToken tests whether a valid token is generated
func TestJWTToken(t *testing.T) {
	var token = m.JWTToken{}.GenerateToken("user1", "admin")

	valid := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6ImFkbWluIiwidXNlcm5hbWUiOiJ1c2VyMSJ9.r5qF4wU0k6v5LD6CibWpWaU0DS5nz5N39bin8Kw-_9U"

	if token != valid {
		t.Error("Invalid token", token)
	}
}
