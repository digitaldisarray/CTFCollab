package handler

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)
// GetJoinedCTFs returns all joined CTFs
// @Summary returns joined CTFs
// @Description Uses JWT to return all joined CTFs
// @Tags ctfs
// @Accept json
// @Produce json
// @Param user body db.GetJoinedCTFsParams true "Joined CTF"
// @Success 200 {object} map[string]interface{} "CTF: ctf name"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /ctfs [post]

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

// CreateCTF creates a new CTF 
// @Summary Create a CTF
// @Description Creates a CTF and verifies that the CTF has been created
// @Tags ctfs
// @Accept json
// @Produce json
// @Param user body db.CreateCTFParams true "Create CTF"
// @Success 200 {object} map[string]interface{} "ctf_id: ID of the created CTF"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {string} map[string]string "Internal server error"
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
func (h *Handler) GetChallenges(c echo.Context) error {
	// TODO: Make sure user is a member of the CTF

	ctx := context.Background()
	challenges, err := h.Queries.GetCTFChallenges(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, challenges)
}

// CreateCTF creates a new CTF challenge 
// @Summary Create a CTF challenge
// @Description Creates a CTF challenge 
// @Tags challenges
// @Accept json
// @Produce json
// @Param user body db.CreateChallengeParams true "Create challenge"
// @Success 200 {object} map[string]interface{} "challenge updated"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /ctfs [post]
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
