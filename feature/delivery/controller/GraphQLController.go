package controller

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/feature/domain/graph"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// GraphQLController handle the graphql request, parse request to schema and return results
type GraphQLController struct {
	graphiQLEnable bool
	mongodb        *mongo.Connection

	logger *zap.Logger
}

// GraphQLControllerTarget is parameter object for geting all GraphQLController's dependency
type GraphQLControllerTarget struct {
	fx.In
	GraphiQLEnable bool `name:"graphiql_enable"`
	MongoDB        *mongo.Connection
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
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// GraphiQL is defining as the GraphiQL Page handler
func (m *GraphQLController) GraphiQL() gin.HandlerFunc {
	// h := playground.Handler("GraphQL", "/")
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Register is function to register all controller's endpoint handler
func (m *GraphQLController) Register(r *gin.Engine) {
	r.Use(m.mongodb.Connect()).
		Use(m.Middleware()).
		// Use(m.auth.Middleware()).
		POST("/query", m.GrqphQL())
	if !m.graphiQLEnable {
		r.GET("/", m.GraphiQL())
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
