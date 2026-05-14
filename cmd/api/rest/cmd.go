package rest

import (
	"context"
	"log"

	"go-simple-template/internal/infrastructure"
	"go-simple-template/internal/interfaces/http/rest/router"
	"go-simple-template/internal/interfaces/http/rest/server"
)

func Start(ctx context.Context) {
	// init factory infrastructure
	factory := infrastructure.NewFactory().BuildRestFactory(ctx)

	// create a new router with the factory
	// this will allow the router to access the factory's resources
	// such as database connections, cache clients, etc.
	router := router.New(
		router.WithFactory(factory),
	)

	// initialize the router
	srv := server.New(ctx, router)

	// run the server
	err := srv.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run rest server: %v", err)
	}
}
