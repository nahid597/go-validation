package middleware

import (
	"bytes"
	"encoding/json"
	"go-validator/models"

	"github.com/gofiber/fiber/v2"
)

func ValidateUser(c *fiber.Ctx) error {
	var user models.User

	decoder := json.NewDecoder(bytes.NewReader(c.Body()))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	// validate name is not empty
	if len(user.Name) < 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Name must be at least 5 characters long",
		})
	}

	// Age must be between 18 and 100
	if user.Age < 18 || user.Age > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Age must be between 18 and 100",
		})
	}

	// Email must contain "@" symbol
	if len(user.Email) < 5 || !bytes.Contains([]byte(user.Email), []byte("@")) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email must be a valid email address",
		})
	}

	c.Locals("user", user)

	return c.Next()
}
