package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/go-skeleton/internal/delivery/kafka"
	"github.com/yourusername/go-skeleton/internal/gateway"
	"github.com/yourusername/go-skeleton/internal/repository"
	"github.com/yourusername/go-skeleton/internal/usecase"
	"github.com/yourusername/go-skeleton/pkg/config"
	"github.com/yourusername/go-skeleton/pkg/database"
	pkgKafka "github.com/yourusername/go-skeleton/pkg/kafka"
	"github.com/yourusername/go-skeleton/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger.InitLogger(&cfg.Log)
	logInstance := logger.GetLogger()

	// Initialize database
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		logInstance.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repository
	exampleRepo := repository.NewExampleRepository(db)

	// Initialize gateway
	exampleGateway := gateway.NewExampleGateway("http://external-api.example.com")

	// Initialize use case
	exampleUseCase := usecase.NewExampleUseCase(exampleRepo, exampleGateway)

	// Initialize Kafka consumer handler
	consumerHandler := kafka.NewExampleConsumerHandler(exampleUseCase, logInstance)

	// Get topics from config
	topics := []string{cfg.Kafka.Topics["example"]}

	// Initialize Kafka consumer
	consumer, err := pkgKafka.NewConsumer(&cfg.Kafka, topics, consumerHandler, logInstance)
	if err != nil {
		logInstance.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logInstance.Info("Starting Kafka consumer worker...")
		if err := consumer.Start(ctx); err != nil {
			logInstance.Errorf("Kafka consumer error: %v", err)
		}
	}()

	<-sigterm
	logInstance.Info("Shutting down worker gracefully...")
	cancel()
}
