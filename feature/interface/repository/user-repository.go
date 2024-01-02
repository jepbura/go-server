package repository

import (
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
)

type UserRepository struct {
	handler DBHandler
}

func NewUserRepository(handler DBHandler) UserRepository {
	return UserRepository{handler}
}

func (repo UserRepository) SaveUser_Repository(user model.NewUser) error {
	err := repo.handler.SaveUser_DBHandler(user)
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepository) FindAllUser_Repository() ([]*model.User, error) {
	results, err := repo.handler.FindAllUsers_DBHandler()
	if err != nil {
		return results, err
	}
	return results, nil
}
