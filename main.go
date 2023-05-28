package main

import (
	"fmt"
	"image"
	"net/http"
	"os"

	_ "image/jpeg" // Import untuk mendukung format JPEG
	_ "image/png"  // Import untuk mendukung format PNG
	"log"
	"paa/database"
	"paa/handler"
	"paa/model"
	"paa/repository"

	"github.com/gin-gonic/gin"
)

// Book struct
type Book struct {
	Title    string
	ImageURL string
}

var credDB = model.Cred{
	Host:     "localhost",
	User:     "postgres",
	Password: "wswas321",
	DBName:   "books",
	Port:     5432,
}

func main() {
	db, err := database.ConnectDB(credDB)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	file, err := os.Open("static/image/buku.jpg")
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membuka file:", err)
		return
	}
	defer file.Close()

	// Dekode gambar
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat mendekode gambar:", err)
		return
	}

	// Menampilkan informasi gambar
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	fmt.Println("Lebar gambar:", width)
	fmt.Println("Tinggi gambar:", height)

	repo := repository.NewBooksRepository(db)
	handler := handler.NewBooksHandler(repo)

	r := gin.Default()

	// load file html & css
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")

	// handler api auth
	r.POST("/register", handler.CreateUser)
	r.POST("/login", handler.LoginUser)

	// handler api book
	r.POST("/book", handler.IsLogin, handler.CreateBook)
	r.GET("/book", handler.IsLogin, handler.GetAllBooks)
	r.POST("/book/:id", handler.IsLogin, handler.UpdateBook)
	r.DELETE("/book/:id", handler.IsLogin, handler.DeleteBook)

	// handler page auth
	r.GET("/", HomeHandler)
	r.GET("/daftar",DaftarHandler)
	r.GET("/login", handler.ShowLoginPage)
	r.GET("/register", handler.ShowRegisterPage)

	// handler page book
	r.POST("/add-book", handler.IsLogin, handler.ShowAddBookPage)
	r.POST("/edit-book/:id", handler.IsLogin, handler.ShowEditBookPage)
	r.POST("/delete-book/:id", handler.IsLogin, handler.DeletePage)

	// Menampilkan pesan server berjalan
	fmt.Println("Server berjalan di http://localhost:8080")

	r.Run(":8080")
}

// HomeHandler adalah handler untuk halaman home ("/")
func HomeHandler(c *gin.Context) {
	books := []Book{
		{Title: "Harry Potter", ImageURL: "/static/images/harry_potter.jpg"},
		{Title: "The Great Gatsby", ImageURL: "/static/images/great_gatsby.jpg"},
		// Tambahkan data buku lainnya di sini
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
		"books": books,
	})
}
func DaftarHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "boks.html", gin.H{
		"title": "Daftar Buku",

	})
}
