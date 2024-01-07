package mongo

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoDBHandler) FindAll(ctx context.Context) ([]*domain.User, error) {
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

func (m *MongoDBHandler) FindByID(ctx context.Context, id string) (*domain.User, error) {
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
	var user domain.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *MongoDBHandler) Save(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
	fmt.Print("*********************************************\n")

	_user := &domain.User{
		// ID:          fmt.Sprintf("T%d", rand.Int()),
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

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return nil, nil
	}

	// Insert an user
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("result is: ", result)

	return _user, nil
}

func (m *MongoDBHandler) Delete(ctx context.Context, id string) (string, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")
	// deleteUserId := domain.User{
	// 	// ID:          fmt.Sprintf("T%d", rand.Int()),
	// 	Name:        "John",
	// 	Surname:     "Doe",
	// 	UserName:    "john_doe",
	// 	Password:    "password123",
	// 	NationalID:  "123456789",
	// 	BirthYear:   "1990",
	// 	PhoneNumber: "1234567890",
	// 	FatherName:  "Doe Sr.",
	// 	City:        "New York",
	// 	Email:       "john.doe@example.com",
	// 	Gender:      "Male",
	// 	Role:        "User",
	// 	PhotoURL:    "https://example.com/profile.jpg",
	// 	Settings:    "default",
	// }

	return "deleteUserId.ID", nil
}
