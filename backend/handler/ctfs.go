package handler

import (
	"context"
	"net/http"
	"time"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetJoinedCTFs(c echo.Context) error {
	// Get JWT claims
	// If JWT claims exist, get all CTFs user is member of
	// If JWT claims don't exist (unauthenticated), return an error
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetCTF(c echo.Context) error {
	// TODO: Make sure user is member of CTF through JWT, or is admin

	// Lookup CTF by phrase and return it
	ctx := context.Background()

	ctf, err := h.Queries.GetCTFByPhrase(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctf)
}

func (h *Handler) CreateCTF(c echo.Context) error {
	// TODO: Make sure user is signed in

	// Parse parameters from the request
	ctf := new(db.CreateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: Make sure author_id is set based off of their authentication
	// TODO: Generate mnemonic phrase here

	// Create the CTF in the database
	ctx := context.Background()
	result, err := h.Queries.CreateCTF(ctx, *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure CTF was created
	ctf_id, err := VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"ctf_id": ctf_id})
}

func (h *Handler) DeleteCTF(c echo.Context) error {
	// TODO Make sure user is CTF owner or an admin
	ctx := context.Background()
	err := h.Queries.DeleteCTFByPhrase(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}

func (h *Handler) UpdateCTF(c echo.Context) error {
	// TODO Make sure user is CTF owner or an admin

	ctf := new(db.UpdateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctf.Phrase = c.Param("phrase") // Set the target CTF to update

	ctx := context.Background()
	_, err := h.Queries.UpdateCTF(ctx, *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Updated")
}

func (h *Handler) JoinCTF(c echo.Context) error {
	// Get JWT claims
	// If no claims exist, generate them (anonymous account)
	// Add user to CTF
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetChallenges(c echo.Context) error {
	// TODO: Make sure user is a member of the CTF

	ctx := context.Background()
	challenges, err := h.Queries.GetCTFChallenges(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, challenges)
}

func (h *Handler) CreateChallenge(c echo.Context) error {
	// TODO: Make sure user is a member of the CTF

	challenge := new(db.CreateChallengeParams)
	if err := c.Bind(challenge); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	challenge.Phrase = c.Param("phrase") // Set the target CTF to add challenge under

	ctx := context.Background()
	_, err := h.Queries.CreateChallenge(ctx, *challenge)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Updated")
}

func (h *Handler) GetAllCTFs(c echo.Context) error {

	ctx := context.Background()
	ctfs, err := h.Queries.ListAllCTFs(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctfs)
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

	ctx := context.Background()
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
	ctfs, err := h.Queries.SearchCTFs(ctx, *sqlcSearch)

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
