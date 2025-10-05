package service

import (
	"context"
	"errors"
	"time"

	"flashlight-go/internal/dto"
	"flashlight-go/internal/models"
	"flashlight-go/internal/repository"

	"gorm.io/datatypes"
)

type PaymentService struct {
	paymentRepo   *repository.PaymentRepository
	workOrderRepo *repository.WorkOrderRepository
}

func NewPaymentService(
	paymentRepo *repository.PaymentRepository,
	workOrderRepo *repository.WorkOrderRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepo:   paymentRepo,
		workOrderRepo: workOrderRepo,
	}
}

func (s *PaymentService) Create(ctx context.Context, req dto.CreatePaymentRequest, cashierUserID, shiftID *uint) (*models.Payment, error) {
	// Verify work order exists
	workOrder, err := s.workOrderRepo.FindByID(ctx, req.WorkOrderID)
	if err != nil {
		return nil, errors.New("work order not found")
	}

	// Generate payment number
	paymentNumber, err := s.paymentRepo.GeneratePaymentNumber(ctx)
	if err != nil {
		return nil, err
	}

	// Calculate change for cash payments
	var changeAmount float64
	if req.Method == string(models.MethodCash) && req.AmountPaid > workOrder.TotalAmount {
		changeAmount = req.AmountPaid - workOrder.TotalAmount
	}

	// Create payment
	payment := &models.Payment{
		WorkOrderID:     req.WorkOrderID,
		CashierUserID:   cashierUserID,
		ShiftID:         shiftID,
		PaymentNumber:   paymentNumber,
		Method:          models.PaymentMethod(req.Method),
		Status:          models.PaymentStatusCompleted,
		AmountPaid:      req.AmountPaid,
		ChangeAmount:    changeAmount,
		ReferenceNumber: req.ReferenceNumber,
		PaidAt:          &time.Time{},
	}

	// Set paid at to now
	now := time.Now()
	payment.PaidAt = &now

	// Convert raw payload to JSON
	if req.RawPayload != nil {
		jsonData, _ := datatypes.NewJSONType(req.RawPayload).MarshalJSON()
		payment.RawPayload = jsonData
	}

	if err := s.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	// Check if work order is fully paid
	totalPaid, err := s.paymentRepo.GetTotalPaidForWorkOrder(ctx, req.WorkOrderID)
	if err == nil && totalPaid >= workOrder.TotalAmount {
		// Update work order status to completed
		workOrder.Status = models.StatusCompleted
		now := time.Now()
		workOrder.CompletedAt = &now
		_ = s.workOrderRepo.Update(ctx, workOrder)
	}

	return payment, nil
}

func (s *PaymentService) GetByID(ctx context.Context, id uint) (*models.Payment, error) {
	return s.paymentRepo.FindByID(ctx, id)
}

func (s *PaymentService) GetAll(ctx context.Context, page, perPage int) ([]models.Payment, *dto.PaginationMeta, error) {
	payments, total, err := s.paymentRepo.FindAll(ctx, page, perPage, "WorkOrder", "CashierUser")
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	meta := &dto.PaginationMeta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}

	return payments, meta, nil
}

func (s *PaymentService) Update(ctx context.Context, id uint, req dto.UpdatePaymentRequest) (*models.Payment, error) {
	payment, err := s.paymentRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Status != nil {
		payment.Status = models.PaymentStatus(*req.Status)
	}
	if req.ReferenceNumber != nil {
		payment.ReferenceNumber = req.ReferenceNumber
	}
	if req.RawPayload != nil {
		jsonData, _ := datatypes.NewJSONType(req.RawPayload).MarshalJSON()
		payment.RawPayload = jsonData
	}

	if err := s.paymentRepo.Update(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) Delete(ctx context.Context, id uint) error {
	return s.paymentRepo.Delete(ctx, id)
}

func (s *PaymentService) GetByWorkOrder(ctx context.Context, workOrderID uint) ([]models.Payment, error) {
	return s.paymentRepo.FindByWorkOrder(ctx, workOrderID)
}
