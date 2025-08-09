package router

import (
	"blog/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetUpRoutes(app *fiber.App) {
	app.Use(cors.New())

	app.Static("/", "./web")

	// serve static files
	app.Static("/uploads", "./uploads")

	api := app.Group("/api")
	api.Post("/blog", handlers.CreateBlog)
	api.Get("/blogs/search", handlers.SearchBlogs)
	api.Get("/blogs/:slug", handlers.GetBlogBySlug)
	api.Get("/blogs", handlers.GetAllBlogs)
}
