package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global variable that will hold the database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("no env found")
	}

	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Adjust the connection string as per your database setup
	connStr := fmt.Sprintf("user=postgres dbname=%s password=%s host=%s sslmode=disable", dbName, dbPassword, host)
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Database connected successfully!")
}
