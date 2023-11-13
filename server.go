package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
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
		tmpl, err := template.New("notes").Parse(`<div id="notes-container" class="h-[50%] w-full flex flex-col justify-center items-center rounded-xl p-10">
            {{range .Notes}}
                <div class="w-full min-h-20 h-auto bg-gray-800 rounded-lg p-4">
                    <h1 class="text-center text-gray-400">{{.Text}}</h1>
                </div>
            {{else}}
                <h1 class="text-center text-gray-400 animate-pulse">Notes will appear here</h1>
            {{end}}
        </div>`)
		if err != nil {
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
