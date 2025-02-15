package handler

import (
	"context"
	"net/http"

	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/digitaldisarray/ctfcollab/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllCTFs(c echo.Context) error {
	ctx := context.Background()

	ctfs, err := h.Queries.ListAllCTFs(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctfs)
}

func (h *Handler) GetJoinedCTFs(c echo.Context) error {
	// Get JWT claims
	// If JWT claims exist, get all CTFs user is member of
	// If JWT claims don't exist (unauthenticated), return an error
	return c.String(http.StatusNotImplemented, "Not implemented")
}

func (h *Handler) GetCTF(c echo.Context) error {
	// Lookup CTF by phrase and return it
	ctx := context.Background()

	ctf, err := h.Queries.GetCTFByPhrase(ctx, c.Param("phrase"))
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

	// TODO: Make sure author_id is set based off of their JWT claims

	// Generate CTF phrase
	mnemonic, err := util.GenerateMnemonic(6)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	ctf.Phrase = mnemonic

	// Create the CTF in the database
	ctx := context.Background()
	result, err := h.Queries.CreateCTF(ctx, *ctf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure CTF was created
	_, err = VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"phrase": mnemonic})
}

func (h *Handler) DeleteCTF(c echo.Context) error {
	ctx := context.Background()
	err := h.Queries.DeleteCTFByPhrase(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}

func (h *Handler) UpdateCTF(c echo.Context) error {
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
	ctx := context.Background()
	challenges, err := h.Queries.GetCTFChallenges(ctx, c.Param("phrase"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, challenges)
}

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
