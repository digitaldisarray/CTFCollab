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

// GetAllCTFs gets all CTF's a user belongs to
// @Summary Get All CTFs
// @Description Gets every CTF a user belongs to
// @Tags ctfs
// @Produce json
// @Success 200 {araray} db.Ctf
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs [get]
func (h *Handler) GetAllCTFs(c echo.Context) error {
	ctfs, err := h.Queries.ListAllCTFs(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctfs)
}

// GetJoinedCTFs gets all ctf's a user is joined to
// @Summary Get Joined CTFs
// @Description Gets all CTFs a user is joined to based on their JWT token
// @Tags ctfs
// @Produce json
// @Success 200 {array} db.ListUsersJoinedCTFsRow
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/joined [get]
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
		// Look up CTF that guest belongs to
		ctfs, err := h.Queries.ListGuestsJoinedCTF(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, ctfs)
	}
}

// GetCTF gets a CTF based on its phrase
// @Summary Get CTF
// @Description Gets a CTF whose phrase matches the given phrase
// @Tags ctfs
// @Produce json
// @Success 200 {object} db.Ctf
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase} [get]
func (h *Handler) GetCTF(c echo.Context) error {
	ctf, err := h.Queries.GetCTFByPhrase(c.Request().Context(), c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctf)
}

// CreateCTF creates a new CTF with the requesting user as the author
// @Summary Create CTF
// @Description Creates a new CTF, creator is a participant of the new CTF
// @Tags ctfs
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs [post]
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

	newCTF := map[string]any{
		"id":      mnemonic,
		"name":    ctf.Name,
		"date":    ctf.StartDate,
		"status":  "pending", //TODO: Need to add as param for ctf maybe?
		"members": "0",       //TODO: Need to add as param for ctf maybe?
	}
	h.WsHub.Broadcast(websocket.Message{
		Type:    "ctf_added",
		Payload: newCTF,
	})

	return c.JSON(http.StatusOK, echo.Map{"phrase": mnemonic})
}

// DeleteCTF deletes a CTF by phrase
// @Summary Delete CTF
// @Description Deletes the CTF that matches the provided phrase
// @Tags ctfs
// @Produce json
// @Param phrase path string true "Phrase"
// @Success 200 {string} string "CTF deleted successfully"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase} [delete]
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

// UpdateCTF updates a CTF's information
// @Summary Update CTF
// @Description Updates the information of an existing CTF
// @Tags ctfs
// @Accept json
// @Produce json
// @Param phrase path string true "Phrase"
// @Success 200 {string} string "CTF updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase} [put]
func (h *Handler) UpdateCTF(c echo.Context) error {
	phrase := c.Param("phrase")
	ctf := new(db.UpdateCTFParams)
	if err := c.Bind(ctf); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctf.Phrase = c.Param("phrase") // Set the target CTF to update

	// TODO: Make sure we aren't modifying things we shouldn't here

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

// JoinCTF requesting user joins a CTF by phrase
// @Summary Join CTF (Logged-in user)
// @Description Joins a logged-in user to a CTF by phrase
// @Tags ctfs
// @Produce json
// @Param phrase path string true "Phrase"
// @Success 200 {string} string "Joined successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase}/join [post]
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
		return c.NoContent(http.StatusBadRequest) // Can't join CTF as guest w/ this endpoint
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// Parse nickname from request body
type GuestParams struct {
	Nickname string `json:"nickname"`
}

// JoinCTFGuest allows a guest user to join a CTF
// @Summary Join CTF (Guest)
// @Description Joins a guest user to a CTF based on nickname and returns a JWT token
// @Tags ctfs
// @Accept json
// @Produce json
// @Param phrase path string true "Phrase"
// @Param nickname body GuestParams false "Nickname"
// @Success 200 {string} string "Guest joined and token issued"
// @Failure 400 {object} map[string]string "Bad request or nickname taken"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /ctfs/{phrase}/join-as-guest [post]
func (h *Handler) JoinCTFGuest(c echo.Context) error {
	// Make sure CTF phrase is valid
	phrase := c.Param("phrase")
	ctf, err := h.Queries.GetCTFByPhrase(c.Request().Context(), phrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	guestParams := new(GuestParams)
	if err := c.Bind(guestParams); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Generate a nickname if not provided
	if guestParams.Nickname == "" {
		guestParams.Nickname = GenerateGuestNickname()
	}

	// Check if the nickname already exists for the specific CTF
	_, err = h.Queries.GetGuestByNameAndCTF(c.Request().Context(), db.GetGuestByNameAndCTFParams{
		Nickname: guestParams.Nickname,
		CtfID:    ctf.ID,
	})
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Nickname already taken for this CTF",
		})
	} else if err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create a new guest account
	newGuest := db.CreateGuestParams{
		Nickname: guestParams.Nickname,
		CtfID:    ctf.ID,
	}
	result, err := h.Queries.CreateGuest(c.Request().Context(), newGuest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure user was created
	_, err = VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Fetch user from DB to get info
	guest, err := h.Queries.GetGuestByNameAndCTF(c.Request().Context(), db.GetGuestByNameAndCTFParams{
		Nickname: guestParams.Nickname,
		CtfID:    ctf.ID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Create JWT token for the guest
	claims := &auth.CustomClaims{
		Name:     guestParams.Nickname,
		LoggedIn: false, // Guest accounts are not logged in
		Id:       int(guest.ID),
		IsAdmin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // Token valid for 72 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	// Return the token to the client
	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    encodedToken,
		HttpOnly: true,
		Secure:   false, // TODO: Change to true when not localhost
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	return c.NoContent(http.StatusOK) // Might need c.JSON(http.StatusOK, echo.Map{"token": encoded_token})
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
// SearchCTFs searches for CTFs by optional fields
// @Summary Search CTFs
// @Description Searches for CTFs by name, description, and optional start/end date ranges
// @Tags ctfs
// @Accept json
// @Produce json
// @Param search body OuterSearchParams false "Search filters"
// @Success 200 {array} db.Ctf
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/search [post]
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

/**Used to check if a ctf exists without returning anymore data than necessary, meant to be used for those with and without roles*/
// GetCTFExists checks if a CTF exists
// @Summary Check CTF Existence
// @Description Checks whether a CTF exists with the given phrase
// @Tags ctfs
// @Produce json
// @Param phrase path string true "CTF Phrase"
// @Success 200 {object} map[string]bool "Exists"
// @Failure 404 {object} map[string]string "CTF not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /ctfs/{phrase}/exists [get]
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
