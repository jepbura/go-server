package main

import (
	"log"

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
		diApp.Http.Start()
	}
}
