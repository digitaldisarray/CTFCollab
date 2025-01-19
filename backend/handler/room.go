package handler

import (
	"context"
	"net/http"
	"strconv"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

// GetRooms retrieves all rooms.
func (h *Handler) GetRooms(c echo.Context) error {
	rooms, err := h.Queries.GetAllRooms(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

// GetRoom retrieves a specific room by ID.
func (h *Handler) GetRoom(c echo.Context) error {
	roomID := c.Param("id")
	i, err := strconv.Atoi(roomID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id := int32(i)
	room, err := h.Queries.GetRoom(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}

// CreateRoom creates a new room.
func (h *Handler) CreateRoom(c echo.Context) error {
	room := new(db.CreateRoomParams)
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rooms := context.Background()
	result, err := h.Queries.CreateRoom(rooms, *room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure room was created
	roomID, err := VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roomID)
}

// DeleteRoom deletes a room.
func (h *Handler) DeleteRoom(c echo.Context) error {
	roomID := c.Param("id")
	i, err := strconv.Atoi(roomID)
	id := int32(i)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	rooms := context.Background()
	result, err := h.Queries.DeleteRoom(rooms, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check how many rows were affected to make sure of deletion
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "Room not found")
	}

	return c.JSON(http.StatusOK, "Deleted")
}
