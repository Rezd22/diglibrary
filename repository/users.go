package repository

import "paa/model"

func (b *booksRepo) CreateUser(user *model.User) error {
	err := b.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *booksRepo) GetUserByUsername(username string) (model.User, error) {
	var user model.User

	query := `select * from users where username = ?`
	err := b.db.Raw(query, username).Scan(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
