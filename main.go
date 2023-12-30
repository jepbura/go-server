package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/di"
	"github.com/jepbura/go-server/feature/delivery/controller"
	"go.uber.org/fx"
)

// // Register function register all API controllers to Mux
//
//	func Register(target controller.Target) {
//		for _, controller := range target.Controllers {
//			controller.Register(target.Gin)
//		}
//	}
//
// Register function register all API controllers to Mux
func Register(lc fx.Lifecycle, gin *gin.Engine, ctrls []controller.Controller) {
	for _, ctrl := range ctrls {
		ctrl.Register(gin)
	}
}

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
