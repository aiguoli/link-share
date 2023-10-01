package handler

import (
	"github.com/gofiber/fiber/v2"
	"link-share/database"
	"link-share/models"
	"time"
)

func GetLink(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var link models.Link
	if err := db.Find(&link, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"msg":  "Link not found",
		})
	}
	return c.JSON(fiber.Map{
		"code": 200,
		"data": link,
	})
}

func CreateLink(c *fiber.Ctx) error {
	db := database.DB
	link := new(models.Link)
	if err := c.BodyParser(link); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	if c.Locals("user") != nil {
		link.UserID = c.Locals("user").(models.User).ID
	}
	link.ExpirationDate, _ = time.Parse("2006-01-02", c.FormValue("expiration_date"))
	db.Create(&link)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Link created successfully",
		"data": link,
	})
}

func UpdateLink(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var link models.Link
	var newLink models.Link

	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	user := c.Locals("user").(models.User)
	if err := db.First(&link, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"msg":  "Link not found",
		})
	}
	db.Model(&link).Omit("UserID").Updates(newLink)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Link updated successfully",
		"data": link,
	})
}

func DeleteLink(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var link models.Link
	user := c.Locals("user").(models.User)
	if err := db.First(&link, "id = ? AND user_id = ?", id, user.ID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code": 404,
			"msg":  "Link not found",
		})
	}
	db.Delete(&link)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Link deleted successfully",
	})
}
