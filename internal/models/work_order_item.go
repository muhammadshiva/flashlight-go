package models

import (
	"time"
)

type WorkOrderItem struct {
	ID                    uint    `gorm:"primaryKey" json:"id"`
	WorkOrderID           uint    `gorm:"not null;index" json:"work_order_id"`
	ProductID             uint    `gorm:"not null;index" json:"product_id"`
	ProductNameSnapshot   string  `gorm:"type:varchar(255);not null" json:"product_name_snapshot"`
	PriceSnapshot         float64 `gorm:"type:decimal(15,2);not null" json:"price_snapshot"`
	Quantity              int     `gorm:"not null" json:"quantity"`
	Subtotal              float64 `gorm:"type:decimal(15,2);not null" json:"subtotal"`
	AssignedStaffUserID   *uint   `gorm:"index" json:"assigned_staff_user_id"`
	ItemNote              *string `gorm:"type:text" json:"item_note"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`

	// Relations
	WorkOrder      WorkOrder `gorm:"foreignKey:WorkOrderID" json:"work_order,omitempty"`
	Product        Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	AssignedStaff  *User     `gorm:"foreignKey:AssignedStaffUserID" json:"assigned_staff,omitempty"`
}

func (WorkOrderItem) TableName() string {
	return "work_order_items"
}
