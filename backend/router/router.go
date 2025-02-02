package router

import (
	"net/http"

	"github.com/digitaldisarray/ctfcollab/auth"
	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(handler *handler.Handler) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	//e.Use(middleware.CORS()) // Dev env only?
	e.Use(middleware.Recover())

	// CTF routes
	{
		e.GET("/ctfs", handler.GetJoinedCTFs)
		e.GET("/ctfs/:phrase", handler.GetCTF)
		e.POST("/ctfs", handler.CreateCTF)
		e.DELETE("/ctfs/:phrase", handler.DeleteCTF)
		e.PUT("/ctfs/:phrase", handler.UpdateCTF)
		e.POST("/ctfs/:phrase/join", handler.JoinCTF)
		e.GET("/ctfs/:phrase/challenges", handler.GetChallenges)
		e.POST("/ctfs/:phrase/challenges", handler.CreateChallenge)
	}

	// Challenge routes
	{
		//e.GET("/challenges/:id", ) // detailed information about a challenge, session has to be in the ctf it belongs to, or admin
		e.DELETE("/challenges/:id", handler.DeleteChallenge) // session has to belong to ctf
		//e.PUT("/challenges/:id", )
	}

	// User routes
	{
		e.POST("/user", handler.CreateUser)
		e.POST("/user/login", handler.LoginUser)
		e.POST("/users/password", handler.ChangePassword) // admin, or account owner
		e.DELETE("/users/:id", handler.DeleteUser)        // admin, or account owner
	}

	return e
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract token from Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			// Handle anonymous user case or error if needed
			c.Set("user", nil)
			return next(c)
		}

		// Parse token
		tokenString := authHeader[len("Bearer "):]
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Set claims in context
		c.Set("user", claims)
		return next(c)
	}
}
