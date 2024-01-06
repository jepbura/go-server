package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/domain"
)

type MongoDbProvider interface {
	Disconnect(ctx context.Context) error
	Connect() gin.HandlerFunc
	WithContext(ctx context.Context) context.Context
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	Save(ctx context.Context, user domain.NewUser) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
}
