package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkOrderSource string
type WorkOrderType string
type WorkOrderStatus string

const (
	SourceKiosk   WorkOrderSource = "kiosk"
	SourceCashier WorkOrderSource = "cashier"
	SourceOnline  WorkOrderSource = "online"

	TypeService WorkOrderType = "service"
	TypeRetail  WorkOrderType = "retail"
	TypeMix     WorkOrderType = "mix"

	StatusPending    WorkOrderStatus = "pending"
	StatusConfirmed  WorkOrderStatus = "confirmed"
	StatusInProgress WorkOrderStatus = "in_progress"
	StatusReady      WorkOrderStatus = "ready"
	StatusCompleted  WorkOrderStatus = "completed"
	StatusCancelled  WorkOrderStatus = "cancelled"
)

type WorkOrder struct {
	ID                   uint             `gorm:"primaryKey" json:"id"`
	OrderNumber          string           `gorm:"type:varchar(100);uniqueIndex;not null" json:"order_number"`
	Source               WorkOrderSource  `gorm:"type:varchar(20);not null" json:"source"`
	Type                 WorkOrderType    `gorm:"type:varchar(20);not null" json:"type"`
	CustomerUserID       *uint            `gorm:"index" json:"customer_user_id"`
	CustomerVehicleID    *uint            `gorm:"index" json:"customer_vehicle_id"`
	CashierUserID        *uint            `gorm:"index" json:"cashier_user_id"`
	ShiftID              *uint            `gorm:"index" json:"shift_id"`
	QueueNumber          *int             `json:"queue_number"`
	Status               WorkOrderStatus  `gorm:"type:varchar(20);not null;index" json:"status"`
	Notes                *string          `gorm:"type:text" json:"notes"`
	SpecialInstructions  *string          `gorm:"type:text" json:"special_instructions"`
	ConfirmedAt          *time.Time       `json:"confirmed_at"`
	StartedAt            *time.Time       `json:"started_at"`
	CompletedAt          *time.Time       `json:"completed_at"`
	Subtotal             float64          `gorm:"type:decimal(15,2);default:0" json:"subtotal"`
	DiscountAmount       float64          `gorm:"type:decimal(15,2);default:0" json:"discount_amount"`
	TaxAmount            float64          `gorm:"type:decimal(15,2);default:0" json:"tax_amount"`
	TotalAmount          float64          `gorm:"type:decimal(15,2);default:0" json:"total_amount"`
	CreatedAt            time.Time        `json:"created_at"`
	UpdatedAt            time.Time        `json:"updated_at"`
	DeletedAt            gorm.DeletedAt   `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	CustomerUser    *User              `gorm:"foreignKey:CustomerUserID" json:"customer_user,omitempty"`
	CustomerVehicle *CustomerVehicle   `gorm:"foreignKey:CustomerVehicleID" json:"customer_vehicle,omitempty"`
	CashierUser     *User              `gorm:"foreignKey:CashierUserID" json:"cashier_user,omitempty"`
	Shift           *Shift             `gorm:"foreignKey:ShiftID" json:"shift,omitempty"`
	Items           []WorkOrderItem    `gorm:"foreignKey:WorkOrderID" json:"items,omitempty"`
	Payments        []Payment          `gorm:"foreignKey:WorkOrderID" json:"payments,omitempty"`
}

func (WorkOrder) TableName() string {
	return "work_orders"
}
