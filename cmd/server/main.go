package main

import (
	"fmt"
	"log"

	"flashlight-go/config"
	"flashlight-go/internal/database"
	"flashlight-go/internal/handler"
	"flashlight-go/internal/repository"
	"flashlight-go/internal/routes"
	"flashlight-go/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Create indexes
	if err := database.CreateIndexes(db); err != nil {
		log.Fatal("Failed to create indexes:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	workOrderRepo := repository.NewWorkOrderRepository(db)
	workOrderItemRepo := repository.NewWorkOrderItemRepository(db)
	productRepo := repository.NewProductRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	shiftRepo := repository.NewShiftRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	workOrderService := service.NewWorkOrderService(workOrderRepo, workOrderItemRepo, productRepo, db)
	paymentService := service.NewPaymentService(paymentRepo, workOrderRepo)
	shiftService := service.NewShiftService(shiftRepo, paymentRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	workOrderHandler := handler.NewWorkOrderHandler(workOrderService)

	// Setup routes
	router := routes.NewRouter(userHandler, workOrderHandler)
	r := router.Setup()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}

	// Suppress unused variable warnings (these will be used when adding more routes)
	_ = paymentService
	_ = shiftService
}
