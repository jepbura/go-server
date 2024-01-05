package http

import (
	"fmt"

	"github.com/gin-gonic/gin"

	handler "github.com/jepbura/go-server/pkg/api/handler"
	middleware "github.com/jepbura/go-server/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	fmt.Print("*********************************************\n")
	fmt.Print("NewServerHTTP\n")
	fmt.Print("*********************************************\n")
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// // Swagger docs
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", middleware.LoginHandler)

	// Auth middleware
	api := engine.Group("/api", middleware.AuthorizationMiddleware)

	api.GET("users", userHandler.FindAll)
	api.GET("users/:id", userHandler.FindByID)
	api.POST("users", userHandler.Save)
	api.DELETE("users/:id", userHandler.Delete)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
