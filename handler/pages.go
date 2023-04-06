package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *Handler) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) ShowAddBookPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add_books.html", nil)
}

func (h *Handler) ShowEditBookPage(c *gin.Context) {
	id := c.Param("id")

	buku, _ := h.booksRepo.GetBookById(id)
	c.HTML(http.StatusOK, "edit_books.html", buku)
}

func (h *Handler) DeletePage(c *gin.Context) {
	h.DeleteBook(c)
}
