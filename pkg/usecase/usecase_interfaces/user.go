package usecase_interfaces

import (
	"context"

	"github.com/jepbura/go-server/pkg/domain"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
}
