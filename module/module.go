package module

import (
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/database"
	"github.com/jepbura/go-server/logging"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.EnvInit,
		logging.LoggerInit,
		database.GetConnection,
		// server.RunServer,
		// server.New,
		// // Database
		// mongodb.New,
		// postgresql.New,
		// // Controller
		// controller.NewGraphQLController,
		// controller.NewAuth,
	),
	ServiceModule,
	RepositoriyModule,
)
