package auth

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TODO: (for all middleware) Make sure that the user referenced by the token still exists/account hasn't been deleted
// TODO: If a user is stripped of admin, they may still have a JWT that gives them admin, so maybe we can add a last modified field to users?
// Makes sure user is an admin

// AdminOnly
// @Summary Checks admin user credentials
// @Description 
// @Tags user
// @Accept json
// @Produce json
// @Param phrase path string true ""
// @Param user body db.CreateAdminParams true "Admin details"
// @Success 200 {string} string "Admin check passed"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/{phrase}/ [post]
func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get claims
		token := c.Get("user").(*jwt.Token)
		user, ok := token.Claims.(*CustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
		}

		if !user.IsAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
		}
		return next(c)
	}
}

// Makes sure user is admin OR their jwt username matches the /:username in the request

// SelfOnly
// @Summary Checks regular user credentials
// @Description 
// @Tags user
// @Accept json
// @Produce json
// @Param phrase path string true ""
// @Param user body db.CreateUserParams true "User details"
// @Success 200 {string} string "User check passed"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/{phrase}/ [post]
func SelfOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get claims
		token := c.Get("user").(*jwt.Token)
		user, ok := token.Claims.(*CustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
		}

		// Check permissions
		if !user.IsAdmin && user.Name != c.Param("username") {
			return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
		}

		return next(c)
	}
}

// Make sure user is admin OR they are the owner of the CTF in /:phrase

// OwnerOnly
// @Summary Checks admin user credentials
// @Description 
// @Tags user
// @Accept json
// @Produce json
// @Param phrase path string true ""
// @Param user body db.CreateAdminParams true "Admin details"
// @Success 200 {string} string "Admin check passed"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/{phrase}/ [post]
func OwnerOnly(h *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get claims
			token := c.Get("user").(*jwt.Token)
			user, ok := token.Claims.(*CustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
			}

			// Get the CTF by phrase
			phrase := c.Param("phrase")
			ctf, err := h.GetCTFByPhrase(context.Background(), phrase)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CTF")
			}

			// Check if the user is the owner of the CTF
			if ctf.AuthorID != int32(user.Id) {
				return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
			}

			return next(c)
		}
	}
}

// Make sure user is admin OR they are a member of the CTF in /:phrase

// MemberOnly
// @Summary Checks member credentials
// @Description 
// @Tags user
// @Accept json
// @Produce json
// @Param phrase path string true ""
// @Param user body db.CreateUserParams true "Admin details"
// @Success 200 {string} string "Member check passed"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/{phrase}/ [post]
// Makes sure user is an admin
func MemberOnly(h *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get claims
			token := c.Get("user").(*jwt.Token)
			user, ok := token.Claims.(*CustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
			}

			// Check if user is a member of the given CTF
			isMember, err := h.IsUserMemberOfCTF(context.Background(), db.IsUserMemberOfCTFParams{
				UserID: int32(user.Id),
				Phrase: c.Param("phrase"),
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to check membership")
			}
			if !isMember {
				return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
			}

			return next(c)
		}
	}
}
