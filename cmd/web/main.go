package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/yourusername/go-skeleton/internal/delivery/http"
	"github.com/yourusername/go-skeleton/internal/gateway"
	"github.com/yourusername/go-skeleton/internal/repository"
	"github.com/yourusername/go-skeleton/internal/usecase"
	"github.com/yourusername/go-skeleton/pkg/config"
	"github.com/yourusername/go-skeleton/pkg/database"
	"github.com/yourusername/go-skeleton/pkg/logger"
	"github.com/yourusername/go-skeleton/pkg/validator"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger.InitLogger(&cfg.Log)
	log := logger.GetLogger()

	// Initialize validator
	validator.InitValidator()

	// Initialize database
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repository
	exampleRepo := repository.NewExampleRepository(db)

	// Initialize gateway
	exampleGateway := gateway.NewExampleGateway("http://external-api.example.com")

	// Initialize use case
	exampleUseCase := usecase.NewExampleUseCase(exampleRepo, exampleGateway)

	// Initialize handler
	exampleHandler := http.NewExampleHandler(exampleUseCase)

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// Setup routes
	http.SetupRouter(app, exampleHandler)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Infof("Starting web server on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
