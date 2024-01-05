package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/infrastructure/database/mongo"
	"github.com/jepbura/go-server/pkg/infrastructure/graph"
	"golang.org/x/crypto/acme/autocert"

	"go.uber.org/zap"
)

type ServerHTTP struct {
	Environment    string `name:"env"`
	Port           string `name:"port"`
	GraphiQLEnable bool   `name:"graphiql_enable"`
	Logger         *zap.Logger
	Server         *http.Server
	Manager        *autocert.Manager
	engine         *gin.Engine
	MongoDB        *mongo.MongoDBHandler
}

func NewServerHTTP(cnf config.Env, Logger *zap.Logger) *ServerHTTP {
	fmt.Print("*********************************************\n")
	fmt.Print("RunServer\n")
	fmt.Print("*********************************************\n")
	var man *autocert.Manager
	engine := gin.New()

	// zap.Logger integration with gin
	engine.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(Logger, true))

	server := &ServerHTTP{
		Environment: cnf.Environment,
		Port:        cnf.Port,
		Logger:      Logger,
		engine:      engine,
	}

	if cnf.Environment != "local" {
		host := ""
		man = &autocert.Manager{
			Prompt: autocert.AcceptTOS,
			Cache:  autocert.DirCache("certs"),
		}

		server.Server = &http.Server{
			Addr:    host + ":443",
			Handler: engine,
			TLSConfig: &tls.Config{
				GetCertificate: man.GetCertificate,
			},
		}
		server.Manager = man
	} else {
		host := "localhost"
		server.Server = &http.Server{
			Addr:    host + ":" + cnf.Port,
			Handler: engine,
		}
	}

	return server
}

func (s *ServerHTTP) StartServer() {
	if s.Environment != "local" {
		s.Logger.Info("Starting HTTPS server at " + s.Server.Addr)
		go s.Server.ListenAndServeTLS("", "")
		go http.ListenAndServe(":80", s.Manager.HTTPHandler(nil))
	} else {
		s.Logger.Info("Starting HTTP server at " + s.Server.Addr)
		go s.Server.ListenAndServe()
	}
}

func (s *ServerHTTP) StopServer(ctx context.Context) error {
	s.Logger.Info("Stopping HTTPS server.")
	return s.Server.Shutdown(ctx)
}

func (s *ServerHTTP) StartGraphQLServer() {
	fmt.Print("*********************************************\n")
	fmt.Print("NewGraphQLController\n")
	fmt.Print("*********************************************\n")
	s.engine.Use(s.MongoDB.Connect()).
		Use(Middleware()).
		// Use(m.auth.Middleware()).
		POST("/query", GrqphQL())
	if !s.GraphiQLEnable {
		s.engine.GET("/", GraphiQL())
	}
}

// GrqphQL is defining as the GraphQL handler
func GrqphQL() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GraphiQL is defining as the GraphiQL Page handler
func GraphiQL() gin.HandlerFunc {
	// h := playground.Handler("GraphQL", "/")
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Middleware for GraphQL resolver to pass services into ctx
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// ctx = context.WithValue(ctx, admin.Key, m.admin)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
