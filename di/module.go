package di

import (
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/database"
	"github.com/jepbura/go-server/feature/infrastructure/logging"
	"go.uber.org/fx"
)

var DependencyInjection = fx.Options(
	fx.Provide(
		config.EnvInit,
		logging.LoggerInit,
		// server.RunServer,
		// Database
		database.GetConnection,
		// mongodb.New,
		// Controller
		// controller.NewGraphQLController,
		// controller.NewAuth,
	),
	ServiceModule,
	RepositoriyModule,
)
