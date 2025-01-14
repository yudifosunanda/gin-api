package userModels
import "github.com/go-playground/validator/v10"

// User struct represents a user model for the database
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}
