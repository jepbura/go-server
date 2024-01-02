package usecase

import (
	"log"

	"github.com/jepbura/go-server/feature/domain"
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

func NewUserInteractor(repository domain.UserRepository) UserInteractor {
	return UserInteractor{repository}
}

func (interactor *UserInteractor) SaveUser_Usecase(user model.User) error {
	err := interactor.UserRepository.SaveUser_Repo_Model(user)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *UserInteractor) FindAllUsers_Usecase() ([]*model.User, error) {
	results, err := interactor.UserRepository.FindAllUsers_Repo_Model()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
