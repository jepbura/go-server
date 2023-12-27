package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/constant"
	"github.com/jepbura/go-server/database"
	"github.com/jepbura/go-server/graph"
)

func main() {
	// Environment variables initialization
	GlobalResult, err := config.EnvInit()
	if err != nil {
		return
	}

	fmt.Println("GlobalResult:", GlobalResult.Port)
	fmt.Println("err:", err)

	port := GlobalResult.Port
	if port == "" {
		port = string(constant.Port)
	}

	database.GetConnection()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
