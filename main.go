package main

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/di"
	"github.com/jepbura/go-server/feature/delivery/controller"
	"go.uber.org/fx"
)

// Register function register all API controllers to Mux

func Register(target controller.Target) {
	for _, controller := range target.Controllers {
		controller.Register(target.Gin)
	}
}

func main() {

	app := fx.New(
		di.DependencyInjection,
		fx.Invoke(Register),
		// fx.Invoke(server.RunServer),
		// fx.Invoke(database.GetConnection),
	)

	app.Run()

	if err := app.Start(context.Background()); err != nil {
		fmt.Println("Error starting the application:", err)
	}

	defer app.Stop(context.Background())
}
