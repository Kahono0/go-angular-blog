package handlers

import (
	"blog/models"
	"blog/service"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := service.CreateBlog(&blog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create blog"})
	}
	return c.Status(fiber.StatusCreated).JSON(blog)
}

func GetAllBlogs(c *fiber.Ctx) error {
	blogs, err := service.GetAllBlogs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve blogs"})
	}

	return c.JSON(blogs)
}

func GetBlogBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Slug is required"})
	}

	blog, err := service.GetBlogBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
	}

	return c.JSON(blog)
}
