package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/alexedwards/argon2id"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/labstack/echo/v4"
)

type Params struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32

	// The number of iterations over the memory.
	Iterations uint32

	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8

	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32

	// Length of the generated key. 16 bytes or more is recommended.
	KeyLength uint32
}

// CreateUser creates a new user with a hashed password
// @Summary Create a new user
// @Description Creates a new user in the database with a hashed password using the Argon2id algorithm
// @Tags users
// @Accept json
// @Produce json
// @Param user body db.CreateUserParams true "Create User"
// @Success 200 {object} map[string]interface{} "user_id: ID of the created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /user [post]
func (h *Handler) CreateUser(c echo.Context) error {
	// Parse request
	user := new(db.CreateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// TODO update for more passes, currently does only 1
	// automatically formats w/ random salt
	hash, err := argon2id.CreateHash(user.PasswordHash, argon2id.DefaultParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.PasswordHash = hash
	// Insert data into db
	ctx := context.Background()
	result, err := h.Queries.CreateUser(ctx, *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure user was created
	if _, err = VerifyParseResult(result, 1); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id, _ := result.LastInsertId()

	return c.JSON(http.StatusOK, map[string]interface{}{"user_id": id})
}

// LoginUser handles user login
// @Summary User login
// @Description Authenticates a user by comparing the submitted password to the hashed password stored in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body db.CreateUserParams true "Login Credentials"
// @Success 200 {object} db.CreateUserParams "Logged in user information"
// @Failure 400 {string} response "Invalid login credentials"
// @Failure 500 {string} response "Internal server error"
// @Router /user/login [post]
func (h *Handler) LoginUser(c echo.Context) error {
	user := new(db.CreateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	dbu, err := h.Queries.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := argon2id.ComparePasswordAndHash(user.PasswordHash, dbu.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !res {
		return c.JSON(http.StatusBadRequest, "Invalid login credentials.")
	}

	// probably attach auth here
	return c.JSON(http.StatusOK, user)
}

// ChangePassword updates the password for a user
// @Summary Update password
// @Description Hashes and updates the password for an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param req body db.ChangePasswordParams true "Password Update Request"
// @Success 200 "Password updated successfully"
// @Failure 400 {object} map[string]string "Invalid input or password criteria not met"
// @Failure 500 {object} map[string]string "Internal server error during password update"
// @Router /users/password [post]
func (h *Handler) ChangePassword(c echo.Context) error {
	var req db.ChangePasswordParams
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	log.Printf("ChangePassword request: %+v", req)

	//TODO: Implement authentication and authorization checks here

	// Hash the new password before updating it in the database
	hash, err := argon2id.CreateHash(req.PasswordHash, argon2id.DefaultParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
	}
	req.PasswordHash = hash

	// Update the password in the database
	_, err = h.Queries.ChangePassword(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUser removes a user from the database
// @Summary Delete user
// @Description Deletes a user by ID from the database, requires admin privileges
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 "User successfully deleted"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 500 {string} map[string]string "Internal server error"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	// TODO: Add authentication and authorization checks here

	_, err = h.Queries.DeleteUser(c.Request().Context(), int32(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusOK)
}
