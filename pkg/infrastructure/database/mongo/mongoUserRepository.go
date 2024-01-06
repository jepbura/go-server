package mongo

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jepbura/go-server/pkg/domain"
)

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

func (m *MongoDBHandler) FindByID(ctx context.Context, id string) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindByID\n")
	fmt.Print("*********************************************\n")
	user := domain.User{
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

func (m *MongoDBHandler) Save(ctx context.Context, newUser domain.NewUser, db MongoDBInputsFunc) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
	fmt.Print("*********************************************\n")

	_user := domain.User{
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
	// collection, err := Collection(ctx, db)
	if db.Client == nil {
		fmt.Println("MongoDB client is nil")
		return domain.User{}, errors.New("MongoDB client is nil")
	}

	collection := db.Client.Database(db.DBName).Collection(db.Col)

	if collection == nil {
		fmt.Println("collection is: ")
		fmt.Println("collection is: ", collection)
		return domain.User{}, errors.New("MongoDB collection is nil")
	}

	user, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return domain.User{}, err
	}
	fmt.Println("user is: ", user)
	return _user, nil

}

func (m *MongoDBHandler) Delete(ctx context.Context, id string) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")
	user := domain.User{
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
