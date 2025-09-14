package controller

import (
	"go-validator/database"
	"go-validator/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	// save to database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	// Create response without password
	userResponse := models.UserResponse{
		ID:    user.ID,
		Age:   user.Age,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"user":    userResponse,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not retrieve users",
			"error":   err.Error(),
		})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"users":  userResponses,
	})
}
