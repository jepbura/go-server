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

func (c *UserDatabase) FindAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := c.DBHandler.FindAllUsers(ctx)
	return users, err
}

func (c *UserDatabase) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := c.DBHandler.FindUserByID(ctx, id)
	return user, err
}

func (c *UserDatabase) SaveUser(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	user, err := c.DBHandler.SaveUser(ctx, newUser)
	return user, err
}

func (c *UserDatabase) DeleteUser(ctx context.Context, id string) (string, error) {
	deleteUserId, err := c.DBHandler.DeleteUser(ctx, id)
	return deleteUserId, err
}
