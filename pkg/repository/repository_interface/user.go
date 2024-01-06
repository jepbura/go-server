package repository_interface

import (
	"context"

	"github.com/jepbura/go-server/pkg/domain"
)

type RepositoryInterface interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	Save(ctx context.Context, newUser domain.NewUser) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
}
