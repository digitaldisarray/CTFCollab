package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/digitaldisarray/ctfcollab/auth"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserAuthParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (h *Handler) CreateUser(c echo.Context) error {
	// Parse request
	user := new(UserAuthParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure username isn't taken
	ctx := context.Background()
	_, err := h.Queries.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	// TODO: update for more passes, currently does only 1
	// Hash password
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Copy over data into a new user struct
	newUser := new(db.CreateUserParams)
	newUser.PasswordHash = hash
	newUser.Username = user.Username

	// Insert data into db
	result, err := h.Queries.CreateUser(ctx, *newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure user was created
	if _, err = VerifyParseResult(result, 1); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetUser(c echo.Context) error {
	username := c.Param("username")
	ctx := context.Background()

	// Get user from database
	user, err := h.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Remove password hash before sending user data
	user.PasswordHash = ""
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) LoginUser(c echo.Context) error {
	user := new(UserAuthParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Lookup user in database
	ctx := context.Background()
	dbu, err := h.Queries.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure passwords match
	res, err := argon2id.ComparePasswordAndHash(user.Password, dbu.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !res {
		return c.JSON(http.StatusBadRequest, "Invalid login credentials.")
	}

	// Create JWT token
	claims := &auth.CustomClaims{
		Name:     user.Username,
		LoggedIn: true,
		Id:       int(dbu.ID),
		IsAdmin:  dbu.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("change_me")) // TODO: get from .env
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h *Handler) ChangePassword(c echo.Context) error {
	var req ChangePasswordParams
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// TODO: Bypass old password check if JWT is admin

	// Make sure old password is correct
	ctx := context.Background()
	dbu, err := h.Queries.GetUserByUsername(ctx, c.Param("username")) // Lookup user in database
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure passwords match
	res, err := argon2id.ComparePasswordAndHash(req.OldPassword, dbu.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !res {
		return c.JSON(http.StatusBadRequest, "Incorrect password")
	}

	// Hash the new password before updating it in the database
	hash, err := argon2id.CreateHash(req.NewPassword, argon2id.DefaultParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
	}

	// Update the password in the database
	var updateReq db.ChangePasswordParams
	updateReq.Username = c.Param("username")
	updateReq.PasswordHash = hash
	_, err = h.Queries.ChangePassword(c.Request().Context(), updateReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	_, err := h.Queries.DeleteUser(c.Request().Context(), c.Param("username"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) BecomeAdmin(c echo.Context) error {
	var req db.SetAdminStatusParams
	req.IsAdmin = true
	req.Username = c.Param("username")

	ctx := context.Background()
	_, err := h.Queries.SetAdminStatus(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to set admin status")
	}

	return c.NoContent(http.StatusOK)
}
