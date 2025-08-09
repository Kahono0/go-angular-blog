package router

import (
	"blog/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetUpRoutes(app *fiber.App) {
	app.Use(cors.New())

	// serve static files
	app.Static("/uploads", "./uploads")

	app.Post("/blog", handlers.CreateBlog)
	app.Get("/blogs", handlers.GetAllBlogs)
	app.Get("/blogs/:slug", handlers.GetBlogBySlug)
}
