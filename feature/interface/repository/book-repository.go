package repository

import "github.com/jepbura/go-server/feature/domain"

type BookRepo struct {
	handler DBHandler
}

func NewBookRepo(handler DBHandler) BookRepo {
	return BookRepo{handler}
}

func (repo BookRepo) SaveBook(book domain.Book) error {
	err := repo.handler.SaveBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (repo BookRepo) FindAll() ([]*domain.Book, error) {
	results, err := repo.handler.FindAllBooks()
	if err != nil {
		return results, err
	}
	return results, nil
}
