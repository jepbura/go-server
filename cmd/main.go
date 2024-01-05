package main

import (
	"fmt"
	"log"

	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/di"
)

func main() {
	config, configErr := config.EnvInit()
	fmt.Println("config is: ", config.DBName)
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	fmt.Println("config is: ", config)

	server, diErr := di.InitializeAPP(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
