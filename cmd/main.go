package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/di"
)

func main() {
	config, configErr := config.EnvInit()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	diApp, diErr := di.InitializeAPP(config)
	// fmt.Println("Server is: ", diApp)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		go diApp.Http.StartServer()
		go diApp.Http.StartGraphQLServer()

		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

		<-stopChan

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := diApp.Http.Server.Shutdown(ctx); err != nil {
			log.Fatal("could not gracefully shutdown the server: ", err)
		}
		log.Println("server gracefully stopped")
	}
}
