package handler

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCTF(c echo.Context) error {
	idParam := c.Param("id")
	i, err := strconv.Atoi(idParam)
	id := int32(i)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := context.Background()
	ctf, err := h.Queries.GetCTF(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctf)
}

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

func (h *Handler) DeleteCTF(c echo.Context) error {
	// TODO need auth
	idParam := c.Param("id")
	i, err := strconv.ParseInt(idParam, 10, 32)
	id := int32(i)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	ctx := context.Background()
	err = h.Queries.DeleteCTF(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}

func (h *Handler) UpdateCTF(c echo.Context) error {
	// TODO need auth
	idParam := c.Param("id")
	i, err := strconv.Atoi(idParam)
	id := int32(i)
	if err != nil || idParam == "" {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctf := new(db.UpdateCTFParams)

	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctf.ID = id
	ctx := context.Background()
	_, err = h.Queries.UpdateCTF(ctx, *ctf)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Updated")
}
