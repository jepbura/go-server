package userRepository

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/config"
	"github.com/jepbura/go-server/pkg/constant"
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
	// return &UserDatabase{
	// 	DB:        DBHandler.Client,
	// 	DBHandler: DBHandler,
	// }
}

func (c *UserDatabase) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := c.DBHandler.FindAll(ctx)
	return users, err
}

func (c *UserDatabase) FindByID(ctx context.Context, id string) (domain.User, error) {
	user, err := c.DBHandler.FindByID(ctx, id)
	return user, err
}

func (c *UserDatabase) Save(ctx context.Context, newUser domain.NewUser) (domain.User, error) {
	dbHandler := c.DBHandler
	DBName := dbHandler.Env.DBName
	DBUserCOL := dbHandler.Env.DBUserCOL
	DBName = config.DefaultIfEmpty(DBName, string(constant.DBName))
	DBUserCOL = config.DefaultIfEmpty(DBUserCOL, string(constant.DB_USER_COL))
	fmt.Println("DBName is: ", DBName)
	fmt.Println("DBUserCOL is: ", DBUserCOL)
	fmt.Println("dbHandler.Client is: ", dbHandler.Client)
	db := mongodb.MongoDBInputsFunc{
		Client: dbHandler.Client,
		DBName: DBName,
		Col:    DBUserCOL,
	}
	user, err := dbHandler.Save(ctx, newUser, db)
	return user, err
}

func (c *UserDatabase) Delete(ctx context.Context, id string) (domain.User, error) {
	deleteUser, err := c.DBHandler.Delete(ctx, id)
	return deleteUser, err
}
