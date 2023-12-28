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
