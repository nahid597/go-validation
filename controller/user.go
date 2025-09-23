package controller

import (
	"go-validator/database"
	"go-validator/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return models.ErrorResponse(c, fiber.StatusBadRequest, "Cannot parse JSON", []string{err.Error()})
	}

	// save to database
	if err := database.DB.Create(&user).Error; err != nil {
		return models.InternalErrorResponse(c, "Could not create user: "+err.Error())
	}

	// Create response without password
	userResponse := models.UserResponse{
		ID:    user.ID,
		Age:   user.Age,
		Name:  user.Name,
		Email: user.Email,
	}

	return models.CreatedResponse(c, "User created successfully", userResponse)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return models.InternalErrorResponse(c, "Could not retrieve users")
	}

	// Convert to response format (without passwords)
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:    user.ID,
			Age:   user.Age,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return models.SuccessResponse(c, "Users retrieved successfully", userResponses)
}
