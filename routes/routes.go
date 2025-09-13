package routes

import (
	"go-validator/controller"
	"go-validator/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Hello, World 3!",
		})
	})

	// API routes group
	api := app.Group("/api")

	// User routes
	userRoutes := api.Group("/users")
	userRoutes.Post("", middleware.ValidateUser, controller.CreateUser)
	userRoutes.Get("", controller.GetUsers)
}
