package models

// User represents a user in the system
// @Description User account information
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
	Age      int    `gorm:"not null" json:"age" example:"25" validate:"required,gte=18,lte=100"`
	Name     string `gorm:"size:100;not null" json:"name" example:"John Doe" validate:"required,min=3,max=100"`
	Email    string `gorm:"size:100;unique;not null" json:"email" example:"john@example.com" validate:"required,email"`
	Password string `gorm:"size:100;not null" json:"password" example:"password123" validate:"custom_password"`
}

// UserRequest represents the request body for creating a user
// @Description User creation request
type UserRequest struct {
	Age      int    `json:"age" example:"25" validate:"required,min=1,max=120"`
	Name     string `json:"name" example:"John Doe" validate:"required,min=2,max=100"`
	Email    string `json:"email" example:"john@example.com" validate:"required,email"`
	Password string `json:"password" example:"password123" validate:"required,min=6"`
}

// UserResponse represents the response when creating/retrieving a user
// @Description User response (without password)
type UserResponse struct {
	ID    uint   `json:"id" example:"1"`
	Age   int    `json:"age" example:"25"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
}
