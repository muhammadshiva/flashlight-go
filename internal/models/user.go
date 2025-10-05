package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleOwner    UserRole = "owner"
	RoleAdmin    UserRole = "admin"
	RoleCashier  UserRole = "cashier"
	RoleStaff    UserRole = "staff"
	RoleCustomer UserRole = "customer"
)

type User struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	Name                string         `gorm:"type:varchar(255);not null" json:"name"`
	Email               string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PhoneNumber         string         `gorm:"type:varchar(50)" json:"phone_number"`
	Password            string         `gorm:"type:varchar(255);not null" json:"-"`
	Role                UserRole       `gorm:"type:varchar(20);not null" json:"role"`
	MembershipTypeID    *uint          `gorm:"index" json:"membership_type_id"`
	MembershipExpiresAt *time.Time     `json:"membership_expires_at"`
	Address             *string        `gorm:"type:text" json:"address"`
	City                *string        `gorm:"type:varchar(100)" json:"city"`
	State               *string        `gorm:"type:varchar(100)" json:"state"`
	PostalCode          *string        `gorm:"type:varchar(20)" json:"postal_code"`
	Country             *string        `gorm:"type:varchar(100)" json:"country"`
	ProfileImage        *string        `gorm:"type:varchar(255)" json:"profile_image"`
	IsActive            bool           `gorm:"default:true" json:"is_active"`
	LastLoginAt         *time.Time     `json:"last_login_at"`
	FCMToken            *string        `gorm:"type:varchar(255)" json:"fcm_token"`
	EmailVerifiedAt     *time.Time     `json:"email_verified_at"`
	RememberToken       *string        `gorm:"type:varchar(100)" json:"-"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	MembershipType    *MembershipType    `gorm:"foreignKey:MembershipTypeID" json:"membership_type,omitempty"`
	DeviceFCMTokens   []DeviceFCMToken   `gorm:"foreignKey:UserID" json:"device_fcm_tokens,omitempty"`
	CustomerVehicles  []CustomerVehicle  `gorm:"foreignKey:CustomerID" json:"customer_vehicles,omitempty"`
	WorkOrders        []WorkOrder        `gorm:"foreignKey:CustomerUserID" json:"work_orders,omitempty"`
	HandledWorkOrders []WorkOrder        `gorm:"foreignKey:CashierUserID" json:"handled_work_orders,omitempty"`
	Payments          []Payment          `gorm:"foreignKey:CashierUserID" json:"payments,omitempty"`
	Shifts            []Shift            `gorm:"foreignKey:UserID" json:"shifts,omitempty"`
}

func (User) TableName() string {
	return "users"
}
