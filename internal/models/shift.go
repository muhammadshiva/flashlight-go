package models

import (
	"time"
)

type ShiftStatus string

const (
	ShiftStatusActive   ShiftStatus = "active"
	ShiftStatusClosed   ShiftStatus = "closed"
	ShiftStatusCanceled ShiftStatus = "canceled"
)

type Shift struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	UserID       uint        `gorm:"not null;index" json:"user_id"`
	StartTime    time.Time   `gorm:"not null" json:"start_time"`
	EndTime      *time.Time  `json:"end_time"`
	InitialCash  float64     `gorm:"type:decimal(15,2);default:0" json:"initial_cash"`
	FinalCash    float64     `gorm:"type:decimal(15,2);default:0" json:"final_cash"`
	TotalSales   float64     `gorm:"type:decimal(15,2);default:0" json:"total_sales"`
	Status       ShiftStatus `gorm:"type:varchar(20);not null" json:"status"`
	ReceivedFrom *string     `gorm:"type:varchar(255)" json:"received_from"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`

	// Relations
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	WorkOrders []WorkOrder `gorm:"foreignKey:ShiftID" json:"work_orders,omitempty"`
	Payments   []Payment   `gorm:"foreignKey:ShiftID" json:"payments,omitempty"`
}

func (Shift) TableName() string {
	return "shifts"
}
