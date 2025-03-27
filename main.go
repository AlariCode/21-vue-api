package main

import (
	"bookmark-api/internal/app/bookmarks"
	"bookmark-api/internal/app/categories"
	"bookmark-api/internal/app/profile"
	"bookmark-api/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}

	fiberApp := fiber.New()

	api := fiberApp.Group("/api")

	// Profile route
	api.Get("/profile", profile.Get)

	// Categories routes
	api.Post("/categories", categories.Create)
	api.Get("/categories", categories.GetAll)
	api.Put("/categories/:id", categories.Update)
	api.Delete("/categories/:id", categories.Delete)

	// Bookmarks routes
	api.Post("/bookmarks", bookmarks.Create)
	api.Delete("/bookmarks/:id", bookmarks.Delete)
	api.Get("/categories/:categoryId/bookmarks", bookmarks.GetByCategory)

	log.Fatal(fiberApp.Listen(":3000"))
}
