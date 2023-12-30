package controller

import (
	"fmt"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/feature/delivery/graph"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// GraphQLController handle the graphql request, parse request to schema and return results
type GraphQLController struct {
	graphiQLEnable bool
	mongodb        *mongo.Client

	logger *zap.Logger
}

// GraphQLControllerTarget is parameter object for geting all GraphQLController's dependency
type GraphQLControllerTarget struct {
	fx.In
	GraphiQLEnable bool `name:"graphiql_enable"`
	MongoDB        *mongo.Client
	Logger         *zap.Logger
}

// NewGraphQLController is a constructor for GraphQLController
func NewGraphQLController(target GraphQLControllerTarget) Result {
	return Result{
		Controller: &GraphQLController{
			graphiQLEnable: target.GraphiQLEnable,
			mongodb:        target.MongoDB,
			logger:         target.Logger,
		},
	}
}

// GrqphQL is defining as the GraphQL handler
func (m *GraphQLController) GrqphQL() gin.HandlerFunc {
	h := handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	fmt.Println("GraphQLController")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GraphiQL is defining as the GraphiQL Page handler
func (m *GraphQLController) GraphiQL() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Register is function to register all controller's endpoint handler
func (m *GraphQLController) Register(r *gin.Engine) {
	// r.Use(m.mongodb.Connect()).
	// 	Use(m.Middleware()).
	// 	// Use(m.auth.Middleware()).
	// 	POST("/v1/graphql", m.GrqphQL())
	if !m.graphiQLEnable {
		r.GET("/v1/graphiql", m.GraphiQL())
	}
}

// Middleware for GraphQL resolver to pass services into ctx
func (m *GraphQLController) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// ctx = context.WithValue(ctx, admin.Key, m.admin)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
