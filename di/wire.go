//go:build wireinject
// +build wireinject

package di2

import (
	"github.com/google/wire"
	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/infrastructure/logging"
	"github.com/jepbura/go-server/pkg/infrastructure/server"
)

// var dbSet = wire.NewSet(
// 	mongo.NewMongoDatabase,
// 	wire.Bind(new(mongo.DBHandlerProvider), new(*mongo.MongoDBHandler)),
// )

// var repositorySet = wire.NewSet(
// 	repository.NewUserRepository,
// 	wire.Bind(new(repository.NewUserRepositoryProvider), new(*repository.RepositoryHandler)),
// )

// var usecaseSet = wire.NewSet(
// 	usecase.NewUserInteractor,
// 	wire.Bind(new(usecase.NewUserInteractorProvider), new(*usecase.UserInteractor)),
// )

// var controllersSet = wire.NewSet(
// 	controllers.NewUserController,
// 	wire.Bind(new(controllers.NewUserControllerProvider), new(*controllers.UserController)),
// )

// var NewUserRepository = wire.NewSet(
// 	wire.Struct(new(UserRepository), "*"),
// 	wire.Bind(new(IUserRepository), new(*UserRepository)))

func InitializeAPI(cfg config.Env) (*server.ServerHTTP, error) {
	wire.Build(
		logging.LoggerInit,
		// Database
		// mongo.NewMongoDatabase,
		// dbSet,
		// Repository
		// repository.NewUserRepository,
		// repositorySet,
		// Use case
		// usecase.NewUserInteractor,
		// usecaseSet,
		// Controller
		// controllers.NewUserController,
		// controllersSet,
		// Server
		server.RunServer,
	)

	// logger, err := logging.LoggerInit(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	// mongoDB, err := mongo.NewMongoDatabase(cfg, logger)
	// if err != nil {
	// 	return nil, err
	// }

	// userRepo := repository.NewUserRepository(*mongoDB)
	// userUsecase := usecase.NewUserInteractor(userRepo)
	// userController := controllers.NewUserController(userUsecase)

	// use `mongoDB`, `userRepo`, and `logger` in the rest of your function...

	return &server.ServerHTTP{}, nil
}
