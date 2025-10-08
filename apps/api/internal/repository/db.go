package repository

import (
	"paddletraffic/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	return gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&model.Location{})
}
