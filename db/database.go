package db

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // PostgreSQL driver
	"os"
  "github.com/joho/godotenv"
	
)

// DB is the global variable that will hold the database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	errEnv := godotenv.Load()

	if errEnv != nil{
		fmt.Println("no env found")
	}

  host := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")
  dbPassword := os.Getenv("DB_PASSWORD")
	
	// Adjust the connection string as per your database setup
	connStr := fmt.Sprintf("user=postgres dbname=%s password=%s host=%s sslmode=disable", dbName, dbPassword, host)
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Database connected successfully!")
}

// CloseDB closes the database connection when the application exits
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
