package model

type Book struct {
	ID          int    `gorm:"type:int;primaryKey"`
	Title       string `gorm:"type:varchar(255);not null" form:"title"`
	Author      string `gorm:"type:varchar(50);not null" form:"author"`
	ReleaseYear int    `gorm:"type:int;not null" json:"release_year" form:"release_year"`
	Description string `gorm:"type:varchar(255);not null" form:"description"`
}

type GetBooks struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear int
}
