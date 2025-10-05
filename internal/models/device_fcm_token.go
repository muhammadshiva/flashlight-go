package models

import (
	"time"
)

type DeviceFCMToken struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `gorm:"not null;index" json:"user_id"`
	DeviceID   string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"device_id"`
	FCMToken   string     `gorm:"type:varchar(255);not null" json:"fcm_token"`
	DeviceName *string    `gorm:"type:varchar(255)" json:"device_name"`
	DeviceType *string    `gorm:"type:varchar(50)" json:"device_type"`
	LastUsedAt *time.Time `json:"last_used_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (DeviceFCMToken) TableName() string {
	return "device_fcm_tokens"
}
