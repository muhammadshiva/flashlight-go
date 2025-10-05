package entity

import (
	"time"
)

// BaseEntity contains common fields for all entities
type BaseEntity struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
