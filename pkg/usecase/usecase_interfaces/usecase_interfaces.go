package usecase_interfaces

import (
	"context"

	"github.com/jepbura/go-server/pkg/domain"
)

type UseCasesInterface struct {
	UserUsecaseInterface UserUsecaseInterface
}

type UserUsecaseInterface interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Save(ctx context.Context, newUser domain.NewUser) (*domain.User, error)
	Delete(ctx context.Context, id string) (string, error)
}
