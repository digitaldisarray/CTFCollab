package handler

import (
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/digitaldisarray/ctfcollab/auth"
	db "github.com/digitaldisarray/ctfcollab/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserAuthParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordParams struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// CreateUser creates a new user with a hashed password
// @Summary Create a new user
// @Description Creates a new user in the database with a hashed password using the Argon2id algorithm
// @Tags users
// @Accept json
// @Produce json
// @Param user body db.CreateUserParams true "Create User"
// @Success 200 {object} map[string]interface{} "user_id: ID of the created user"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user [post]
func (h *Handler) CreateUser(c echo.Context) error {
	// Parse request
	user := new(UserAuthParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure username isn't taken
	_, err := h.Queries.GetUserByUsername(c.Request().Context(), user.Username)
	if err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	// TODO: update for more passes, currently does only 1
	// Hash password
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Copy over data into a new user struct
	newUser := new(db.CreateUserParams)
	newUser.PasswordHash = hash
	newUser.Username = user.Username

	// Insert data into db
	result, err := h.Queries.CreateUser(c.Request().Context(), *newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Make sure user was created and extract user ID
	userID, err := VerifyParseResult(result, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return user ID in response
	return c.JSON(http.StatusOK, map[string]interface{}{"user_id": userID})
}

func (h *Handler) GetUser(c echo.Context) error {
	username := c.Param("username")

	// Get user from database
	user, err := h.Queries.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Remove password hash before sending user data
	user.PasswordHash = ""
	return c.JSON(http.StatusOK, user)
}

// LoginUser handles user login
// @Summary User login
// @Description Authenticates a user by comparing the submitted password to the hashed password stored in the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body db.CreateUserParams true "Login Credentials"
// @Success 200 {object} db.CreateUserParams "Logged in user information"
// @Failure 400 {string} string "Invalid login credentials"
// @Failure 500 {string} string "Internal server error"
// @Router /user/login [post]
func (h *Handler) LoginUser(c echo.Context) error {
	user := new(UserAuthParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Lookup user in database
	dbu, err := h.Queries.GetUserByUsername(c.Request().Context(), user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Make sure passwords match
	res, err := argon2id.ComparePasswordAndHash(user.Password, dbu.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if !res {
		return c.JSON(http.StatusBadRequest, "Invalid login credentials.")
	}

	// Create JWT token
	claims := &auth.CustomClaims{
		Name:     user.Username,
		LoggedIn: true,
		Id:       int(dbu.ID),
		IsAdmin:  dbu.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encoded_token, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		return err
	}

	// auth should be included in subsequent requests if you use:
	//			credentials: "include"

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    encoded_token,
		HttpOnly: true,
		Secure:   false, // TODO: Change to true when not localhost
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	return c.JSON(http.StatusOK, echo.Map{"token": encoded_token})
}

// ChangePassword updates the password for a user
// @Summary Update password
// @Description Hashes and updates the password for an existing user, requires admin privileges or account owner
// @Tags users
// @Accept json
// @Produce json
// @Param req body db.ChangePasswordParams true "Password Update Request"
// @Success 200 "Password updated successfully"
// @Failure 400 {string} string "Invalid input or password criteria not met"
// @Failure 500 {string} string "Internal server error during password update"
// @Router /users/password [post]
func (h *Handler) ChangePassword(c echo.Context) error {
	var req ChangePasswordParams
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// Bypass old password check if JWT is admin
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*auth.CustomClaims)
	if !claims.IsAdmin {
		// Make sure old passowwrd is correct
		dbu, err := h.Queries.GetUserByUsername(c.Request().Context(), c.Param("username")) // Lookup user in database
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Make sure passwords match
		res, err := argon2id.ComparePasswordAndHash(req.OldPassword, dbu.PasswordHash)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if !res {
			return c.JSON(http.StatusBadRequest, "Incorrect password")
		}
	}

	// Hash the new password before updating it in the database
	hash, err := argon2id.CreateHash(req.NewPassword, argon2id.DefaultParams)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Password hashing failed")
	}

	// Update the password in the database
	var updateReq db.ChangePasswordParams
	updateReq.Username = c.Param("username")
	updateReq.PasswordHash = hash
	_, err = h.Queries.ChangePassword(c.Request().Context(), updateReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password")
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUser removes a user from the database
// @Summary Delete user
// @Description Deletes a user by ID from the database, requires admin privileges or account owner
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 "User successfully deleted"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Internal server error"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	_, err := h.Queries.DeleteUser(c.Request().Context(), c.Param("username"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) BecomeAdmin(c echo.Context) error {
	var req db.SetAdminStatusParams
	req.IsAdmin = true
	req.Username = c.Param("username")

	_, err := h.Queries.SetAdminStatus(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to set admin status")
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) CreateGuest(c echo.Context) error {
	// Parse request body json
	type GuestParams struct {
		Nickname string `json:"nickname"`
	}
	guest := new(GuestParams)
	if err := c.Bind(guest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Check if nickname exists in guest database
	_, err := h.Queries.GetGuestByName(c.Request().Context(), guest.Nickname)
	if err == nil {
		return c.JSON(http.StatusBadRequest, "Guest already exists")
	}

	// Create JWT token for the guest
	claims := &auth.CustomClaims{
		Name:     guest.Nickname,
		LoggedIn: false, // Guest so false!
		IsAdmin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encoded_token, err := token.SignedString([]byte(h.JWTSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": encoded_token})
}
