package domain

import (
	"github.com/jepbura/go-server/pkg/infrastructure/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User Model
type User = model.User

type UserWithId struct {
	User `bson:",inline"`
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
}

type NewUser = model.NewUser
