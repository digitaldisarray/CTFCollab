package router

import (
	"github.com/digitaldisarray/ctfcollab/auth"
	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/digitaldisarray/ctfcollab/websocket"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:4173"}, // Your Svelte dev server
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Requested-With",
		},
		AllowCredentials: true,
	}))

	e.Use(middleware.Recover())
	// TODO: Rate limit?

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.CustomClaims)
		},
		SigningKey:  []byte(handler.JWTSecret),
		TokenLookup: "cookie:token",
	}

	// CTF routes
	{
		// No JWT needed

		e.POST("ctfs/:phrase/join-as-guest", handler.JoinCTFGuest)

		ctfs := e.Group("/ctfs")
		ctfs.Use(echojwt.WithConfig(config))

		ctfs.GET("", handler.GetAllCTFs, auth.AdminOnly)
		ctfs.POST("", handler.CreateCTF, auth.AdminOnly)
		ctfs.GET("/search", handler.SearchCTFs, auth.AdminOnly) // TODO: Search joined CTFs for regular users
		ctfs.DELETE("/:phrase", handler.DeleteCTF, auth.AdminOnly)
		ctfs.POST("/:phrase/join", handler.JoinCTF)
		ctfs.GET("/joined", handler.GetJoinedCTFs, auth.AdminOnly) // Changed to admin only cause getting errors loading ctfs in admin dashboard
		ctfs.GET("/:phrase", handler.GetCTF, auth.MemberOnly(handler.Queries))
		ctfs.PUT("/:phrase", handler.UpdateCTF, auth.MemberOnly(handler.Queries))
		ctfs.GET("/:phrase/participants", handler.GetParticipants, auth.MemberOnly(handler.Queries))
		ctfs.POST("/:phrase/add-participant", handler.AddParticipant, auth.MemberOnly(handler.Queries))
		ctfs.DELETE("/:phrase/remove-participant", handler.RemoveParticipant, auth.MemberOnly(handler.Queries))
		ctfs.POST("/:phrase/remove-participant", handler.RemoveParticipant, auth.MemberOnly(handler.Queries)) // for sendBeacon

		ctfs.GET("/:phrase/challenges", handler.GetChallenges, auth.MemberOnly(handler.Queries))
		ctfs.POST("/:phrase/challenges", handler.CreateChallenge, auth.MemberOnly(handler.Queries))
		ctfs.GET("/:phrase/challenge/:id", handler.GetChallenge, auth.MemberOnly(handler.Queries))
		ctfs.DELETE("/:phrase/challenges/:id", handler.DeleteChallenge, auth.MemberOnly(handler.Queries)) // session has to belong to ctf
		ctfs.PUT("/:phrase/challenges/:id/submit", handler.SubmitFlag, auth.MemberOnly(handler.Queries))
		// TODO: Route to get participants for a CTF, accessible to CTF members
		// TODO: Route to get participants for a challenge, accessible to CTF members

	}
	e.GET("/ctfs/:phrase/exists", handler.GetCTFExists)

	e.GET("/ws", func(c echo.Context) error {
		websocket.ServeWs(handler.WsHub, c.Response().Writer, c.Request())
		return nil
	})

	// Challenge routes
	{
		// chals := e.Group("/challenges")
		// chals.Use(echojwt.WithConfig(config))

		// TODO: Move the /phrase/challenge requests here & make authenticated & ctf membership checked
	}

	// User routes
	// e.POST("/users/guest", handler.CreateGuest) // Accessible without JWT // I think now we should have the only way of creating a guest be joining a ctf
	e.POST("/users", handler.CreateUser)      // Accessible without JWT
	e.POST("/users/login", handler.LoginUser) // Accessible without JWT
	{
		users := e.Group("/users")
		users.Use(echojwt.WithConfig(config))
		users.GET("/:username", handler.GetUser, auth.SelfOnly)
		users.DELETE("/:username", handler.DeleteUser, auth.SelfOnly)
		users.POST("/:username/password", handler.ChangePassword, auth.SelfOnly)
		users.POST("/logout", handler.LogoutUser, auth.SelfOnly)
	}

	return e
}
