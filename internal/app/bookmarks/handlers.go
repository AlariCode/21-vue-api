package bookmarks

import (
	"bookmark-api/internal/database"
	"bookmark-api/internal/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/opengraph/v2"
)

func Create(c *fiber.Ctx) error {
	bookmark := new(models.Bookmark)
	if err := c.BodyParser(bookmark); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)", bookmark.CategoryID).Scan(&exists)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(400).JSON(fiber.Map{"error": "Category not found"})
	}

	og, err := opengraph.Fetch(bookmark.URL)
	if err == nil {
		bookmark.Title = og.Title
		if len(og.Image) > 0 {
			bookmark.Image = og.Image[0].URL
		}
	}

	result, err := database.DB.Exec("INSERT INTO bookmarks (category_id, url, title, image) VALUES (?, ?, ?, ?)",
		bookmark.CategoryID, bookmark.URL, bookmark.Title, bookmark.Image)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := result.LastInsertId()
	bookmark.ID = id

	return c.JSON(bookmark)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM bookmarks WHERE id = ?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}

func GetByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("categoryId")
	sortBy := c.Query("sort", "date") // по умолчанию сортировка по дате

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)", categoryID).Scan(&exists)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if !exists {
		return c.Status(400).JSON(fiber.Map{"error": "Category not found"})
	}

	var orderBy string
	switch sortBy {
	case "title":
		orderBy = "title"
	case "date":
		orderBy = "created_at DESC"
	default:
		return c.Status(400).JSON(fiber.Map{"error": "Invalid sort parameter. Use 'date' or 'title'"})
	}

	query := fmt.Sprintf("SELECT id, category_id, url, title, image, created_at FROM bookmarks WHERE category_id = ? ORDER BY %s", orderBy)
	rows, err := database.DB.Query(query, categoryID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	bookmarks := make([]models.Bookmark, 0)
	for rows.Next() {
		var bookmark models.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.CategoryID, &bookmark.URL, &bookmark.Title, &bookmark.Image, &bookmark.CreatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return c.JSON(bookmarks)
}
