package handler

import (
	"net/http"
	"paa/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBook(c *gin.Context) {
	var book model.Book
	err := c.ShouldBind(&book)

	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error binding book",
		}

		c.HTML(http.StatusBadRequest, "add_books.html", resp)
		return
	}

	err = h.booksRepo.CreateBook(&book)
	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error creating book",
		}

		c.HTML(http.StatusBadRequest, "add_books.html", resp)
		return
	}

	_ = model.Response{
		Data:    book,
		Message: "success creating book",
	}

	h.GetAllBooks(c)
}

func (h *Handler) GetAllBooks(c *gin.Context) {
	books, err := h.booksRepo.GetAllBooks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"message": "Error getting all books",
		})
		return
	}

	resp := model.Response{
		Data:    books,
		Message: "",
	}

	c.HTML(http.StatusOK, "books.html", resp)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	param := c.Param("id")
	var book model.Book
	err := c.ShouldBind(&book)

	paramint, _ := strconv.Atoi(param)

	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error binding book",
		}

		c.HTML(http.StatusBadRequest, "edit_books.html", resp)
		return
	}

	err = h.booksRepo.UpdateBook(param, book)
	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error updating book",
		}

		c.HTML(http.StatusBadRequest, "edit_books.html", resp)
		return
	}

	book.ID = paramint

	_ = model.Response{
		Data:    book,
		Message: "success updating book",
	}

	h.GetAllBooks(c)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	param := c.Param("id")
	// paramint, _ := strconv.Atoi(param)

	deletedBook, err := h.booksRepo.GetBookById(param)
	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error getting book",
		}

		c.HTML(http.StatusBadRequest, "books.html", resp)
		return
	}

	err = h.booksRepo.DeleteBook(param)
	if err != nil {
		resp := model.Response{
			Data:    nil,
			Message: "error deleting book",
		}

		c.HTML(http.StatusBadRequest, "books.html", resp)
		return
	}

	_ = model.Response{
		Data:    deletedBook,
		Message: "success deleting book",
	}

	h.GetAllBooks(c)
}
