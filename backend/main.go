package main

import (
	"fmt"
	"log"

	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/digitaldisarray/ctfcollab/router"
)

func main() {
	// Setup handlers (eventually specify db info too)
	h, err := handler.LoadHandler()
	if err != nil {
		log.Panic(err)
	}

	// Setup API routes
	e := router.SetupRouter(h)

	// Start the API with a specified port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "1337")))
}
