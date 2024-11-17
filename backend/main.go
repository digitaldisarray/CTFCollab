package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/digitaldisarray/ctfcollab/handler"
	"github.com/digitaldisarray/ctfcollab/router"
)

func main() {
	// Setup handlers
	h, err := handler.LoadHandler("root:FakePassword@/ctfcollab")
	if err != nil {
		log.Panic(err)
	}

	// Setup API routes
	e := router.SetupRouter(h)

	// Start the API with a specified port
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", "1337")))
}
