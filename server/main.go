package main

import (
	"blog/db"
	"blog/router"
	"log"

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

	log.Fatal(app.Listen(":3000"))
}
