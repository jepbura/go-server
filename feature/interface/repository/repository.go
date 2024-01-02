package repository

import (
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
)

type DBHandler interface {
	SaveUser_DBHandler(book model.NewUser) error
	FindAllUsers_DBHandler() ([]*model.User, error)
}
