package di

import (
	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/feature/delivery/controller"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo/mongo_controller"
	"github.com/jepbura/go-server/feature/infrastructure/logging"
	"github.com/jepbura/go-server/feature/infrastructure/server"
	"github.com/jepbura/go-server/feature/interface/repository"
	"github.com/jepbura/go-server/feature/usecase"
	"go.uber.org/fx"
)

var DependencyInjection = fx.Options(
	fx.Provide(
		config.EnvInit,
		logging.LoggerInit,
		server.RunServer,
		// Controller
		controller.NewGraphQLController,
		// Database
		mongo.NewMongoDatabase,
		mongo_controller.NewDBHandler,
		// Repository
		repository.NewUserRepository,
		// Usecase
		usecase.NewUserInteractor,
	),
	ServiceModule,
	RepositoriyModule,
)
