package main

import (
	"go-validator/database"
	"go-validator/models"
	"go-validator/routes"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "go-validator/docs" // Import generated docs
)

func main() {

	// connect to database
	database.ConnectDatabase()
	// run migrations
	database.DB.AutoMigrate(&models.User{})

	// setup fiber app
	app := fiber.New()

	// setup routes
	routes.SetupRoutes(app)

	// Swagger documentation route
	app.Get("/api/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":3000")
}
