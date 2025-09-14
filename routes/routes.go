package routes

import (
	"go-validator/controller"
	"go-validator/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Go Validator API is running!",
			"version": "1.0.0",
		})
	})

	// API routes group
	api := app.Group("/api")

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Get("/", controller.GetUsers)
	userRoutes.Post("/", middleware.ValidateUser, controller.CreateUser)
}
