package middleware

import (
	"bytes"
	"encoding/json"
	"go-validator/models"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type rule struct {
	pattern *regexp.Regexp
	message string
}

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

	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email is required",
		})
	}

	// Email must contain "@" symbol
	if len(user.Email) < 5 || !bytes.Contains([]byte(user.Email), []byte("@")) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Email must be a valid email address",
		})
	}

	if user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is required",
		})
	}

	// Password check
	if !doesPasswordMeetCriteria(user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password must be at least 8 characters long and include at least one uppercase letter, one lowercase letter, one digit, and one special character",
		})
	}

	c.Locals("user", user)

	return c.Next()
}

func doesPasswordMeetCriteria(password string) bool {
	if len(password) < 8 {
		return false
	}

	var rules = []rule{
		{regexp.MustCompile(`[A-Z]`), "at least one uppercase letter"},
		{regexp.MustCompile(`[a-z]`), "at least one lowercase letter"},
		{regexp.MustCompile(`[0-9]`), "at least one digit"},
		{regexp.MustCompile(`[!@#~$%^&*()+|_]`), "at least one special character"},
	}

	for _, r := range rules {
		if !r.pattern.MatchString(password) {
			return false
		}
	}

	return true
}
