package dto

import "time"

// Generic CRUD DTOs for simpler entities

type CreateMembershipTypeRequest struct {
	Name     string      `json:"name" binding:"required"`
	Benefits interface{} `json:"benefits"`
	IsActive *bool       `json:"is_active"`
}

type UpdateMembershipTypeRequest struct {
	Name     *string     `json:"name"`
	Benefits interface{} `json:"benefits"`
	IsActive *bool       `json:"is_active"`
}

type CreateProductCategoryRequest struct {
	Name      string  `json:"name" binding:"required"`
	IconImage *string `json:"icon_image"`
	IsActive  *bool   `json:"is_active"`
}

type UpdateProductCategoryRequest struct {
	Name      *string `json:"name"`
	IconImage *string `json:"icon_image"`
	IsActive  *bool   `json:"is_active"`
}

type CreateProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Price       float64  `json:"price" binding:"required,gt=0"`
	Image       *string  `json:"image"`
	CategoryID  uint     `json:"category_id" binding:"required"`
	Kind        string   `json:"kind" binding:"required,oneof=service addon retail"`
	IsActive    *bool    `json:"is_active"`
	IsPremium   *bool    `json:"is_premium"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price,omitempty" binding:"omitempty,gt=0"`
	Image       *string  `json:"image"`
	CategoryID  *uint    `json:"category_id"`
	Kind        *string  `json:"kind,omitempty" binding:"omitempty,oneof=service addon retail"`
	IsActive    *bool    `json:"is_active"`
	IsPremium   *bool    `json:"is_premium"`
}

type CreateVehicleRequest struct {
	Brand       string `json:"brand" binding:"required"`
	Model       string `json:"model" binding:"required"`
	VehicleType string `json:"vehicle_type" binding:"required"`
}

type UpdateVehicleRequest struct {
	Brand       *string `json:"brand"`
	Model       *string `json:"model"`
	VehicleType *string `json:"vehicle_type"`
}

type CreateCustomerVehicleRequest struct {
	CustomerID   uint   `json:"customer_id" binding:"required"`
	VehicleID    uint   `json:"vehicle_id" binding:"required"`
	LicensePlate string `json:"license_plate" binding:"required"`
}

type UpdateCustomerVehicleRequest struct {
	VehicleID    *uint   `json:"vehicle_id"`
	LicensePlate *string `json:"license_plate"`
}

type CreatePaymentRequest struct {
	WorkOrderID     uint        `json:"work_order_id" binding:"required"`
	Method          string      `json:"method" binding:"required,oneof=cash qris transfer e_wallet"`
	AmountPaid      float64     `json:"amount_paid" binding:"required,gt=0"`
	ReferenceNumber *string     `json:"reference_number"`
	RawPayload      interface{} `json:"raw_payload"`
}

type UpdatePaymentRequest struct {
	Status          *string     `json:"status,omitempty" binding:"omitempty,oneof=pending completed failed refunded"`
	ReferenceNumber *string     `json:"reference_number"`
	RawPayload      interface{} `json:"raw_payload"`
}

type CreateShiftRequest struct {
	InitialCash  float64 `json:"initial_cash"`
	ReceivedFrom *string `json:"received_from"`
}

type CloseShiftRequest struct {
	FinalCash float64 `json:"final_cash" binding:"required"`
}

type CreateDeviceFCMTokenRequest struct {
	UserID     uint    `json:"user_id" binding:"required"`
	DeviceID   string  `json:"device_id" binding:"required"`
	FCMToken   string  `json:"fcm_token" binding:"required"`
	DeviceName *string `json:"device_name"`
	DeviceType *string `json:"device_type"`
}

type UpdateDeviceFCMTokenRequest struct {
	FCMToken   *string `json:"fcm_token"`
	DeviceName *string `json:"device_name"`
	DeviceType *string `json:"device_type"`
}

// Generic response types
type GenericResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
