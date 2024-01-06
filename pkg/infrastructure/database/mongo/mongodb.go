package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/constant"
	"github.com/jepbura/go-server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// MongoDBHandler is connection provider to access to global mongodb client
type MongoDBHandler struct {
	Client *mongo.Client
}

type MongoDbProvider interface {
	Disconnect(ctx context.Context) error
	Connect() gin.HandlerFunc
	WithContext(ctx context.Context) context.Context
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
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

func (m *MongoDBHandler) FindAll(ctx context.Context) ([]*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindAll\n")
	fmt.Print("*********************************************\n")
	user := &domain.User{
		ID:          "1",
		Name:        "John",
		Surname:     "Doe",
		UserName:    "john_doe",
		Password:    "password123",
		NationalID:  "123456789",
		BirthYear:   "1990",
		PhoneNumber: "1234567890",
		FatherName:  "Doe Sr.",
		City:        "New York",
		Email:       "john.doe@example.com",
		Gender:      "Male",
		Role:        "User",
		PhotoURL:    "https://example.com/profile.jpg",
		Settings:    "default",
	}

	users := []*domain.User{user}

	return users, nil
}

func (m *MongoDBHandler) FindByID(ctx context.Context, id uint) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindByID\n")
	fmt.Print("*********************************************\n")
	var user domain.User
	// err := c.DB.First(&user, id).Error

	return user, nil
}

func (m *MongoDBHandler) Save(ctx context.Context, user domain.User) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
	fmt.Print("*********************************************\n")
	// err := c.DB.Save(&user).Error

	return user, nil
}

func (m *MongoDBHandler) Delete(ctx context.Context, user domain.User) error {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")
	// err := c.DB.Delete(&user).Error

	return nil
}
