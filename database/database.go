package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// This method closes mongoDB connection and cancel context.
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// CancelFunc to cancel to context
	defer cancel()
	// client provides a method to close a mongoDB connection.
	defer func() {
		// client.Disconnect method also has deadline.
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// ctx will be used to set deadline for process, here deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	// Confirmation of connection
	return client, ctx, cancel, err
}

// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {
	// mongo.Client has Ping to ping mongoDB, deadline of the Ping method will be determined by cxt
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func GetConnection() {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	// Release resource when the main, function is returned.
	defer Close(client, ctx, cancel)
	// Ping mongoDB with Ping method
	ping(client, ctx)
}

func SaveOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database
	// method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)
	// InsertMany accept two argument of type Context and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string) (result *mongo.Cursor, err error) {
	// select database and collection.
	collection := client.Database(dataBase).Collection(col)
	// collection has an method Find, that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, bson.D{})
	return
}
