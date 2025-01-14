package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	dsn := "host=localhost user=postgres password= dbname=mydb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Println("Database connection successful")

	// Auto-migrate models
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
}
