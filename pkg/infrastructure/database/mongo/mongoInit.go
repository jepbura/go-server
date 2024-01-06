package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// MongoDBHandler is connection provider to access to global mongodb client
type MongoDBHandler struct {
	Env    config.Env
	Client *mongo.Client
}

type MongoDBInputsFunc struct {
	Client *mongo.Client
	DBName string
	Col    string
}

// func NewMongoDatabase(cnf config.Env, Logger *zap.Logger) (*mongo.Client, error) {
func NewMongoDatabase(cnf config.Env, Logger *zap.Logger) (*MongoDBHandler, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("NewMongoDatabase\n")
	fmt.Print("*********************************************\n")

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
	mongoDbHandler := &MongoDBHandler{
		Client: client,
		Env:    cnf,
	}
	return mongoDbHandler, nil
	// return client, nil
}

type key string

const (
	// mongoClient key for mongo session in each request
	mongoClient key = "mongo_client"
)

// Disconnect is method return adpater for http request that
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

func Collection(ctx context.Context, db MongoDBInputsFunc) (*mongo.Collection, error) {

	if db.Client == nil {
		fmt.Println("MongoDB client is nil")
		return nil, errors.New("MongoDB client is nil")
	}

	collection := db.Client.Database(db.DBName).Collection(db.Col)
	if collection == nil {
		fmt.Println("MongoDB collection is nil")
		return nil, errors.New("MongoDB collection is nil")
	}

	return collection, nil
}
