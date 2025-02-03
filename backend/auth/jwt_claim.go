package auth

import "github.com/golang-jwt/jwt/v5"

// JWT custom claims structure
type CustomClaims struct {
	Name     string `json:"name,omitempty"`
	LoggedIn bool   `json:"logged_in,omitempty"` // If false, it is a guest session
	Id       int    `json:"user_id,omitempty"`   // Either session ID or user ID
	IsAdmin  bool   `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

// // Create a JWT for an anonymous session
// func CreateAnonymousSession(username string, ctfs []string) (string, error) {
// 	claims := Claims{
// 		Username: username,
// 		Ctfs:     ctfs,
// 		IsAdmin:  false,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(secretKey)
// }

// // Create a JWT for an authenticated session
// func CreateAuthedSession(username string, userid int, isAdmin bool) (string, error) {
// 	claims := Claims{
// 		Username: username,
// 		UserId:   userid,
// 		IsAdmin:  isAdmin,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(secretKey)
// }

// // Parse and validate JWT
// func ParseToken(tokenString string) (*Claims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		// Ensure the signing method matches
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("invalid signing method")
// 		}
// 		return secretKey, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	// Extract and return claims
// 	claims, ok := token.Claims.(*Claims)
// 	if !ok {
// 		return nil, errors.New("invalid claims")
// 	}

// 	return claims, nil
// }
