package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerVehicle struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CustomerID   uint           `gorm:"not null;index" json:"customer_id"`
	VehicleID    uint           `gorm:"not null;index" json:"vehicle_id"`
	LicensePlate string         `gorm:"type:varchar(50);not null" json:"license_plate"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Customer   User        `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Vehicle    Vehicle     `gorm:"foreignKey:VehicleID" json:"vehicle,omitempty"`
	WorkOrders []WorkOrder `gorm:"foreignKey:CustomerVehicleID" json:"work_orders,omitempty"`
}

func (CustomerVehicle) TableName() string {
	return "customer_vehicles"
}
