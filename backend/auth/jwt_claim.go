package auth

import "github.com/golang-jwt/jwt/v5"

// JWT custom claims structure
type CustomClaims struct {
	Name     string `json:"name"`
	LoggedIn bool   `json:"logged_in"` // If false, it is a guest session
	Id       int    `json:"id"`        // Either session ID or user ID
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}
