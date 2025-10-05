package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	IconImage *string        `gorm:"type:varchar(255)" json:"icon_image"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}
