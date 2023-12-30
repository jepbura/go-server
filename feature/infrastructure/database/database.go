package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jepbura/go-server/config"
	"github.com/jepbura/go-server/constant"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// GraphQLControllerTarget is parameter object for geting all GraphQLController's dependency
type DatabaseTarget struct {
	fx.In
	GraphiQLEnable bool   `name:"graphiql_enable"`
	MONGO_URL      string `name:"MONGO_URL"`
	DBHost         string `name:"DB_HOST"`
	DBPort         string `name:"DB_PORT"`
	DBUser         string `name:"DB_USER"`
	DBPass         string `name:"DB_PASS"`
	Lc             fx.Lifecycle
	Logger         *zap.Logger
}

func NewMongoDatabase(target DatabaseTarget) (mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := target.DBHost
	dbPort := target.DBPort
	dbUser := target.DBUser
	dbPass := target.DBPass

	dbHost = config.DefaultIfEmpty(dbHost, string(constant.DBHost))
	dbPort = config.DefaultIfEmpty(dbPort, string(constant.DBPort))

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbURI)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer CloseMongoDBConnection(client)

	return client, nil
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
