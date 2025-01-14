package models

import "gorm.io/gorm"

// User struct represents a user model for the database
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
