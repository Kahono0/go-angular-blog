package main

import (
	"blog/db"
	"blog/models"
	"blog/service"
	"fmt"
	"log"
	"time"
)

func SeedBlogs(count int) error {
	placeHolderImage := "https://picsum.photos/id/%d/200/300"

	for i := 1; i <= count; i++ {
		blog := models.Blog{
			Title:     fmt.Sprintf("Blog Title %d", i),
			Content:   fmt.Sprintf("This is the content of blog number %d.", i),
			Image:     fmt.Sprintf(placeHolderImage, i),
			CreatedAt: time.Now(),
		}

		if err := service.CreateBlog(&blog); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	dbName := "blog.db"
	if err := db.Connect(dbName); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := SeedBlogs(10); err != nil {
		log.Fatalf("Failed to seed blogs: %v", err)
	}

	log.Println("Successfully seeded blogs")
}
