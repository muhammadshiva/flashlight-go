package repository

import (
	"context"
	"fmt"
	"time"

	"flashlight-go/internal/models"

	"gorm.io/gorm"
)

type WorkOrderRepository struct {
	*BaseRepository[models.WorkOrder]
}

func NewWorkOrderRepository(db *gorm.DB) *WorkOrderRepository {
	return &WorkOrderRepository{
		BaseRepository: NewBaseRepository[models.WorkOrder](db),
	}
}

func (r *WorkOrderRepository) GenerateOrderNumber(ctx context.Context) (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("WO-%s", now.Format("20060102"))

	var count int64
	err := r.DB().WithContext(ctx).Model(&models.WorkOrder{}).
		Where("order_number LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	orderNumber := fmt.Sprintf("%s-%04d", prefix, count+1)
	return orderNumber, nil
}

func (r *WorkOrderRepository) FindWithItems(ctx context.Context, id uint) (*models.WorkOrder, error) {
	var workOrder models.WorkOrder
	err := r.DB().WithContext(ctx).
		Preload("Items").
		Preload("Items.Product").
		Preload("CustomerUser").
		Preload("CustomerVehicle").
		Preload("CustomerVehicle.Vehicle").
		Preload("CashierUser").
		Preload("Payments").
		First(&workOrder, id).Error
	if err != nil {
		return nil, err
	}
	return &workOrder, nil
}

func (r *WorkOrderRepository) FindByStatus(ctx context.Context, status models.WorkOrderStatus) ([]models.WorkOrder, error) {
	var orders []models.WorkOrder
	err := r.DB().WithContext(ctx).
		Where("status = ?", status).
		Preload("CustomerUser").
		Preload("CustomerVehicle").
		Preload("CustomerVehicle.Vehicle").
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}

func (r *WorkOrderRepository) FindByShift(ctx context.Context, shiftID uint) ([]models.WorkOrder, error) {
	var orders []models.WorkOrder
	err := r.DB().WithContext(ctx).
		Where("shift_id = ?", shiftID).
		Preload("Items").
		Preload("Payments").
		Find(&orders).Error
	return orders, err
}

func (r *WorkOrderRepository) GetNextQueueNumber(ctx context.Context) (int, error) {
	var maxQueue int
	today := time.Now().Format("2006-01-02")

	err := r.DB().WithContext(ctx).Model(&models.WorkOrder{}).
		Where("DATE(created_at) = ?", today).
		Select("COALESCE(MAX(queue_number), 0)").
		Scan(&maxQueue).Error

	return maxQueue + 1, err
}
