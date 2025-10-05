package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductKind string

const (
	ProductKindService ProductKind = "service"
	ProductKindAddon   ProductKind = "addon"
	ProductKindRetail  ProductKind = "retail"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description *string        `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	Image       *string        `gorm:"type:varchar(255)" json:"image"`
	CategoryID  uint           `gorm:"not null;index" json:"category_id"`
	Kind        ProductKind    `gorm:"type:varchar(20);not null" json:"kind"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	IsPremium   bool           `gorm:"default:false" json:"is_premium"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Category       ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	WorkOrderItems []WorkOrderItem `gorm:"foreignKey:ProductID" json:"work_order_items,omitempty"`
}

func (Product) TableName() string {
	return "products"
}
