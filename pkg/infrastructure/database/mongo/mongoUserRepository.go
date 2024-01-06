package mongo

import (
	"context"
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

func (m *MongoDBHandler) FindByID(ctx context.Context, id uint) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler FindByID\n")
	fmt.Print("*********************************************\n")
	var user domain.User
	// err := c.DB.First(&user, id).Error

	return user, nil
}

func (m *MongoDBHandler) Save(ctx context.Context, newUser domain.NewUser) (domain.User, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Save\n")
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

func (m *MongoDBHandler) Delete(ctx context.Context, user domain.User) error {
	fmt.Print("*********************************************\n")
	fmt.Print("MongoDBHandler Delete\n")
	fmt.Print("*********************************************\n")
	// err := c.DB.Delete(&user).Error

	return nil
}
