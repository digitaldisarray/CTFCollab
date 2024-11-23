package handler

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	// Parse request
	user := new(db.CreateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: hash password with argon2id (at least 2-3 passes) + salt

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

	return c.JSON(http.StatusOK, "User created")
}
