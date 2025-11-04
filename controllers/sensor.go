package controllers

import (
	"crud-go-fiber/database"
	"crud-go-fiber/helpers"
	"crud-go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetSensors(c *fiber.Ctx) error {
	var sensors []models.Sensor
	database.DB.Find(&sensors)

	if len(sensors) == 0 {
		helpers.ApiResponse(c, 404, "Data sensor empty", nil, "Data empty")
		return nil
	}

	helpers.ApiResponse(c, 200, "Data sensor", sensors, nil)
	return nil
}

func GetSensor(c *fiber.Ctx) error {
	id := c.Params("id")
	var sensor models.Sensor

	if err := database.DB.First(&sensor, id).Error; err != nil {
		helpers.ApiResponse(c, 404, "Data not found", nil, "Data not found")
		return nil
	}

	helpers.ApiResponse(c, 200, "Data sensor", sensor, nil)
	return nil
}

func CreateSensor(c *fiber.Ctx) error {
	var sensor models.Sensor

	if err := c.BodyParser(&sensor); err != nil {
		helpers.ApiResponse(c, 500, "Failed add data", nil, "Invalid request body format")
		return nil
	}

	database.DB.Create(&sensor)
	helpers.ApiResponse(c, 200, "Data sensor", sensor, nil)
	return nil
}

func UpdateSensor(c *fiber.Ctx) error {
	id := c.Params("id")
	var sensor models.Sensor

	if err := database.DB.First(&sensor, id).Error; err != nil {
		helpers.ApiResponse(c, 404, "Data not found", nil, "Data not found")
		return nil
	}

	if err := c.BodyParser(&sensor); err != nil {
		helpers.ApiResponse(c, 500, "Failed update data", nil, "Invalid request body format")
		return nil
	}

	database.DB.Save(&sensor)
	helpers.ApiResponse(c, 200, "Data sensor", sensor, nil)
	return nil
}

func DeleteSensor(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Sensor{}, id)
	helpers.ApiResponse(c, 204, "Delete data sensor", nil, nil)
	return nil
}
