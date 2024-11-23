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
	e.GET("/ctfs", handler.GetCTFs)
	e.POST("/ctfs", handler.CreateCTF)
	e.POST("/user", handler.CreateUser)

	return e
}
