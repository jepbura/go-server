package mongo_controller

import (
	"context"
	"fmt"
	"time"

	"github.com/jepbura/go-server/constant"
	"github.com/jepbura/go-server/feature/infrastructure/database/mongo"
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// DB Target is parameters to get all mux's dependencies
type DBTarget struct {
	fx.In
	MongoURL string `name:"mongo_url" optional:"true"`
	Lc       fx.Lifecycle
	Logger   *zap.Logger
}

// func Save(document *model.Book) interface{} {
// 	// Get Client, Context, CalcelFunc and err from connect method.
// 	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Free the resource when mainn dunction is  returned
// 	defer database.Close(client, ctx, cancel)
// 	cursor, err := database.SaveOne(client, ctx, string(constant.DB), string(constant.COL), document)
// 	// handle the errors.
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cursor.InsertedID
// }

// func FindAll() []*model.Book {
// 	// Get Client, Context, CalcelFunc and err from connect method.

// 	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Free the resource when mainn dunction is  returned
// 	defer database.Close(client, ctx, cancel)

// 	cursor, err := database.
// 		Query(client, ctx, string(constant.DB), string(constant.COL))
// 	// handle the errors.
// 	if err != nil {
// 		panic(err)
// 	}

// 	var results []*model.Book
// 	for cursor.Next(ctx) {
// 		var v *model.Book
// 		err := cursor.Decode(&v)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		results = append(results, v)
// 	}
// 	return results
// }

type DBHandler struct {
	Connection *mongo.Connection
}

func NewDBHandler(conn *mongo.Connection) *DBHandler {
	return &DBHandler{
		Connection: conn,
	}
}

func (db *DBHandler) SaveUser_DBHandler(book *model.Book) interface{} {
	// Get Client from Connection
	client := db.Connection.Client()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Database(string(constant.DB)).Collection(string(constant.COL))
	cursor, err := collection.InsertOne(ctx, book)

	// handle the errors.
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}

func (db *DBHandler) FindAllUsers_DBHandler() []*model.Book {
	// Get Client from Connection
	fmt.Println("*********************************************")
	client := db.Connection.Client()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Database(string(constant.DB)).Collection(string(constant.COL))
	cursor, err := collection.Find(ctx, bson.M{})

	// handle the errors.
	if err != nil {
		panic(err)
	}

	var results []*model.Book
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}
	return results
}
