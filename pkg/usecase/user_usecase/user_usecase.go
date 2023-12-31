package userUsecase

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	"github.com/jepbura/go-server/pkg/repository/repository_interface"
	"github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"
)

type UserUseCase struct {
	userRepo repository_interface.RepositoryInterface
}

func NewUserUseCase(repo repository_interface.RepositoryInterface) usecase_interfaces.UserUsecaseInterface {
	fmt.Print("*********************************************\n")
	fmt.Print("NewUserUseCase\n")
	fmt.Print("*********************************************\n")
	return &UserUseCase{
		userRepo: repo,
	}
}

func (c *UserUseCase) FindAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := c.userRepo.FindAllUsers(ctx)
	return users, err
}

func (c *UserUseCase) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := c.userRepo.FindUserByID(ctx, id)
	return user, err
}

func (c *UserUseCase) SaveUser(ctx context.Context, newUser domain.NewUser) (*domain.User, error) {
	user, err := c.userRepo.SaveUser(ctx, newUser)

	return user, err
}

func (c *UserUseCase) DeleteUser(ctx context.Context, id string) (string, error) {
	deleteUserId, err := c.userRepo.DeleteUser(ctx, id)

	return deleteUserId, err
}
