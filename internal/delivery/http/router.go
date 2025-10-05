package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRouter(app *fiber.App, exampleHandler *ExampleHandler) {
	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// API routes
	api := app.Group("/api/v1")

	// Example routes
	examples := api.Group("/examples")
	examples.Post("/", exampleHandler.Create)
	examples.Get("/", exampleHandler.GetAll)
	examples.Get("/:id", exampleHandler.GetByID)
	examples.Put("/:id", exampleHandler.Update)
	examples.Delete("/:id", exampleHandler.Delete)
}
