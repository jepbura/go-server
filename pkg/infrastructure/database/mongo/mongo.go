package mongo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jepbura/go-server/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbInitProvider interface {
	Disconnect(ctx context.Context) error
	Connect() gin.HandlerFunc
	WithContext(ctx context.Context) context.Context
	Collection(ctx context.Context, col string) (*mongo.Collection, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Save(ctx context.Context, newUser domain.NewUser) (*domain.User, error)
	Delete(ctx context.Context, id string) (string, error)
}
