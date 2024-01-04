package usecase

import (
	"context"

	"github.com/jepbura/go-server/pkg/domain"
	repository_interface "github.com/jepbura/go-server/pkg/repository/interface"
	usecase_interfaces "github.com/jepbura/go-server/pkg/usecase/interface"
)

type UserUseCase struct {
	userRepo repository_interface.UserRepository
}

func NewUserUseCase(repo repository_interface.UserRepository) usecase_interfaces.UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

func (c *UserUseCase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *UserUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *UserUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *UserUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
