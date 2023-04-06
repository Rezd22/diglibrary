package database

import (
	"fmt"
	"paa/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cred model.Cred) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", cred.Host, cred.User, cred.Password, cred.DBName, cred.Port)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Book{})

	return db, nil
}
