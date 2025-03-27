package categories

import (
	"bookmark-api/internal/database"
	"bookmark-api/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	result, err := database.DB.Exec("INSERT INTO categories (name, alias) VALUES (?, ?)",
		category.Name, category.Alias)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := result.LastInsertId()
	category.ID = id

	return c.JSON(category)
}

func GetAll(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, name, alias FROM categories")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Alias); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		categories = append(categories, category)
	}

	return c.JSON(categories)
}

func Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err = database.DB.Exec("UPDATE categories SET name = ?, alias = ? WHERE id = ?",
		category.Name, category.Alias, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	category.ID = id
	return c.JSON(category)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(204)
}
