package main

import (
	"go-validator/database"
	"go-validator/models"
	"go-validator/routes"

	"github.com/gofiber/fiber/v2"
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

	app.Listen(":3000")
}
