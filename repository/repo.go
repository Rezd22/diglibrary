package repository

import (
	"paa/model"

	"gorm.io/gorm"
)

type BooksRepo interface {
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (model.User, error)

	CreateBook(book *model.Book) error
	GetAllBooks() ([]model.GetBooks, error)
	GetBookById(id string) (model.Book, error)
	UpdateBook(id string, book model.Book) error
	DeleteBook(id string) error
}

type booksRepo struct {
	db *gorm.DB
}

func NewBooksRepository(db *gorm.DB) BooksRepo {
	return &booksRepo{db}
}
