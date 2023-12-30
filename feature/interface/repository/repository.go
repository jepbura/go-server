package repository

import "github.com/jepbura/go-server/feature/domain"

type DBHandler interface {
	FindAllBooks() ([]*domain.Book, error)
	SaveBook(book domain.Book) error
	SaveAuthor(author domain.Author) error
}
