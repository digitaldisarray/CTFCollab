package handler

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetJoinedCTFs(c echo.Context) error {
	// Get JWT claims
	// If JWT claims exist, get all CTFs user is member of
	// If JWT claims don't exist (unauthenticated), return an error
	return c.String(http.StatusNotImplemented, "Not implemented")
}

// GetCTF returns a CTF by phrase
// @Summary Get CTF details
// @Description Returns CTF details by its mnemonic phrase
// @Tags ctfs
// @Produce json
// @Param phrase path string true "CTF phrase"
// @Success 200 {object} db.Ctf "CTF details"
// @Failure 500 {string} string "Internal server error"
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

// CreateCTF creates a new CTF
// @Summary Create a CTF
// @Description Creates a new CTF and returns its ID
// @Tags ctfs
// @Accept json
// @Produce json
// @Param ctf body db.CreateCTFParams true "CTF details"
// @Success 200 {object} map[string]interface{} "{\"ctf_id\": string}"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs [post]
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

// DeleteCTF deletes a CTF by phrase
// @Summary Delete a CTF
// @Description Deletes a CTF by its mnemonic phrase
// @Tags ctfs
// @Produce json
// @Param phrase path string true "CTF phrase"
// @Success 200 {string} string "Deleted"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase} [delete]
func (h *Handler) DeleteCTF(c echo.Context) error {
	// TODO Make sure user is CTF owner or an admin
	ctx := context.Background()
	err := h.Queries.DeleteCTFByPhrase(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}

// UpdateCTF updates a CTF by phrase
// @Summary Update CTF details
// @Description Updates CTF details by its mnemonic phrase
// @Tags ctfs
// @Accept json
// @Produce json
// @Param phrase path string true "CTF phrase"
// @Param ctf body db.UpdateCTFParams true "CTF update parameters"
// @Success 200 {string} string "Updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase} [put]
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

<<<<<<< HEAD

// GetChallenges returns all CTF challenges
// @Summary CTF challenges
// @Description returns CTF on success
// @Tags ctfs
// @Accept json
// @Produce json
// @Param user body db.GetCTFChallengesRow true "Get challenge"
// @Success 200 {object} map[string]interface{} "challenge returned"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /ctfs [post]
=======
// GetChallenges returns challenges for a CTF
// @Summary Get CTF challenges
// @Description Returns all challenges for a specific CTF
// @Tags challenges
// @Produce json
// @Param phrase path string true "CTF phrase"
// @Success 200 {string} string "List of challenges"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase}/challenges [get]
>>>>>>> 29259a58797c573c9dc87be9f1abd42ef1d4d80c
func (h *Handler) GetChallenges(c echo.Context) error {
	// TODO: Make sure user is a member of the CTF

	ctx := context.Background()
	challenges, err := h.Queries.GetCTFChallenges(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, challenges)
}

// CreateChallenge creates a new challenge
// @Summary Create a challenge
// @Description Creates a new challenge for a CTF
// @Tags challenges
// @Accept json
// @Produce json
// @Param phrase path string true "CTF phrase"
// @Param challenge body db.CreateChallengeParams true "Challenge details"
// @Success 200 {string} string "Challenge created"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /ctfs/{phrase}/challenges [post]
func (h *Handler) CreateChallenge(c echo.Context) error {

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
