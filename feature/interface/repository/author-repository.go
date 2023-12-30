package repository

import "github.com/jepbura/go-server/feature/domain"

type AuthorRepo struct {
	handler DBHandler
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{handler}
}

func (repo AuthorRepo) SaveAuthor(author domain.Author) error {
	err := repo.handler.SaveAuthor(author)
	if err != nil {
		return err
	}
	return nil
}
