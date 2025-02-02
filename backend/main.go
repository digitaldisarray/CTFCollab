package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/digitaldisarray/ctfcollab/router"
	"github.com/joho/godotenv"
)

func main() {
	// Get ready to read environment variables
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load environment from .env, assuming they are already set in environment")
	}

	log.Println("MYSQL_URL:", os.Getenv("MYSQL_URL"))

	// Setup handlers
	dbUrl, found := os.LookupEnv("MYSQL_URL")
	if !found || dbUrl == "" {
		log.Panic("MYSQL_URL environment variable not set")
	}
	h, err := handler.LoadHandler(dbUrl)
	if err != nil {
		log.Panic(err)
	}

	// Setup API routes
	e := router.SetupRouter(h)

	// Start the API with a specified port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "1337")))
}
