package handlers

import (
	"blog/models"
	"blog/service"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

var serverURL = "http://localhost:3000"

func CreateBlog(c *fiber.Ctx) error {
	title := c.FormValue("title")
	slug := c.FormValue("slug")
	dateStr := c.FormValue("date")
	content := c.FormValue("content")

	var blogDate time.Time
	if dateStr == "" {
		blogDate = time.Now()
	} else {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid date format")
		}
		blogDate = t
	}

	file, err := c.FormFile("image")
	var filename string
	if err == nil && file != nil {
		uploadDir := "./uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create upload directory")
		}
		filename = fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveFile(file, filePath); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
		}
	}

	blog := models.Blog{
		Title:     title,
		Slug:      slug,
		CreatedAt: blogDate,
		Content:   content,
		Image:     fmt.Sprintf("%s/uploads/%s", serverURL, filename),
	}

	if err := service.CreateBlog(&blog); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to create blog: %v", err))
	}

	return c.Status(fiber.StatusCreated).JSON(blog)
}

func GetAllBlogs(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	itemsPerPage := c.QueryInt("itemsPerPage", 0)
	blogsResponse, err := service.QueryBlogs(page, itemsPerPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve blogs"})
	}

	return c.JSON(blogsResponse)
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

func SearchBlogs(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter is required"})
	}

	blogs, err := service.SearchBlogs(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to search blogs"})
	}

	return c.JSON(blogs)
}
