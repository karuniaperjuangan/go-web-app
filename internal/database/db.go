package database

import (
	"log"

	"go-web-crud/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	log.Printf("Successfully connected")

	DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migrated")
}
