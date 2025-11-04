package routes

import (
	"crud-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Device Routes
	deviceGroup := api.Group("/device")
	{
		deviceGroup.Get("/", controllers.GetDevices)
		deviceGroup.Get("/:id", controllers.GetDevice)
		deviceGroup.Post("/", controllers.CreateDevice)
		deviceGroup.Put("/:id", controllers.UpdateDevice)
		deviceGroup.Delete("/:id", controllers.DeleteDevice)
	}

	// Sensor Routes
	sensorGroup := api.Group("/sensor")
	{
		sensorGroup.Get("/", controllers.GetSensors)
		sensorGroup.Get("/:id", controllers.GetSensor)
		sensorGroup.Post("/", controllers.CreateSensor)
		sensorGroup.Put("/:id", controllers.UpdateSensor)
		sensorGroup.Delete("/:id", controllers.DeleteSensor)
	}
}
