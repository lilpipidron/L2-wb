package postgresql

import (
	"dev11/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.Event{}); err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}
