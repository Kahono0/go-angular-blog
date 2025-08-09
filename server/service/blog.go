package service

import (
	"blog/db"
	"blog/models"
	"errors"
	"strings"
)

func generateSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func contentWordCount(content string) int {
	words := strings.Fields(content)
	return len(words)
}

var WordsPerMinute = 200

var (
	ErrInvalidInput = errors.New("invalid input: title and content cannot be empty")
)

func estimateReadingTime(content string) int {
	wordCount := contentWordCount(content)
	if wordCount == 0 {
		return 0
	}
	return (wordCount + WordsPerMinute - 1) / WordsPerMinute
}

func CreateBlog(b *models.Blog) error {
	if b.Title == "" || b.Content == "" {
		return ErrInvalidInput
	}

	b.Slug = generateSlug(b.Title)
	b.ReadingTime = estimateReadingTime(b.Content)

	if err := db.DB.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := db.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogBySlug(slug string) (*models.Blog, error) {
	var blog models.Blog
	if err := db.DB.Where("slug = ?", slug).First(&blog).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}
