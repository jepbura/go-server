package di

import (
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/feature/infrastructure/database"
	"github.com/jepbura/go-server/feature/infrastructure/logging"
	"github.com/jepbura/go-server/server"
	"go.uber.org/fx"
)

var DependencyInjection = fx.Options(
	fx.Provide(
		config.EnvInit,
		logging.LoggerInit,
		server.RunServer,
		// Database
		// database.GetConnection,
		database.NewMongoDatabase,
		// mongodb.New,
		// Controller
		// controller.NewGraphQLController,
		// controller.NewAuth,
	),
	ServiceModule,
	RepositoriyModule,
)
