package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	_ "github.com/digitaldisarray/ctfcollab/docs"
	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/digitaldisarray/ctfcollab/router"
	"github.com/joho/godotenv"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func main() {
	// Get ready to read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load environment from .env, assuming they are already set in environment")
	}

	// Make sure JWT secret is found
	jwtSecret, found := os.LookupEnv("JWT_SECRET")
	if !found || jwtSecret == "" {
		log.Panic("JWT_SECRET environment variable not set")
	}

	// Make sure MySQL url is found
	dbUrl, found := os.LookupEnv("MYSQL_URL")
	if !found || dbUrl == "" {
		log.Panic("MYSQL_URL environment variable not set")
	}

	// Setup handlers
	h, err := handler.LoadHandler(dbUrl, jwtSecret)
	if err != nil {
		log.Panic(err)
	}

	// Setup API routes
	e := router.SetupRouter(h)

	//http://localhost:1337/swagger/index.html#/  (to look at docs)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start the API with a specified port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "1337")))
}
