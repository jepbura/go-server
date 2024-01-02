package repository

import (
	"github.com/jepbura/go-server/feature/domain"
	"github.com/jepbura/go-server/feature/infrastructure/graph/model"
)

type DBHandler interface {
	FindAllUsers_DBHandler() ([]*model.User, error)
	SaveUser_DBHandler(book model.User) error
	FindAllBooks() ([]*domain.Book, error)
	SaveBook(book domain.Book) error
	SaveAuthor(author domain.Author) error
}
