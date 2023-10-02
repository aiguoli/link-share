package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"link-share/database"
	"link-share/models"
)

func Collect(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	linkId := c.Params("id")
	db := database.DB
	var link models.Link
	var user models.User
	if err := db.Find(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't uncollect link",
		})
	}
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
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	linkId := c.Params("id")
	db := database.DB
	var user models.User
	if err := db.Find(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't uncollect link",
		})
	}
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
	db := database.DB
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	var user models.User
	if err := db.Find(&user, userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "User not found!",
		})
	}
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
