package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fraud-detector/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Failed to connect DB: ", err)
	}

	// Auto-migrate
	err = DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal(" Failed to migrate DB: ", err)
	}

	fmt.Println("Database connected & migrated")
}
