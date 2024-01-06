package repository

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	mongodb "github.com/jepbura/go-server/pkg/infrastructure/database/mongo"
	"github.com/jepbura/go-server/pkg/repository/repository_interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	DB        *mongo.Client
	DBHandler mongodb.MongoDBHandler
}

func NewUserRepository(DB *mongo.Client, DBHandler mongodb.MongoDBHandler) repository_interface.UserRepository {
	fmt.Print("*********************************************\n")
	fmt.Print("NewUserRepository\n")
	fmt.Print("*********************************************\n")
	return &UserDatabase{
		DB:        DB,
		DBHandler: DBHandler,
	}
}

func (c *UserDatabase) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := c.DBHandler.FindAll(ctx)
	return users, err
}

func (c *UserDatabase) FindByID(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.DBHandler.FindByID(ctx, id)
	return user, err
}

func (c *UserDatabase) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := c.DBHandler.Save(ctx, user)
	return user, err
}

func (c *UserDatabase) Delete(ctx context.Context, user domain.User) error {
	err := c.DBHandler.Delete(ctx, user)
	return err
}
