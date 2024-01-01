package di

import (
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/feature/delivery/controller"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"github.com/jepbura/go-server/feature/infrastructure/logging"
	"github.com/jepbura/go-server/feature/infrastructure/server"
	"go.uber.org/fx"
)

var DependencyInjection = fx.Options(
	fx.Provide(
		config.EnvInit,
		logging.LoggerInit,
		server.RunServer,
		// Database
		// database.GetConnection,
		mongo.NewMongoDatabase,
		// mongodb.New,
		// Controller
		controller.NewGraphQLController,
		// controller.NewAuth,
	),
	ServiceModule,
	RepositoriyModule,
)
