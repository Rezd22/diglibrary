package handler

import (
	"paa/repository"
)

type Handler struct {
	booksRepo repository.BooksRepo
}

func NewBooksHandler(booksRepo repository.BooksRepo) *Handler {
	return &Handler{booksRepo}
}
