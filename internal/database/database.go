package database

import (
	"fmt"
	"log"

	"flashlight-go/config"
	"flashlight-go/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Database.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.MembershipType{},
		&models.DeviceFCMToken{},
		&models.Vehicle{},
		&models.CustomerVehicle{},
		&models.ProductCategory{},
		&models.Product{},
		&models.WorkOrder{},
		&models.WorkOrderItem{},
		&models.Payment{},
		&models.Shift{},
	)

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func CreateIndexes(db *gorm.DB) error {
	log.Println("Creating database indexes...")

	// Work Orders indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_orders_order_number ON work_orders(order_number)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_orders_customer_user_id ON work_orders(customer_user_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_orders_shift_id ON work_orders(shift_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_orders_status_source_created ON work_orders(status, source, created_at)")

	// Work Order Items indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_order_items_work_order_id ON work_order_items(work_order_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_work_order_items_product_id ON work_order_items(product_id)")

	// Payments indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_payments_work_order_id ON payments(work_order_id)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_payments_payment_number ON payments(payment_number)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_payments_shift_id ON payments(shift_id)")

	// Customer Vehicles indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_customer_vehicles_customer_id_license ON customer_vehicles(customer_id, license_plate)")

	// Products indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_products_category_kind_active ON products(category_id, kind, is_active)")

	log.Println("Database indexes created successfully")
	return nil
}
