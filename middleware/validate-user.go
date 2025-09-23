package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-validator/models"
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type rule struct {
	pattern *regexp.Regexp
	message string
}

var validate *validator.Validate
var validatorInitOnce sync.Once

func InitValidator() {
	validate = validator.New()

	// Custom validation for password complexity
	validate.RegisterValidation("custom_password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		errs := DoesPasswordMeetCriteria(password)
		return len(errs) == 0
	})
}

func ValidateUser(c *fiber.Ctx) error {
	var user models.User

	decoder := json.NewDecoder(bytes.NewReader(c.Body()))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		return models.ErrorResponse(c, fiber.StatusBadRequest, "Cannot parse JSON", []string{err.Error()})
	}

	// Ensure validator is initialized (safe on first use)
	validatorInitOnce.Do(InitValidator)

	// run validator
	if err := validate.Struct(user); err != nil {
		errors := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Password" {
				errors = append(errors, DoesPasswordMeetCriteria(user.Password)...)
			} else {
				errors = append(errors, fmt.Sprintf("Field '%s': %s", err.Field(), err.Error()))
			}
		}

		return models.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed", errors)
	}

	c.Locals("user", user)

	return c.Next()
}

func DoesPasswordMeetCriteria(password string) []string {
	var errs []string
	if len(password) < 8 {
		errs = append(errs, "Password must be at least 8 characters long")
	}

	var rules = []rule{
		{regexp.MustCompile(`[A-Z]`), "at least one uppercase letter"},
		{regexp.MustCompile(`[a-z]`), "at least one lowercase letter"},
		{regexp.MustCompile(`[0-9]`), "at least one digit"},
		{regexp.MustCompile(`[!@#~$%^&*()+|_]`), "at least one special character"},
	}

	for _, r := range rules {
		if !r.pattern.MatchString(password) {
			errs = append(errs, "Password must contain "+r.message)
		}
	}

	return errs
}
