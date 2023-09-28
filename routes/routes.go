package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"link-share/handler"
	"link-share/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)

	// Link
	link := api.Group("/link")
	link.Get("/:id", handler.GetLink)
	link.Post("/", handler.CreateLink)
	link.Patch("/:id", middleware.Protected(), handler.UpdateLink)
	link.Delete("/:id", middleware.Protected(), handler.DeleteLink)
}
