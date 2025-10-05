package entity

// Example entity for demonstration
type Example struct {
	BaseEntity
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Status      string `gorm:"type:varchar(50);default:'active'"`
}

func (Example) TableName() string {
	return "examples"
}
