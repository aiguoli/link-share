package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"link-share/database"
	"link-share/routes"
	"os"
)

func main() {
	app := fiber.New()

	file, _ := os.OpenFile("./link-share.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	app.Use(logger.New(logger.Config{
		Format:     "[${time}][${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     file,
	}))
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
