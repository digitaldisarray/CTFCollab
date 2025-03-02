package router

import (
	"os"

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
	e.Use(middleware.CORS()) // Dev env only?
	e.Use(middleware.Recover())
	// TODO: Rate limit?

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.CustomClaims)
		},
		SigningKey: []byte(handler.JWTSecret),
	}

	// CTF routes
	{

		ctfs := e.Group("/ctfs")
		ctfs.Use(echojwt.WithConfig(config))
		ctfs.GET("", handler.GetAllCTFs, auth.AdminOnly)
		ctfs.POST("", handler.CreateCTF)
		ctfs.GET("/joined", handler.GetJoinedCTFs)
		ctfs.GET("/search", handler.SearchCTFs, auth.AdminOnly) // can be changed I think
		ctfs.GET("/:phrase", handler.GetCTF, auth.MemberOnly(handler.Queries))
		ctfs.PUT("/:phrase", handler.UpdateCTF, auth.MemberOnly(handler.Queries))
		ctfs.DELETE("/:phrase", handler.DeleteCTF, auth.MemberOnly(handler.Queries))
		ctfs.POST("/:phrase/join", handler.JoinCTF)
		ctfs.GET("/:phrase/challenges", handler.GetChallenges, auth.MemberOnly(handler.Queries))
		ctfs.POST("/:phrase/challenges", handler.CreateChallenge, auth.MemberOnly(handler.Queries))
		// TODO: Route to get participants for a CTF, accessible to CTF members
	}

	// Challenge routes
	{
		//e.GET("/challenges/:id", ) // detailed information about a challenge, session has to be in the ctf it belongs to, or admin
		//e.DELETE("/challenges/:id", handler.DeleteChallenge) // session has to belong to ctf
		//e.PUT("/challenges/:id", )
	}

	// User routes
	e.POST("/users/guest", handler.CreateGuest) // Accessible without JWT
	e.POST("/users", handler.CreateUser)        // Accessible without JWT
	e.POST("/users/login", handler.LoginUser)   // Accessible without JWT
	{
		users := e.Group("/users")
		users.Use(echojwt.WithConfig(config))
		users.GET("/:username", handler.GetUser, auth.SelfOnly)
		users.DELETE("/:username", handler.DeleteUser, auth.SelfOnly)
		users.POST("/:username/password", handler.ChangePassword, auth.SelfOnly)
	}

	// Testing only
	// TODO: Rewrite tests to use built in admin instead of this endpoint
	if os.Getenv("TEST_MODE") == "True" {
		e.GET("/users2/:username/become_admin", handler.BecomeAdmin)
	}

	return e
}
