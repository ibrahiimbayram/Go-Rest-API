package handlers

import (
	. "../database"
	model "../models"
	"github.com/gofiber/fiber"
)

func GetProduct(c *fiber.Ctx) error {
	db := DB
	var product []model.Product

	db.Find(&product)

	if len(product) == 0 {

		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "succes", "data": product})
}

func CreateProduct(c *fiber.Ctx) error {
	db := DB
	var err error
	product := model.Product{}

	err = c.BodyParser(&product)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	err = db.Create(&product).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.JSON(fiber.Map{"status": "succes", "data": product})

}

func UpdateProduct(c *fiber.Ctx) error {
	db := DB
	var err error
	product := model.Product{}

	err = c.BodyParser(&product)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	err = db.Save(&product).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.JSON(fiber.Map{"status": "succes", "data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := DB
	var err error
	product := model.Product{}

	err = c.BodyParser(&product)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	err = db.Delete(&product).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": err})
	}

	return c.JSON(fiber.Map{"status": "succes", "message": "Deleted product"})
}
