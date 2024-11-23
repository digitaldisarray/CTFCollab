package handler

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCTFs(c echo.Context) error {
	ctx := context.Background()
	ctfs, err := h.Queries.ListCTFs(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctfs)
}

func (h *Handler) CreateCTF(c echo.Context) error {
	// TODO: Authentication and authorization

	// Parse parameters from the request
	ctf := new(db.CreateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: Make sure author_id is set based off of their authentication

	// Create the CTF in the database
	ctx := context.Background()
	result, err := h.Queries.CreateCTF(ctx, *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure ctf was created
	ctf_id, err := VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctf_id)
}
