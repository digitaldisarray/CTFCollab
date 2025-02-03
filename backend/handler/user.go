package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
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
	var req db.ChangePasswordParams
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	log.Printf("ChangePassword request: %+v", req)

	// Hash the new password before updating it in the database
	hash, err := argon2id.CreateHash(req.PasswordHash, argon2id.DefaultParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
	}
	req.PasswordHash = hash

	// Update the password in the database
	_, err = h.Queries.ChangePassword(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	// TODO: Add authentication and authorization checks here

	_, err = h.Queries.DeleteUser(c.Request().Context(), int32(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusOK)
}
