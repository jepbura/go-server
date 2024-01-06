package userRepository

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	mongodb "github.com/jepbura/go-server/pkg/infrastructure/database/mongo"
	"github.com/jepbura/go-server/pkg/repository/repository_interface"
)

type UserDatabase struct {
	// DB        *mongo.Client
	DBHandler *mongodb.MongoDBHandler
}

func NewUserRepository(DBHandler *mongodb.MongoDBHandler) repository_interface.RepositoryInterface {
	fmt.Print("*********************************************\n")
	fmt.Print("NewUserRepository\n")
	fmt.Print("*********************************************\n")
	return &UserDatabase{DBHandler}
}

func (c *UserDatabase) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := c.DBHandler.FindAll(ctx)
	return users, err
}

func (c *UserDatabase) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := c.DBHandler.FindByID(ctx, id)
	return user, err
}

func (c *UserDatabase) Save(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	user, err := c.DBHandler.Save(ctx, newUser)
	return user, err
}

func (c *UserDatabase) Delete(ctx context.Context, id string) (*domain.User, error) {
	deleteUser, err := c.DBHandler.Delete(ctx, id)
	return deleteUser, err
}
