//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/jepbura/go-server/pkg/infrastructure/server"

	"github.com/jepbura/go-server/pkg/config"
	mongodb "github.com/jepbura/go-server/pkg/infrastructure/database/mongo"
	"github.com/jepbura/go-server/pkg/infrastructure/graph"
	"github.com/jepbura/go-server/pkg/infrastructure/logging"
	repository "github.com/jepbura/go-server/pkg/repository/user_repository"
	"github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"
	userUsecase "github.com/jepbura/go-server/pkg/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Resolver *graph.Resolver
	Client   *mongo.Client
	// MongoDBHandler *mongodb.MongoDBHandler
	// Repo        *repository.UserDatabase
	// Usecase     *usecase.UserUseCase
	// UserHandler *handler.UserHandler
	// Usecase *usecase_interfaces.AllUseCaseInterface
	Http *http.ServerHTTP
}

var dbSet = wire.NewSet(
	mongodb.NewMongoDatabase,
	// mongodb.ProvideMongoClient,
	wire.Struct(new(mongodb.MongoDBHandler), "*"),
	wire.Bind(new(mongodb.MongoDbProvider), new(*mongodb.MongoDBHandler)),
)

var NewUserRepository = wire.NewSet(
	repository.NewUserRepository,
	// wire.Struct(new(repository.UserDatabase), "*"),
	// wire.Bind(new(repositoryProvider.UserRepository), new(*repository.UserDatabase)),
)

var usecaseSet = wire.NewSet(
	userUsecase.NewUserUseCase,

	wire.Struct(new(usecase_interfaces.UseCasesInterface), "*"),
)

// func InitializeAPP(cnf config.Env) (*http.ServerHTTP, error) {
// func InitializeAPP(cnf config.Env) (*mongodb.MongoDBHandler, error) {
func InitializeAPP(cnf config.Env) (*App, error) {
	panic(wire.Build(
		logging.LoggerInit,
		// Database
		// mongodb.NewMongoDatabase,
		dbSet,
		// Repository
		// repository.NewUserRepository,
		NewUserRepository,
		// Use case
		// usecase.NewUserUseCase,
		usecaseSet,
		// Server
		http.NewServerHTTP,
		wire.Struct(new(graph.Resolver), "*"),
		wire.Struct(new(App), "*"),
	))

	// logger, err := logging.LoggerInit(cnf)
	// if err != nil {
	// 	return nil, err
	// }

	// mongoDB, err := mongodb.NewMongoDatabase(cnf, logger)
	// if err != nil {
	// 	return nil, err
	// }

	// userRepo := repository.NewUserRepository(mongoDB)
	// userUsecase := usecase.NewUserUseCase(userRepo)
	// userController := handler.NewUserHandler(userUsecase)
	// myHttp := http.NewServerHTTP(userController)

	// return &http.ServerHTTP{}, nil
	// return &mongodb.MongoDBHandler{}, nil
	return &App{}, nil
}
