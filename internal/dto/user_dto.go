package dto

import "time"

type CreateUserRequest struct {
	Name                string  `json:"name" binding:"required"`
	Email               string  `json:"email" binding:"required,email"`
	PhoneNumber         string  `json:"phone_number"`
	Password            string  `json:"password" binding:"required,min=6"`
	Role                string  `json:"role" binding:"required,oneof=owner admin cashier staff customer"`
	MembershipTypeID    *uint   `json:"membership_type_id"`
	MembershipExpiresAt *string `json:"membership_expires_at"`
	Address             *string `json:"address"`
	City                *string `json:"city"`
	State               *string `json:"state"`
	PostalCode          *string `json:"postal_code"`
	Country             *string `json:"country"`
}

type UpdateUserRequest struct {
	Name                *string `json:"name"`
	Email               *string `json:"email,omitempty" binding:"omitempty,email"`
	PhoneNumber         *string `json:"phone_number"`
	Password            *string `json:"password,omitempty" binding:"omitempty,min=6"`
	Role                *string `json:"role,omitempty" binding:"omitempty,oneof=owner admin cashier staff customer"`
	MembershipTypeID    *uint   `json:"membership_type_id"`
	MembershipExpiresAt *string `json:"membership_expires_at"`
	Address             *string `json:"address"`
	City                *string `json:"city"`
	State               *string `json:"state"`
	PostalCode          *string `json:"postal_code"`
	Country             *string `json:"country"`
	IsActive            *bool   `json:"is_active"`
}

type UserResponse struct {
	ID                  uint       `json:"id"`
	Name                string     `json:"name"`
	Email               string     `json:"email"`
	PhoneNumber         string     `json:"phone_number"`
	Role                string     `json:"role"`
	MembershipTypeID    *uint      `json:"membership_type_id"`
	MembershipExpiresAt *time.Time `json:"membership_expires_at"`
	Address             *string    `json:"address"`
	City                *string    `json:"city"`
	State               *string    `json:"state"`
	PostalCode          *string    `json:"postal_code"`
	Country             *string    `json:"country"`
	ProfileImage        *string    `json:"profile_image"`
	IsActive            bool       `json:"is_active"`
	LastLoginAt         *time.Time `json:"last_login_at"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
