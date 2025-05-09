package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/digitaldisarray/ctfcollab/websocket"
	"github.com/labstack/echo/v4"
)

// DeleteUser removes a challenge from the database
// @Summary Delete Challenge
// @Description Deletes a Challenge by ID from the database, requires admin privileges
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path int true "Challenge ID"
// @Success 200 {string} string "Challenge successfully deleted"
// @Failure 400 {string} string "Invalid Challenge ID"
// @Failure 500 {string} string "Internal server error"
// @Router /challenges/{id} [delete]
func (h *Handler) DeleteChallenge(c echo.Context) error {
	challengeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	ctx := context.Background()
	err = h.Queries.DeleteChallenge(ctx, int32(challengeID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	h.WsHub.Broadcast(websocket.Message{
		Type:    "chal_deleted",
		Payload: map[string]int{"id": challengeID},
	})

	return c.JSON(http.StatusOK, "Deleted")
}

func createHedgeDocNote() (string, error) {
	hedgeDocURL := os.Getenv("HEDGEDOC_URL")
	if hedgeDocURL == "" {
		hedgeDocURL = "http://hedgedoc:3000"
	}
	hedgedocAPI := hedgeDocURL + "/new"

	req, err := http.NewRequest("POST", hedgedocAPI, strings.NewReader("# Challenge Notes"))
	if err != nil {
		return "", fmt.Errorf("failed to create HedgeDoc request: %v", err)
	}

	// Build a custom client that doesn't follow redirects.
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to HedgeDoc: %v", err)
	}
	defer resp.Body.Close()

	// Should now have the response. Extract the 'Location' header:
	// location might look like: "http://localhost:3001/zVl9cPTpStGxb5R05x3M5w"
	location := resp.Header.Get("Location")
	if location == "" {
		return "", fmt.Errorf("HedgeDoc API did not return a valid Location header")
	}

	return location, nil
}

func (h *Handler) GetChallenges(c echo.Context) error {
	ctx := context.Background()
	challenges, err := h.Queries.GetCTFChallenges(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, challenges)
}

func (h *Handler) GetChallenge(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid challenge ID")
	}

	ctx := context.Background()

	challenge, err := h.Queries.GetChallenge(ctx, int32(id))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid challenge ID")
	}

	return c.JSON(http.StatusOK, challenge)
}

func (h *Handler) CreateChallenge(c echo.Context) error {
	challenge := new(db.CreateChallengeParams)
	if err := c.Bind(challenge); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	challenge.Phrase = c.Param("phrase") // Set the target CTF to add challenge under

	// Step 1: Create a new HedgeDoc note
	hurl, err := createHedgeDocNote()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	challenge.HedgedocUrl = hurl // Assign HedgeDoc note ID

	// Step 2: Insert the challenge into the database
	ctx := context.Background()

	res, err := h.Queries.CreateChallenge(ctx, *challenge)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get chal id
	insertedID, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get inserted challenge ID")
	}

	newChallenge := map[string]any{
		"id":           insertedID,
		"name":         challenge.Name,
		"description":  challenge.Description,
		"flag":         challenge.Flag,
		"hedgedoc_url": challenge.HedgedocUrl,
		"phrase":       challenge.Phrase,
		"status":       "pending",
	}

	h.WsHub.Broadcast(websocket.Message{
		Type:    "chal_added",
		Payload: newChallenge,
	})

	// Return the created challenge with the HedgeDoc URL
	return c.JSON(http.StatusOK, map[string]string{
		"challenge_name": challenge.Name,
		"hedgedoc_url":   hurl,
	})
}
