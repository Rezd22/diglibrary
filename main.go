package main

import (
	"log"
	"paa/database"
	"paa/handler"
	"paa/model"
	"paa/repository"

	"github.com/gin-gonic/gin"
)

var credDB = model.Cred{
	Host:     "localhost",
	User:     "postgres",
	Password: "1377",
	DBName:   "books",
	Port:     5432,
}

func main() {
	db, err := database.ConnectDB(credDB)
	if err != nil {
		log.Fatalf("error connecting database : %v", err)
	}

	repo := repository.NewBooksRepository(db)
	handler := handler.NewBooksHandler(repo)

	r := gin.Default()

	// load file html & css
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")

	// handler api auth
	r.POST("/register", handler.CreateUser)
	r.POST("/login", handler.LoginUser)

	//handler api book
	r.POST("/book", handler.IsLogin, handler.CreateBook)
	r.GET("/book", handler.IsLogin, handler.GetAllBooks)
	r.POST("/book/:id", handler.IsLogin, handler.UpdateBook)
	r.DELETE("/book/:id", handler.IsLogin, handler.DeleteBook)

	// handler page auth
	r.GET("/", handler.ShowLoginPage)
	r.GET("/register", handler.ShowRegisterPage)

	// handler page book
	r.POST("/add-book", handler.IsLogin, handler.ShowAddBookPage)
	r.POST("/edit-book/:id", handler.IsLogin, handler.ShowEditBookPage)
	r.POST("/delete-book/:id", handler.IsLogin, handler.DeletePage)

	r.Run(":8080")
}
