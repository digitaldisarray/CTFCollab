package auth

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TODO: Make JWTs rotate quickly to minimize the chance of them being used after an account gets deleted or loses permissions in some way

// Makes sure user is an admin
func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
func SelfOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		user, ok := token.Claims.(*CustomClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
		}

		if !user.IsAdmin && user.Name != c.Param("username") {
			return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
		}

		return next(c)
	}
}

// Make sure user is admin OR they are the owner of the CTF in /:phrase
func OwnerOnly(h *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			user, ok := token.Claims.(*CustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
			}

			phrase := c.Param("phrase")
			ctf, err := h.GetCTFByPhrase(context.Background(), phrase)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get CTF")
			}

			// Make sure not a guest and check if the user is the owner of the CTF
			if !user.LoggedIn || ctf.AuthorID != int32(user.Id) {
				return echo.NewHTTPError(http.StatusForbidden, "Access forbidden")
			}

			return next(c)
		}
	}
}

// Make sure user is admin OR they are a member of the CTF in /:phrase
func MemberOnly(h *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get claims
			token := c.Get("user").(*jwt.Token)
			user, ok := token.Claims.(*CustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or token claims")
			}

			// Check if user or guest is a member of the given CTF
			var isMember bool
			var err error
			if user.LoggedIn {
				isMember, err = h.IsUserMemberOfCTF(context.Background(), db.IsUserMemberOfCTFParams{
					UserID: int32(user.Id),
					Phrase: c.Param("phrase"),
				})
			} else {
				isMember, err = h.IsGuestMemberOfCTF(context.Background(), db.IsGuestMemberOfCTFParams{
					ID:     int32(user.Id),
					Phrase: c.Param("phrase"),
				})
			}
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to check membership")
			}
			if !isMember {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}

			return next(c)
		}
	}
}
