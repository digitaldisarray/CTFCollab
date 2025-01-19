package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: Switch to https://github.com/labstack/echo-jwt for jwt if it cleans up the code?

// Secret key for signing JWTs
var secretKey = []byte("your-secret-key") // TODO: Get from .env

// Claims structure
type Claims struct {
	Username string   `json:"name,omitempty"`
	Ctfs     []string `json:"ctfs,omitempty"`    // Only anonymous sessions have this
	UserId   int      `json:"user_id,omitempty"` // Only authenticated sessions have this
	IsAdmin  bool     `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

// Create a JWT for an anonymous session
func CreateAnonymousSession(username string, ctfs []string) (string, error) {
	claims := Claims{
		Username: username,
		Ctfs:     ctfs,
		IsAdmin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Create a JWT for an authenticated session
func CreateAuthedSession(username string, userid int, isAdmin bool) (string, error) {
	claims := Claims{
		Username: username,
		UserId:   userid,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Parse and validate JWT
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract and return claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
