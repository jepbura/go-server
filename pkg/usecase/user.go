package usecase

import (
	"context"
	"fmt"

	"github.com/jepbura/go-server/pkg/domain"
	"github.com/jepbura/go-server/pkg/repository/repository_interface"
	"github.com/jepbura/go-server/pkg/usecase/usecase_interfaces"
)

type UserUseCase struct {
	userRepo repository_interface.UserRepository
}

func NewUserUseCase(repo repository_interface.UserRepository) usecase_interfaces.UserUseCase {
	fmt.Print("*********************************************\n")
	fmt.Print("NewUserUseCase\n")
	fmt.Print("*********************************************\n")
	return &UserUseCase{
		userRepo: repo,
	}
}

func (c *UserUseCase) FindAll(ctx context.Context) ([]*domain.User, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *UserUseCase) FindByID(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *UserUseCase) Save(ctx context.Context, newUser domain.NewUser) (domain.User, error) {
	user, err := c.userRepo.Save(ctx, newUser)

	return user, err
}

func (c *UserUseCase) Delete(ctx context.Context, user domain.User) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
