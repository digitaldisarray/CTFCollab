package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeleteUser removes a challenge from the database
// @Summary Delete Challenge
// @Description Deletes a Challenge by ID from the database, requires admin privileges
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path int true "Challenge ID"
// @Success 200 "Challenge successfully deleted"
// @Failure 400 {object} map[string]string "Invalid Challenge ID"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /challenges/{id} [delete]
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
