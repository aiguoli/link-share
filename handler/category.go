package handler

import (
	"github.com/gofiber/fiber/v2"
	"link-share/database"
	"link-share/models"
	"os"
)

func GetCategories(c *fiber.Ctx) error {
	db := database.DB
	var categories []models.Category
	db.Find(&categories)
	return c.JSON(fiber.Map{
		"code": 200,
		"data": categories,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	db := database.DB
	user := c.Locals("user").(models.User)
	if user.Email != os.Getenv("ADMIN_EMAIL") {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"msg":  "Forbidden",
		})
	}
	var category models.Category
	if err := c.BodyParser(&category);err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	db.Create(&category)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Category created successfully",
		"data": category,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	user := c.Locals("user").(models.User)
	if user.Email != os.Getenv("ADMIN_EMAIL") {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"msg":  "Forbidden",
		})
	}
	var category models.Category
	var newCategory models.Category
	if err := c.BodyParser(&newCategory);err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	db.First(&category, id)
	category.Name = newCategory.Name
	db.Save(&category)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Category updated successfully",
		"data": category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	user := c.Locals("user").(models.User)
	if user.Email != os.Getenv("ADMIN_EMAIL") {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"msg":  "Forbidden",
		})
	}
	var category models.Category
	db.First(&category, id)
	db.Delete(&category)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Category deleted successfully",
	})
}
