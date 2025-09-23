package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// APIResponse represents a standard API response
// @Description Standard API response format
type APIResponse struct {
	Success   bool        `json:"success" example:"true"`
	Message   string      `json:"message" example:"Operation completed successfully"`
	Data      interface{} `json:"data,omitempty"`
	Errors    []string    `json:"errors,omitempty"`
	Timestamp time.Time   `json:"timestamp" example:"2025-09-24T10:00:00Z"`
	Path      string      `json:"path,omitempty" example:"/api/users"`
}

// SuccessResponse creates a successful API response
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	response := APIResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		Path:      c.Path(),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreatedResponse creates a successful creation response
func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	response := APIResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
		Path:      c.Path(),
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// ErrorResponse creates an error API response
func ErrorResponse(c *fiber.Ctx, status int, message string, errors []string) error {
	response := APIResponse{
		Success:   false,
		Message:   message,
		Errors:    errors,
		Timestamp: time.Now(),
		Path:      c.Path(),
	}
	return c.Status(status).JSON(response)
}

// ValidationErrorResponse creates a validation error response
func ValidationErrorResponse(c *fiber.Ctx, errors []string) error {
	return ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errors)
}

// InternalErrorResponse creates an internal server error response
func InternalErrorResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusInternalServerError, message, nil)
}

// NotFoundResponse creates a not found error response
func NotFoundResponse(c *fiber.Ctx, message string) error {
	return ErrorResponse(c, fiber.StatusNotFound, message, nil)
}
