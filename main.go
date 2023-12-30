package main

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/di"
	"github.com/jepbura/go-server/server"
	"go.uber.org/fx"
)

func main() {

	// app := fx.New(
	// 	fx.Provide(
	// 		config.EnvInit, // Add this line
	// 		func() *zap.Logger {
	// 			// Initialize and return a *zap.Logger instance here
	// 			// Example:
	// 			logger, err := zap.NewProduction()
	// 			if err != nil {
	// 				fmt.Println("Error initializing logger:", err)
	// 			}
	// 			return logger
	// 		},
	// 	),
	// 	fx.Invoke(server.RunServer),
	// )

	app := fx.New(
		di.DependencyInjection,
		fx.Invoke(server.RunServer),
		// fx.Invoke(database.GetConnection),
	)

	app.Run()

	if err := app.Start(context.Background()); err != nil {
		fmt.Println("Error starting the application:", err)
	}

	defer app.Stop(context.Background())
}
