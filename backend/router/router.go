package router

import (
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

	// Routes
	e.GET("/hello", handler.HelloWorld)

	/* CTFS */
	e.GET("/ctfs", handler.GetCTFs)
	e.GET("/ctfs/:id", handler.GetCTF)
	e.DELETE("/ctfs/:id", handler.DeleteCTF)
	e.PATCH("/ctfs/:id", handler.UpdateCTF)
	e.POST("/ctfs", handler.CreateCTF)

	/* USERS */
	e.POST("/user", handler.CreateUser)
	e.POST("/user/login", handler.LoginUser)
	e.PUT("/users/password", handler.ChangePassword) // admin, or account owner
	e.DELETE("/users/:id", handler.DeleteUser)       // admin, or account owner

	/*
		GET user/:id
		DELETE user/:id
		UPDATE user/:id

		GET sessions/:id
		PATCH sessions
		DELETE sessions

		---- maybe?----
		GET document
		POST document
		UPDATE document
		DELETE document
		-----------------

		GET challenges
		POST challenges


		GET rooms
		GET room/:id
		POST room
		UPDATE room/:id
		DELETE room/:id

	*/

	return e
}
