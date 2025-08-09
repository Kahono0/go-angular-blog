package main

import (
	"blog/db"
	"blog/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dbName := "blog.db"
	if err := db.Connect(dbName); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	// Set up routes
	router.SetUpRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
