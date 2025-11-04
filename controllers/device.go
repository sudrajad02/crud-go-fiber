package controllers

import (
	"crud-go-fiber/database"
	"crud-go-fiber/helpers"
	"crud-go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetDevices(c *fiber.Ctx) error {
	var devices []models.Device
	database.DB.Find(&devices)

	if len(devices) == 0 {
		helpers.ApiResponse(c, 404, "Data device empty", nil, "Data empty")
		return nil
	}

	helpers.ApiResponse(c, 200, "Data device", devices, nil)
	return nil
}

func GetDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	var device models.Device

	if err := database.DB.First(&device, id).Error; err != nil {
		helpers.ApiResponse(c, 404, "Data not found", nil, "Data not found")
		return nil
	}

	helpers.ApiResponse(c, 200, "Data device", device, nil)
	return nil
}

func CreateDevice(c *fiber.Ctx) error {
	var device models.Device

	if err := c.BodyParser(&device); err != nil {
		helpers.ApiResponse(c, 500, "Failed add data", nil, "Invalid request body format")
		return nil
	}

	database.DB.Create(&device)
	helpers.ApiResponse(c, 200, "Data device", device, nil)
	return nil
}

func UpdateDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	var device models.Device

	if err := database.DB.First(&device, id).Error; err != nil {
		helpers.ApiResponse(c, 404, "Data not found", nil, "Data not found")
		return nil
	}

	if err := c.BodyParser(&device); err != nil {
		helpers.ApiResponse(c, 500, "Failed update data", nil, "Invalid request body format")
		return nil
	}

	database.DB.Save(&device)
	helpers.ApiResponse(c, 200, "Data device", device, nil)
	return nil
}

func DeleteDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Device{}, id)
	helpers.ApiResponse(c, 204, "Delete data device", nil, nil)
	return nil
}
