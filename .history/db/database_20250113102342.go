package main

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global variable that will hold the database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	// Adjust the connection string as per your database setup
	connStr := "user=yourusername dbname=yourdb password=yourpassword host=localhost sslmode=disable"
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
