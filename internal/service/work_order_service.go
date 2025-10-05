package service

import (
	"context"
	"errors"
	"time"

	"flashlight-go/internal/dto"
	"flashlight-go/internal/models"
	"flashlight-go/internal/repository"

	"gorm.io/gorm"
)

type WorkOrderService struct {
	workOrderRepo     *repository.WorkOrderRepository
	workOrderItemRepo *repository.WorkOrderItemRepository
	productRepo       *repository.ProductRepository
	db                *gorm.DB
}

func NewWorkOrderService(
	workOrderRepo *repository.WorkOrderRepository,
	workOrderItemRepo *repository.WorkOrderItemRepository,
	productRepo *repository.ProductRepository,
	db *gorm.DB,
) *WorkOrderService {
	return &WorkOrderService{
		workOrderRepo:     workOrderRepo,
		workOrderItemRepo: workOrderItemRepo,
		productRepo:       productRepo,
		db:                db,
	}
}

func (s *WorkOrderService) Create(ctx context.Context, req dto.CreateWorkOrderRequest, cashierUserID *uint) (*dto.WorkOrderResponse, error) {
	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Generate order number
	orderNumber, err := s.workOrderRepo.GenerateOrderNumber(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get next queue number
	queueNumber, err := s.workOrderRepo.GetNextQueueNumber(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create work order
	workOrder := &models.WorkOrder{
		OrderNumber:         orderNumber,
		Source:              models.WorkOrderSource(req.Source),
		Type:                models.WorkOrderType(req.Type),
		CustomerUserID:      req.CustomerUserID,
		CustomerVehicleID:   req.CustomerVehicleID,
		CashierUserID:       cashierUserID,
		QueueNumber:         &queueNumber,
		Status:              models.StatusPending,
		Notes:               req.Notes,
		SpecialInstructions: req.SpecialInstructions,
	}

	if err := tx.Create(workOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create work order items and calculate totals
	var subtotal float64
	for _, itemReq := range req.Items {
		product, err := s.productRepo.FindByID(ctx, itemReq.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("product not found")
		}

		itemSubtotal := product.Price * float64(itemReq.Quantity)
		subtotal += itemSubtotal

		item := &models.WorkOrderItem{
			WorkOrderID:         workOrder.ID,
			ProductID:           product.ID,
			ProductNameSnapshot: product.Name,
			PriceSnapshot:       product.Price,
			Quantity:            itemReq.Quantity,
			Subtotal:            itemSubtotal,
			AssignedStaffUserID: itemReq.AssignedStaffUserID,
			ItemNote:            itemReq.ItemNote,
		}

		if err := tx.Create(item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Update work order totals
	workOrder.Subtotal = subtotal
	workOrder.TotalAmount = subtotal

	if err := tx.Save(workOrder).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch complete work order with items
	result, err := s.workOrderRepo.FindWithItems(ctx, workOrder.ID)
	if err != nil {
		return nil, err
	}

	return s.toResponse(result), nil
}

func (s *WorkOrderService) GetByID(ctx context.Context, id uint) (*dto.WorkOrderResponse, error) {
	workOrder, err := s.workOrderRepo.FindWithItems(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(workOrder), nil
}

func (s *WorkOrderService) GetAll(ctx context.Context, page, perPage int) ([]dto.WorkOrderResponse, *dto.PaginationMeta, error) {
	workOrders, total, err := s.workOrderRepo.FindAll(ctx, page, perPage, "Items", "CustomerUser", "CustomerVehicle", "CustomerVehicle.Vehicle")
	if err != nil {
		return nil, nil, err
	}

	responses := make([]dto.WorkOrderResponse, len(workOrders))
	for i, wo := range workOrders {
		responses[i] = *s.toResponse(&wo)
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

	return responses, meta, nil
}

func (s *WorkOrderService) Update(ctx context.Context, id uint, req dto.UpdateWorkOrderRequest) (*dto.WorkOrderResponse, error) {
	workOrder, err := s.workOrderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Status != nil {
		workOrder.Status = models.WorkOrderStatus(*req.Status)

		// Update timestamps based on status
		now := time.Now()
		switch workOrder.Status {
		case models.StatusConfirmed:
			workOrder.ConfirmedAt = &now
		case models.StatusInProgress:
			workOrder.StartedAt = &now
		case models.StatusCompleted:
			workOrder.CompletedAt = &now
		}
	}

	if req.CashierUserID != nil {
		workOrder.CashierUserID = req.CashierUserID
	}
	if req.ShiftID != nil {
		workOrder.ShiftID = req.ShiftID
	}
	if req.Notes != nil {
		workOrder.Notes = req.Notes
	}
	if req.SpecialInstructions != nil {
		workOrder.SpecialInstructions = req.SpecialInstructions
	}
	if req.DiscountAmount != nil {
		workOrder.DiscountAmount = *req.DiscountAmount
	}
	if req.TaxAmount != nil {
		workOrder.TaxAmount = *req.TaxAmount
	}

	// Recalculate total
	workOrder.TotalAmount = workOrder.Subtotal - workOrder.DiscountAmount + workOrder.TaxAmount

	if err := s.workOrderRepo.Update(ctx, workOrder); err != nil {
		return nil, err
	}

	result, err := s.workOrderRepo.FindWithItems(ctx, workOrder.ID)
	if err != nil {
		return nil, err
	}

	return s.toResponse(result), nil
}

func (s *WorkOrderService) Delete(ctx context.Context, id uint) error {
	return s.workOrderRepo.Delete(ctx, id)
}

func (s *WorkOrderService) GetByStatus(ctx context.Context, status string) ([]dto.WorkOrderResponse, error) {
	workOrders, err := s.workOrderRepo.FindByStatus(ctx, models.WorkOrderStatus(status))
	if err != nil {
		return nil, err
	}

	responses := make([]dto.WorkOrderResponse, len(workOrders))
	for i, wo := range workOrders {
		responses[i] = *s.toResponse(&wo)
	}

	return responses, nil
}

func (s *WorkOrderService) toResponse(wo *models.WorkOrder) *dto.WorkOrderResponse {
	response := &dto.WorkOrderResponse{
		ID:                  wo.ID,
		OrderNumber:         wo.OrderNumber,
		Source:              string(wo.Source),
		Type:                string(wo.Type),
		CustomerUserID:      wo.CustomerUserID,
		CustomerVehicleID:   wo.CustomerVehicleID,
		CashierUserID:       wo.CashierUserID,
		ShiftID:             wo.ShiftID,
		QueueNumber:         wo.QueueNumber,
		Status:              string(wo.Status),
		Notes:               wo.Notes,
		SpecialInstructions: wo.SpecialInstructions,
		ConfirmedAt:         wo.ConfirmedAt,
		StartedAt:           wo.StartedAt,
		CompletedAt:         wo.CompletedAt,
		Subtotal:            wo.Subtotal,
		DiscountAmount:      wo.DiscountAmount,
		TaxAmount:           wo.TaxAmount,
		TotalAmount:         wo.TotalAmount,
		CreatedAt:           wo.CreatedAt,
		UpdatedAt:           wo.UpdatedAt,
	}

	// Add items if loaded
	if len(wo.Items) > 0 {
		response.Items = make([]dto.WorkOrderItemResponse, len(wo.Items))
		for i, item := range wo.Items {
			response.Items[i] = dto.WorkOrderItemResponse{
				ID:                  item.ID,
				WorkOrderID:         item.WorkOrderID,
				ProductID:           item.ProductID,
				ProductNameSnapshot: item.ProductNameSnapshot,
				PriceSnapshot:       item.PriceSnapshot,
				Quantity:            item.Quantity,
				Subtotal:            item.Subtotal,
				AssignedStaffUserID: item.AssignedStaffUserID,
				ItemNote:            item.ItemNote,
				CreatedAt:           item.CreatedAt,
				UpdatedAt:           item.UpdatedAt,
			}
		}
	}

	return response
}
