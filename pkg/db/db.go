package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := "host=127.0.0.1 user=testuser password=password dbname=testdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = DB
	return err
}
