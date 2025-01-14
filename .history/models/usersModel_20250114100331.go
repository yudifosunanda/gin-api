package userModels

// User struct represents a user model for the database
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
