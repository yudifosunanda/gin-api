package models

// User represents a user entity
type Users struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}
