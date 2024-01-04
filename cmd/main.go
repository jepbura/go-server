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

	api, err := di.InitializeAPP(config)
	if err != nil {
		// handle error
	}

	fmt.Println("api is: ", api)
}
