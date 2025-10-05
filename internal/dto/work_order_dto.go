package dto

import "time"

type CreateWorkOrderRequest struct {
	Source              string                     `json:"source" binding:"required,oneof=kiosk cashier online"`
	Type                string                     `json:"type" binding:"required,oneof=service retail mix"`
	CustomerUserID      *uint                      `json:"customer_user_id"`
	CustomerVehicleID   *uint                      `json:"customer_vehicle_id"`
	Notes               *string                    `json:"notes"`
	SpecialInstructions *string                    `json:"special_instructions"`
	Items               []CreateWorkOrderItemRequest `json:"items" binding:"required,min=1"`
}

type CreateWorkOrderItemRequest struct {
	ProductID           uint    `json:"product_id" binding:"required"`
	Quantity            int     `json:"quantity" binding:"required,min=1"`
	AssignedStaffUserID *uint   `json:"assigned_staff_user_id"`
	ItemNote            *string `json:"item_note"`
}

type UpdateWorkOrderRequest struct {
	Status              *string `json:"status,omitempty" binding:"omitempty,oneof=pending confirmed in_progress ready completed cancelled"`
	CashierUserID       *uint   `json:"cashier_user_id"`
	ShiftID             *uint   `json:"shift_id"`
	Notes               *string `json:"notes"`
	SpecialInstructions *string `json:"special_instructions"`
	DiscountAmount      *float64 `json:"discount_amount"`
	TaxAmount           *float64 `json:"tax_amount"`
}

type WorkOrderResponse struct {
	ID                  uint                    `json:"id"`
	OrderNumber         string                  `json:"order_number"`
	Source              string                  `json:"source"`
	Type                string                  `json:"type"`
	CustomerUserID      *uint                   `json:"customer_user_id"`
	CustomerVehicleID   *uint                   `json:"customer_vehicle_id"`
	CashierUserID       *uint                   `json:"cashier_user_id"`
	ShiftID             *uint                   `json:"shift_id"`
	QueueNumber         *int                    `json:"queue_number"`
	Status              string                  `json:"status"`
	Notes               *string                 `json:"notes"`
	SpecialInstructions *string                 `json:"special_instructions"`
	ConfirmedAt         *time.Time              `json:"confirmed_at"`
	StartedAt           *time.Time              `json:"started_at"`
	CompletedAt         *time.Time              `json:"completed_at"`
	Subtotal            float64                 `json:"subtotal"`
	DiscountAmount      float64                 `json:"discount_amount"`
	TaxAmount           float64                 `json:"tax_amount"`
	TotalAmount         float64                 `json:"total_amount"`
	Items               []WorkOrderItemResponse `json:"items,omitempty"`
	CreatedAt           time.Time               `json:"created_at"`
	UpdatedAt           time.Time               `json:"updated_at"`
}

type WorkOrderItemResponse struct {
	ID                  uint      `json:"id"`
	WorkOrderID         uint      `json:"work_order_id"`
	ProductID           uint      `json:"product_id"`
	ProductNameSnapshot string    `json:"product_name_snapshot"`
	PriceSnapshot       float64   `json:"price_snapshot"`
	Quantity            int       `json:"quantity"`
	Subtotal            float64   `json:"subtotal"`
	AssignedStaffUserID *uint     `json:"assigned_staff_user_id"`
	ItemNote            *string   `json:"item_note"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
