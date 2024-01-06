package mongo

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jepbura/go-server/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDBHandler) FindAll(ctx context.Context) ([]*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindAll\n")
	fmt.Print("*********************************************\n")
	_user := &domain.User{
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

	_users := []*domain.User{_user}

	// Get the collection
	collection, err := m.Collection(ctx, m.Env.DBUserCOL)
	if err != nil {
		return nil, nil
	}

	result, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, nil
	}
	fmt.Println("result is: ", result)

	// users := []*domain.User{result}

	return _users, nil
}

func (m *MongoDBHandler) FindByID(ctx context.Context, id string) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindByID\n")
	fmt.Print("*********************************************\n")

	_user := &domain.User{
		ID:          fmt.Sprintf("T%d", rand.Int()),
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

	// Create a filter for the query
	filter := bson.M{"_id": id}

	// Insert an user
	result, err := collection.Find(ctx, filter)
	fmt.Println("result is: ", result)
	if err != nil {
		return nil, err
	}

	return _user, nil
}

func (m *MongoDBHandler) Save(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
	fmt.Print("*********************************************\n")

	_user := &domain.User{
		ID:          fmt.Sprintf("T%d", rand.Int()),
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

func (m *MongoDBHandler) Delete(ctx context.Context, id string) (*domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")
	user := &domain.User{
		ID:          fmt.Sprintf("T%d", rand.Int()),
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

	return user, nil
}
