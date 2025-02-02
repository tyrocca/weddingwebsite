package main

import (
	"log"
	"os"
	"weddingwebsite/api"
	"weddingwebsite/container"
	"weddingwebsite/openapi/genopenapi"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(echomiddleware.Logger())

	///////////////////////////////////
	// SETUP ALL THE REPOSITORIES
	///////////////////////////////////

	// Select repository based on environment variable
	var svcContainer *container.Container

	dbType := os.Getenv("DB_TYPE") // "postgres" or "memory"

	if dbType == "postgres" {
		// Where postgres would happen
		log.Println("Running with PostgreSQL database")
		svcContainer = nil
	} else {
		log.Println("Running with in-memory database")
		svcContainer = container.NewInMemoryContainer()
	}

	///////////////////////////////////
	// SETUP ALL THE API Handlers
	///////////////////////////////////

	weddingHandlerStrict := api.NewWeddingHandler(svcContainer)

	// this can take middleware?
	handler := genopenapi.NewStrictHandler(weddingHandlerStrict, nil)
	genopenapi.RegisterHandlers(e, handler)

	///////////////////////////////////
	// SETUP ALL THE ROUTES
	///////////////////////////////////

	e.Logger.Fatal(e.Start(":8080"))
}
