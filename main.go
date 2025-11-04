package main

import (
	"crud-go-fiber/database"
	"crud-go-fiber/models"
	"crud-go-fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Connect to DB
	database.Connect()
	database.DB.AutoMigrate(&models.Device{}, &models.Sensor{})

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("🚀 Server running on http://localhost:4321")
	app.Listen(":4321")
}
