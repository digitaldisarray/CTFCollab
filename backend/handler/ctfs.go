package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/digitaldisarray/ctfcollab/auth"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/digitaldisarray/ctfcollab/util"
	"github.com/digitaldisarray/ctfcollab/websocket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCTFs(c echo.Context) error {
	ctfs, err := h.Queries.ListAllCTFs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctfs)
}

func (h *Handler) GetJoinedCTFs(c echo.Context) error {
	// Get user ID from JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.CustomClaims)
	id := int32(claims.Id)

	if claims.LoggedIn {
		// Look up CTFs that logged in user belongs to
		ctfs, err := h.Queries.ListUsersJoinedCTFs(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, ctfs)
	} else {
		// Look up CTFs that guest belongs to
		ctfs, err := h.Queries.ListGuestsJoinedCTFs(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, ctfs)
	}
}

func (h *Handler) GetCTF(c echo.Context) error {
	ctf, err := h.Queries.GetCTFByPhrase(c.Request().Context(), c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctf)
}

func (h *Handler) CreateCTF(c echo.Context) error {
	// Parse parameters from the request
	ctf := new(db.CreateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure author_id is set based off of their JWT claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.CustomClaims)
	ctf.AuthorID = int32(claims.Id)

	// Generate CTF phrase
	mnemonic, err := util.GenerateMnemonic(6)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	ctf.Phrase = mnemonic

	// Create the CTF in the database
	result, err := h.Queries.CreateCTF(c.Request().Context(), *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure CTF was created
	_, err = VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Add the user to the CTF
	join := new(db.JoinCTFUserParams)
	ctf_id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	join.CtfID = int32(ctf_id)
	join.UserID = int32(claims.Id)
	h.Queries.JoinCTFUser(c.Request().Context(), *join)

	newChallenge := map[string]any{
		"id":      mnemonic,
		"name":    ctf.Name,
		"date":    ctf.StartDate,
		"status":  "pending", //TODO: Need to add as param for ctf maybe?
		"members": "0",       //TODO: Need to add as param for ctf maybe?
	}
	h.WsHub.Broadcast(websocket.Message{
		Type:    "ctf_added",
		Payload: newChallenge,
	})

	return c.JSON(http.StatusOK, echo.Map{"phrase": mnemonic})
}

func (h *Handler) DeleteCTF(c echo.Context) error {
	phrase := c.Param("phrase")
	err := h.Queries.DeleteCTFByPhrase(c.Request().Context(), phrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	h.WsHub.Broadcast(websocket.Message{
		Type:    "ctf_deleted",
		Payload: map[string]string{"id": phrase},
	})

	return c.NoContent(http.StatusOK)
}

func (h *Handler) UpdateCTF(c echo.Context) error {
	phrase := c.Param("phrase")
	ctf := new(db.UpdateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctf.Phrase = c.Param("phrase") // Set the target CTF to update

	// TOOD: Make sure we aren't modifying things we shouldn't here

	_, err := h.Queries.UpdateCTF(c.Request().Context(), *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	h.WsHub.Broadcast(websocket.Message{
		Type:    "ctf_updated",
		Payload: map[string]string{"id": phrase},
	})

	return c.NoContent(http.StatusOK)
}

func (h *Handler) JoinCTF(c echo.Context) error {
	// Get CTF ID from phrase
	phrase := c.Param("phrase")
	ctf, err := h.Queries.GetCTFByPhrase(c.Request().Context(), phrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.CustomClaims)

	if claims.LoggedIn {
		join := new(db.JoinCTFUserParams)
		join.CtfID = ctf.ID
		join.UserID = int32(claims.Id)
		_, err = h.Queries.JoinCTFUser(c.Request().Context(), *join)
	} else {
		join := new(db.JoinCTFGuestParams)
		join.CtfID = ctf.ID
		join.GuestID = int32(claims.Id)
		_, err = h.Queries.JoinCTFGuest(c.Request().Context(), *join)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// helper struct for SearchCTFs
type OuterSearchParams struct {
	Name        interface{} `json:"name"`
	Description interface{} `json:"description"`
	StartDate   string      `json:"start_date"`
	EndDate     string      `json:"end_date"`
}

/*
Searches CTFs with optional params:

	name: string
	description: string
	start_date: formatted datetime ex: (2025-04-20T09:00:00Z)
	end_date: formatted datetime ex: (2025-04-20T09:00:00Z)
*/
func (h *Handler) SearchCTFs(c echo.Context) error {
	search := new(OuterSearchParams)

	var startTime, endTime time.Time
	var err error

	if err := c.Bind(search); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// sqlcSearch is a sub struct, doesn't contain the dates
	sqlcSearch := new(db.SearchCTFsParams)
	sqlcSearch.Name = search.Name
	sqlcSearch.Description = search.Description
	ctfs, err := h.Queries.SearchCTFs(c.Request().Context(), *sqlcSearch)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var new_ctfs []db.Ctf

	// if a startDate is given, format it into startTime
	if search.StartDate != "" {

		startTime, err = time.Parse("2006-01-02T15:04:05Z", search.StartDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

	}

	// if an endDate is given, format it into endTime
	if search.EndDate != "" {

		endTime, err = time.Parse("2006-01-02T15:04:05Z", search.EndDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	// loop through each CTF and append the matches that adhere to the start and end dates, if applicable
	// if no dates are given, it'll append anyways, so new_ctfs is always correct
	for _, c := range ctfs {
		if (!startTime.IsZero() && c.StartDate.Before(startTime)) || (!endTime.IsZero() && c.EndDate.After(endTime)) {
			continue
		}
		new_ctfs = append(new_ctfs, c)
	}

	return c.JSON(http.StatusOK, new_ctfs)
}

// Used to check if a ctf exists without returning anymore data than necessary, meant to be used for those with and without roles
func (h *Handler) GetCTFExists(c echo.Context) error {
	phrase := c.Param("phrase")

	// Check if the record exists in DB (disregard the returned object)
	_, err := h.Queries.GetCTFByPhrase(c.Request().Context(), phrase)
	if err != nil {
		// If the error is 'no rows in result set', respond with a 404
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No CTF found with that phrase",
			})
		}
		// Otherwise, something else went wrong (DB issue etc.)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// If found, just return 200 and a minimal payload
	return c.JSON(http.StatusOK, map[string]bool{
		"exists": true,
	})
}
