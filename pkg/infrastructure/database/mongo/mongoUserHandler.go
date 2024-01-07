package mongo

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoDBHandler) FindAllUsers(ctx context.Context) ([]*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindAll\n")
	fmt.Print("*********************************************\n")

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return nil, err
	}

	// Find all users
	filter := bson.M{"_id": bson.M{"$exists": true, "$ne": nil}}
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	// Iterate through the results and decode each user
	var users []*domain.User
	for result.Next(ctx) {
		var user domain.UserWithId

		if err := result.Decode(&user); err != nil {
			return nil, err
		}

		// Create a new variable of type domain.User
		var userModel domain.User

		copier.Copy(&userModel, &user)

		// Convert the id from ObjectId to string
		userModel.ID = user.ID.Hex()

		users = append(users, &userModel)
	}

	return users, nil
}

func (m *MongoDBHandler) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindByID\n")
	fmt.Print("*********************************************\n")

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return nil, nil
	}

	var objID primitive.ObjectID
	var filter bson.M

	// Check if id is a valid ObjectId
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		objID = oid
		filter = bson.M{"_id": objID}
	} else {
		// If id is not a valid ObjectId, use it directly
		filter = bson.M{"_id": id}
	}

	// Find one user based on the filter
	var userWithId domain.UserWithId
	err = collection.FindOne(ctx, filter).Decode(&userWithId)
	if err != nil {
		return nil, err
	}
	var userModel domain.User

	copier.Copy(&userModel, &userWithId)

	// Convert the id from ObjectId to string
	userModel.ID = userWithId.ID.Hex()

	return &userModel, nil
}

func (m *MongoDBHandler) SaveUser(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
	fmt.Print("*********************************************\n")

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return nil, err
	}

	// Insert an user
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("result is: ", result.InsertedID)

	if result.InsertedID != nil {
		// Convert the InsertedID to primitive.ObjectID
		insertedID, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			return nil, fmt.Errorf("failed to convert InsertedID to primitive.ObjectID")
		}

		var userWithId domain.UserWithId

		// Use the InsertedID as _id in the filter for FindOne
		err = collection.FindOne(ctx, bson.M{"_id": insertedID}).Decode(&userWithId)
		if err != nil {
			return nil, err
		}

		var userModel domain.User

		copier.Copy(&userModel, &userWithId)

		// Convert the id from ObjectId to string
		userModel.ID = userWithId.ID.Hex()

		return &userModel, nil
	} else {
		return nil, fmt.Errorf("InsertOne did not return an InsertedID")
	}
}

func (m *MongoDBHandler) DeleteUser(ctx context.Context, id string) (string, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return "", nil
	}

	var objID primitive.ObjectID
	var filter bson.M

	// Check if id is a valid ObjectId
	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		objID = oid
		filter = bson.M{"_id": objID}
	} else {
		// If id is not a valid ObjectId, use it directly
		filter = bson.M{"_id": id}
	}

	// Insert an user
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}

	if result.DeletedCount == 1 {
		return id, nil
	} else {
		return "", nil

	}
}
