package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"link-share/database"
	"link-share/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Working!")
	})
	routes.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
