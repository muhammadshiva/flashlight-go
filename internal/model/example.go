package model

import "time"

// ExampleRequest is the request model for creating/updating example
type ExampleRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=1000"`
	Status      string `json:"status" validate:"omitempty,oneof=active inactive"`
}

// ExampleResponse is the response model for example
type ExampleResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ExampleFilterRequest is the filter model for listing examples
type ExampleFilterRequest struct {
	Name   string `query:"name"`
	Status string `query:"status"`
	Page   int    `query:"page" validate:"min=1"`
	Size   int    `query:"size" validate:"min=1,max=100"`
}
