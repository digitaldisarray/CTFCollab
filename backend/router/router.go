package router

import (
	"github.com/digitaldisarray/ctfcollab/auth"
	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.CustomClaims)
		},
		SigningKey: []byte("change_me"), // TODO: Get from .env
	}

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
		e.GET("/users/:username", handler.GetUser)                                    // admin, or account owner
		e.POST("/users/password", handler.ChangePassword, echojwt.WithConfig(config)) // admin, or account owner
		e.DELETE("/users/:username", handler.DeleteUser, echojwt.WithConfig(config))  // admin, or account owner
	}

	return e
}
