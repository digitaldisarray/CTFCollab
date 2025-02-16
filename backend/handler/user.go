package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/alexedwards/argon2id"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

type Params struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32

	// The number of iterations over the memory.
	Iterations uint32

	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8

	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32

	// Length of the generated key. 16 bytes or more is recommended.
	KeyLength uint32
}

func (h *Handler) CreateUser(c echo.Context) error {
	// Parse request
	user := new(db.CreateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// TODO update for more passes, currently does only 1
	// automatically formats w/ random salt
	hash, err := argon2id.CreateHash(user.PasswordHash, argon2id.DefaultParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.PasswordHash = hash
	// Insert data into db
	ctx := context.Background()
	result, err := h.Queries.CreateUser(ctx, *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure user was created
	if _, err = VerifyParseResult(result, 1); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id, _ := result.LastInsertId()

	return c.JSON(http.StatusOK, map[string]interface{}{"user_id": id})
}

func (h *Handler) LoginUser(c echo.Context) error {
	user := new(db.CreateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	dbu, err := h.Queries.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := argon2id.ComparePasswordAndHash(user.PasswordHash, dbu.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !res {
		return c.JSON(http.StatusBadRequest, "Invalid login credentials.")
	}

	// probably attach auth here
	return c.JSON(http.StatusOK, user)
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
