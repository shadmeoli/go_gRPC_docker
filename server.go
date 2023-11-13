package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Note struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Text string `json:"text"`
}

func main() {
	// Create a Fiber app
	app := fiber.New()

	// Initialize the database
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Note{})

	// Middleware
	app.Use(logger.New())

	// Create a new note
	app.Post("/api/notes", func(c *fiber.Ctx) error {
		var note Note
		if err := c.BodyParser(&note); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if result := db.Create(&note); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(note)
	})

	// Delete a note by ID
	app.Delete("/api/notes/:id", func(c *fiber.Ctx) error {
		noteID := c.Params("id")
		var note Note

		if result := db.First(&note, noteID); result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}

		if result := db.Delete(&note); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
		}

		return c.SendStatus(fiber.StatusOK)
	})

	// Start the Fiber app
	app.Listen(":8080")
}
