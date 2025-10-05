package repository

import (
	"context"
	"errors"

	"flashlight-go/internal/models"

	"gorm.io/gorm"
)

type ShiftRepository struct {
	*BaseRepository[models.Shift]
}

func NewShiftRepository(db *gorm.DB) *ShiftRepository {
	return &ShiftRepository{
		BaseRepository: NewBaseRepository[models.Shift](db),
	}
}

func (r *ShiftRepository) FindActiveShiftByUser(ctx context.Context, userID uint) (*models.Shift, error) {
	var shift models.Shift
	err := r.DB().WithContext(ctx).
		Where("user_id = ? AND status = ?", userID, models.ShiftStatusActive).
		First(&shift).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &shift, nil
}

func (r *ShiftRepository) FindWithDetails(ctx context.Context, id uint) (*models.Shift, error) {
	var shift models.Shift
	err := r.DB().WithContext(ctx).
		Preload("User").
		Preload("WorkOrders").
		Preload("Payments").
		First(&shift, id).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *ShiftRepository) GetShiftSummary(ctx context.Context, shiftID uint) (map[string]interface{}, error) {
	var totalSales float64
	var totalOrders int64

	err := r.DB().WithContext(ctx).Model(&models.Payment{}).
		Where("shift_id = ? AND status = ?", shiftID, models.PaymentStatusCompleted).
		Select("COALESCE(SUM(amount_paid), 0)").
		Scan(&totalSales).Error
	if err != nil {
		return nil, err
	}

	err = r.DB().WithContext(ctx).Model(&models.WorkOrder{}).
		Where("shift_id = ?", shiftID).
		Count(&totalOrders).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_sales":  totalSales,
		"total_orders": totalOrders,
	}, nil
}
