package database

import (
	"log"

	"github.com/mohaali482/todo/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Todo{})

	return db
}