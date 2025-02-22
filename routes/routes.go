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
	user := api.Group("/users")
	user.Get("/:id", handler.GetUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)

	// Link
	link := api.Group("/links")
	link.Get("/", handler.GetLinks)
	link.Get("/:id", handler.GetLink)
	link.Post("/", handler.CreateLink)
	link.Post("/:id/visit", middleware.Protected(), handler.IncrementViews)
	link.Patch("/:id", middleware.Protected(), handler.UpdateLink)
	link.Delete("/:id", middleware.Protected(), handler.DeleteLink)

	// Collection
	collection := api.Group("/")
	collection.Post("/collect/:id", middleware.Protected(), handler.Collect)
	collection.Post("/uncollect/:id", middleware.Protected(), handler.Uncollect)
	collection.Get("/collections", middleware.Protected(), handler.Collections)

	// Category
	category := api.Group("/categories")
	category.Get("/", handler.GetCategories)
	category.Post("/", middleware.Protected(), handler.CreateCategory)
	category.Patch("/:id", middleware.Protected(), handler.UpdateCategory)
	category.Delete("/:id", middleware.Protected(), handler.DeleteCategory)
}
