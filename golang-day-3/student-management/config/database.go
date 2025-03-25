package config

import (
	"fmt"
	"log"
	"student-management/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Use &DB to modify the global variable
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	DB = db // Assign db to the global variable
	fmt.Println("✅ Database connected!")

	// AutoMigrate ensures the schema is up-to-date
	if err := DB.AutoMigrate(&models.Student{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
}
