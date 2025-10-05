package repository

import (
	"context"
	"fmt"
	"time"

	"flashlight-go/internal/models"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	*BaseRepository[models.Payment]
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db),
	}
}

func (r *PaymentRepository) GeneratePaymentNumber(ctx context.Context) (string, error) {
	now := time.Now()
	prefix := fmt.Sprintf("PAY-%s", now.Format("20060102"))

	var count int64
	err := r.DB().WithContext(ctx).Model(&models.Payment{}).
		Where("payment_number LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}

	paymentNumber := fmt.Sprintf("%s-%04d", prefix, count+1)
	return paymentNumber, nil
}

func (r *PaymentRepository) FindByWorkOrder(ctx context.Context, workOrderID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.DB().WithContext(ctx).
		Where("work_order_id = ?", workOrderID).
		Order("created_at DESC").
		Find(&payments).Error
	return payments, err
}

func (r *PaymentRepository) GetTotalPaidForWorkOrder(ctx context.Context, workOrderID uint) (float64, error) {
	var total float64
	err := r.DB().WithContext(ctx).Model(&models.Payment{}).
		Where("work_order_id = ? AND status = ?", workOrderID, models.PaymentStatusCompleted).
		Select("COALESCE(SUM(amount_paid), 0)").
		Scan(&total).Error
	return total, err
}

func (r *PaymentRepository) FindByShift(ctx context.Context, shiftID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.DB().WithContext(ctx).
		Where("shift_id = ?", shiftID).
		Preload("WorkOrder").
		Find(&payments).Error
	return payments, err
}
