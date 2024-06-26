package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gotodo/models"
)

var Gorm *gorm.DB

func Init() error {
	dsn := "host=localhost user=postgres password=password dbname=todos port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Todo{})

	Gorm = db

	return nil
}
