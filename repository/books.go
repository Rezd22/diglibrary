package repository

import "paa/model"

func (b *booksRepo) CreateBook(book *model.Book) error {
	err := b.db.Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) GetAllBooks() ([]model.GetBooks, error) {
	var books []model.GetBooks
	query := `select id, title, author, release_year from books`
	err := b.db.Raw(query).Scan(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *booksRepo) GetBookById(id string) (model.Book, error) {
	var book model.Book
	query := `select * from books where id = ?`
	err := b.db.Raw(query, id).Scan(&book).Error
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (b *booksRepo) UpdateBook(id string, book model.Book) error {
	err := b.db.Model(&model.Book{}).Where("id=?", id).Updates(book).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) DeleteBook(id string) error {
	query := `delete from books where id = ?`
	err := b.db.Exec(query, id).Error
	if err != nil {
		return err
	}

	return nil
}
