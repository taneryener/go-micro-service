package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func New() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
