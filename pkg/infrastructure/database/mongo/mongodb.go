package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// DBHandler is connection provider to access to global mongodb client
type MongoDBHandler struct {
	Client *mongo.Client
}

type MongoDbProvider interface {
	// ProvideMongoClient() *mongo.Client
	Disconnect(ctx context.Context) error
	Connect() gin.HandlerFunc
	WithContext(ctx context.Context) context.Context
}

func NewMongoDatabase(cnf config.Env, Logger *zap.Logger) (*mongo.Client, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("NewMongoDatabase\n")
	fmt.Print("*********************************************\n")
	// var dbHandler MongoDBHandler
	// dbHandlerError := DBHandler{}
	if cnf.MongoURL == "" {
		return nil, nil
	}

	dbHost := cnf.DBHost
	dbPort := cnf.DBPort
	dbUser := cnf.DBUser
	dbPass := cnf.DBPass

	dbHost = config.DefaultIfEmpty(dbHost, string(constant.DBHost))
	dbPort = config.DefaultIfEmpty(dbPort, string(constant.DBPort))

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Error check
	if err != nil {
		log.Fatal(err)
	}

	// Connect check
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// dbHandler := &MongoDBHandler{
	// 	Client: client,
	// }

	// return dbHandler, err
	//
	// dbHandler.Client = client
	// return &dbHandler, nil
	// return dbHandler.Client, err
	return client, nil
}

// // Client is a getter for the client field.
// func (c *MongoDBHandler) ProvideMongoClient() *mongo.Client {
// 	return c.Client
// }

func ProvideMongoClient(cnf config.Env, Logger *zap.Logger) (*mongo.Client, error) {

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type key string

const (
	// mongoClient key for mongo session in each request
	mongoClient key = "mongo_client"
)

func (dbHandler *MongoDBHandler) Disconnect(ctx context.Context) error {
	fmt.Println("Disconnecting from MongoDB!")
	return dbHandler.Client.Disconnect(ctx)
}

// Connect is method return adpater for http request that
func (m *MongoDBHandler) Connect() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m != nil {
			// save it in the mux context
			ctx := context.WithValue(c.Request.Context(), mongoClient, m.Client)
			c.Request = c.Request.WithContext(ctx)
		} else {
			// TODO: Warn
			log.Println("Warning: DBHandler is nil")
		}
		// pass execution to the original handler
		c.Next()
	}
}

// WithContext is method apply mongoClient into context
func (m *MongoDBHandler) WithContext(ctx context.Context) context.Context {
	if m != nil {
		// save it in the mux context
		return context.WithValue(ctx, mongoClient, m.Client)
	} else {
		// TODO: Warn
		log.Println("Warning: DBHandler is nil")
	}
	return ctx
}

// ForContext is method to get mongodb client from context
func ForContext(ctx context.Context) *mongo.Client {
	client, ok := ctx.Value(mongoClient).(*mongo.Client)
	if !ok {
		panic("ctx passing is not contain mongodb client")
	}
	return client
}
