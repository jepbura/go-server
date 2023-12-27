package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jepbura/go-server/constant"
	"github.com/jepbura/go-server/database"
	"github.com/jepbura/go-server/graph"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Target is parameters to get all mux's dependencies
type Target struct {
	fx.In
	Environment string `name:"env"`
	Port        string `name:"port"`
	Lc          fx.Lifecycle
	Logger      *zap.Logger
}

// New is constructor to create Mux server on specific addr and port
func RunServer(target Target) {
	fmt.Println("GlobalResult:", target.Port)
	port := target.Port
	if port == "" {
		port = string(constant.Port)
	}

	database.GetConnection()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// target.Logger.Info("Stopping HTTPS server.")
}
