package usecase_interfaces

import (
	"context"

	"github.com/jepbura/go-server/pkg/domain"
)

type UseCasesInterface struct {
	UserUsecaseInterface UserUsecaseInterface
}

type UserUsecaseInterface interface {
	FindAllUsers(ctx context.Context) ([]*domain.User, error)
	FindUserByID(ctx context.Context, id string) (*domain.User, error)
	SaveUser(ctx context.Context, newUser domain.NewUser) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}
