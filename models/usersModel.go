package models

// User struct represents a user model for the database
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID uint `json:"role_id"`

	Roles   Roles   `gorm:"foreignKey:RoleID;references:ID" json:"role"` // fix tag
	// Role   *Role `gorm:"foreignKey:RoleId" json:"role"`
}

