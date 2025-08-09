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

var DefaultLimit = 10

type QueryBlogsResponse struct {
	Data        []models.Blog `json:"data"`
	TotalPages  int           `json:"totalPages"`
	CurrentPage int           `json:"currentPage"`
}

func QueryBlogs(page, limit int) (*QueryBlogsResponse, error) {
	if limit <= 0 {
		limit = DefaultLimit
	}
	if page < 1 {
		page = 1
	}

	var blogs []models.Blog
	var total int64

	if err := db.DB.Model(&models.Blog{}).Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if err := db.DB.Offset(offset).Limit(limit).Find(&blogs).Error; err != nil {
		return nil, err
	}

	totalPages := (int(total) + limit - 1) / limit

	return &QueryBlogsResponse{
		Data:        blogs,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}

func GetBlogBySlug(slug string) (*models.Blog, error) {
	var blog models.Blog
	if err := db.DB.Where("slug = ?", slug).First(&blog).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}

func SearchBlogs(query string) ([]models.Blog, error) {
	var blogs []models.Blog
	if err := db.DB.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%").Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}
