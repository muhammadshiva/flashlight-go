package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MembershipType struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Benefits  datatypes.JSON `gorm:"type:jsonb" json:"benefits"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Users []User `gorm:"foreignKey:MembershipTypeID" json:"users,omitempty"`
}

func (MembershipType) TableName() string {
	return "membership_types"
}
