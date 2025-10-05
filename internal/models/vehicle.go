package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Brand       string         `gorm:"type:varchar(100);not null" json:"brand"`
	Model       string         `gorm:"type:varchar(100);not null" json:"model"`
	VehicleType string         `gorm:"type:varchar(50);not null" json:"vehicle_type"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	CustomerVehicles []CustomerVehicle `gorm:"foreignKey:VehicleID" json:"customer_vehicles,omitempty"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}
