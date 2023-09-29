package handler

import (
	"github.com/gofiber/fiber/v2"
	"link-share/database"
	"link-share/models"
)

func Collect(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	linkId := c.Params("id")
	db := database.DB
	var link models.Link
	if err := db.First(&link, linkId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code": 404,
			"msg":  "Link not found",
		})
	}
	if err := db.Model(&user).Association("Links").Append(&link).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't collect link",
		})
	}
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Link collected successfully",
	})
}

func Uncollect(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	linkId := c.Params("id")
	db := database.DB
	var link models.Link
	if err := db.First(&link, linkId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code": 404,
			"msg":  "Link not found",
		})
	}
	if err := db.Model(&user).Association("Links").Delete(&link); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't uncollect link",
		})
	}
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Link uncollected successfully",
	})
}

func Collections(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	db := database.DB
	var links []models.Link
	if err := db.Model(&user).Association("Links").Find(&links); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't get collections",
		})
	}
	return c.JSON(fiber.Map{
		"code": 200,
		"data": links,
	})
}
