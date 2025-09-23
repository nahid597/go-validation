package routes

import (
	"go-validator/controller"
	"go-validator/middleware"
	"go-validator/models"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"version": "1.0.0",
		}
		return models.SuccessResponse(c, "Go Validator API is running!", data)
	})

	// API routes group
	api := app.Group("/api")

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Get("/", controller.GetUsers)
	userRoutes.Post("/", middleware.ValidateUser, controller.CreateUser)
}
