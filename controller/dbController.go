package controller

import (
	"log"

	"github.com/jepbura/go-server/constant"
	"github.com/jepbura/go-server/database"
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DB_Controller interface {
	Save(book *model.Book)
	FindAll() []*model.Book
}

// DB Target is parameters to get all mux's dependencies
type DBTarget struct {
	fx.In
	MongoURL string `name:"mongo_url" optional:"true"`
	Lc       fx.Lifecycle
	Logger   *zap.Logger
}

func Save(document *model.Book) interface{} {
	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	// Free the resource when mainn dunction is  returned
	defer database.Close(client, ctx, cancel)
	cursor, err := database.SaveOne(client, ctx, string(constant.DB), string(constant.COL), document)
	// handle the errors.
	if err != nil {
		panic(err)
	}
	return cursor.InsertedID
}

func FindAll() []*model.Book {
	// Get Client, Context, CalcelFunc and err from connect method.

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	// Free the resource when mainn dunction is  returned
	defer database.Close(client, ctx, cancel)

	cursor, err := database.
		Query(client, ctx, string(constant.DB), string(constant.COL))
	// handle the errors.
	if err != nil {
		panic(err)
	}

	var results []*model.Book
	for cursor.Next(ctx) {
		var v *model.Book
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, v)
	}
	return results
}
