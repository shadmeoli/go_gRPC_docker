package main

import (
	"fmt"
	"html/template"

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

	// Serve HTML content
	app.Get("/", func(c *fiber.Ctx) error {
		// Retrieve all notes from the database
		var notes []Note
		db.Find(&notes)

		// Render the HTML template with the notes
		tmpl, err := template.New("notes").Parse(`index.html`)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return tmpl.Execute(c.Response().BodyWriter(), struct {
			Notes []Note
		}{Notes: notes})
	})

	// Create a new note
	app.Post("/api/notes", func(c *fiber.Ctx) error {
		var note Note
		if err := c.BodyParser(&note); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if result := db.Create(&note); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
		}

		// Retrieve all notes from the database
		var notes []Note
		db.Find(&notes)
		fmt.Println("Notes:", notes)

		// Render the HTML template with the updated notes
		tmpl, err := template.New("notes").Parse(`<!-- Your HTML template here, as shown above -->`)

		if err != nil {
			return err
		}

		return tmpl.Execute(c.Response().BodyWriter(), struct {
			Notes []Note
		}{Notes: notes})
	})

	// Start the Fiber app
	app.Listen(":8080")
}
