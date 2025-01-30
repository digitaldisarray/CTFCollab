package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteChallenge(c echo.Context) error {
	challengeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	ctx := context.Background()
	err = h.Queries.DeleteChallenge(ctx, int32(challengeID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}
