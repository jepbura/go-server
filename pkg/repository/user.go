package repository

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	repository_interface "github.com/jepbura/go-server/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	DB *mongo.Client
}

func NewUserRepository(DB *mongo.Client) repository_interface.UserRepository {
	return &UserDatabase{DB}
}

// func NewUserRepository(client mongodb.MongoDBHandler) interfaces.UserRepository {
// 	return &UserDatabase{client: client.Client}
// }

func (c *UserDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("FindAll Repository\n")
	fmt.Print("*********************************************\n")
	var users []domain.Users
	// err := c.DB.Find(&users).Error

	return users, nil
}

func (c *UserDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("FindByID Repository\n")
	fmt.Print("*********************************************\n")
	var user domain.Users
	// err := c.DB.First(&user, id).Error

	return user, nil
}

func (c *UserDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	fmt.Print("*********************************************\n")
	fmt.Print("Save Repository\n")
	fmt.Print("*********************************************\n")
	// err := c.DB.Save(&user).Error

	return user, nil
}

func (c *UserDatabase) Delete(ctx context.Context, user domain.Users) error {
	fmt.Print("*********************************************\n")
	fmt.Print("Delete Repository\n")
	fmt.Print("*********************************************\n")
	// err := c.DB.Delete(&user).Error

	return nil
}
