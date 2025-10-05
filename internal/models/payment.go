package models

import (
	"time"

	"gorm.io/datatypes"
)

type PaymentMethod string
type PaymentStatus string

const (
	MethodCash    PaymentMethod = "cash"
	MethodQRIS    PaymentMethod = "qris"
	MethodTransfer PaymentMethod = "transfer"
	MethodEWallet  PaymentMethod = "e_wallet"

	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

type Payment struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	WorkOrderID     uint           `gorm:"not null;index" json:"work_order_id"`
	CashierUserID   *uint          `gorm:"index" json:"cashier_user_id"`
	ShiftID         *uint          `gorm:"index" json:"shift_id"`
	PaymentNumber   string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"payment_number"`
	Method          PaymentMethod  `gorm:"type:varchar(20);not null" json:"method"`
	Status          PaymentStatus  `gorm:"type:varchar(20);not null" json:"status"`
	AmountPaid      float64        `gorm:"type:decimal(15,2);not null" json:"amount_paid"`
	ChangeAmount    float64        `gorm:"type:decimal(15,2);default:0" json:"change_amount"`
	ReferenceNumber *string        `gorm:"type:varchar(255)" json:"reference_number"`
	RawPayload      datatypes.JSON `gorm:"type:jsonb" json:"raw_payload"`
	PaidAt          *time.Time     `json:"paid_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`

	// Relations
	WorkOrder   WorkOrder `gorm:"foreignKey:WorkOrderID" json:"work_order,omitempty"`
	CashierUser *User     `gorm:"foreignKey:CashierUserID" json:"cashier_user,omitempty"`
	Shift       *Shift    `gorm:"foreignKey:ShiftID" json:"shift,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}
