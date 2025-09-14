package models

// APIResponse represents a standard API response
// @Description Standard API response format
type APIResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error response
// @Description Error response format
type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"An error occurred"`
	Error   string `json:"error,omitempty" example:"Detailed error message"`
}

// UserCreateResponse represents the response when creating a user
// @Description User creation response
type UserCreateResponse struct {
	Status  string       `json:"status" example:"success"`
	Message string       `json:"message" example:"User created successfully"`
	User    UserResponse `json:"user"`
}

// UsersListResponse represents the response when getting all users
// @Description Users list response
type UsersListResponse struct {
	Status string         `json:"status" example:"success"`
	Users  []UserResponse `json:"users"`
}
