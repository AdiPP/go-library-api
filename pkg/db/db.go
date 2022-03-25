package db

import (
	"log"

	"github.com/AdiPP/go/library-api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/library"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}