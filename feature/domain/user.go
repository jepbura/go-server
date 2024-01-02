package domain

import "github.com/jepbura/go-server/feature/infrastructure/graph/model"

type UserRepository interface {
	SaveUser_Repo_Model(user model.User) error
	FindAllUsers_Repo_Model() ([]*model.User, error)
}
