package service

import (
	"context"
	"errors"
	"time"

	"flashlight-go/internal/models"
	"flashlight-go/internal/repository"
)

type ShiftService struct {
	shiftRepo   *repository.ShiftRepository
	paymentRepo *repository.PaymentRepository
}

func NewShiftService(
	shiftRepo *repository.ShiftRepository,
	paymentRepo *repository.PaymentRepository,
) *ShiftService {
	return &ShiftService{
		shiftRepo:   shiftRepo,
		paymentRepo: paymentRepo,
	}
}

func (s *ShiftService) Start(ctx context.Context, userID uint, initialCash float64, receivedFrom *string) (*models.Shift, error) {
	// Check if user already has an active shift
	activeShift, err := s.shiftRepo.FindActiveShiftByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if activeShift != nil {
		return nil, errors.New("user already has an active shift")
	}

	shift := &models.Shift{
		UserID:       userID,
		StartTime:    time.Now(),
		InitialCash:  initialCash,
		Status:       models.ShiftStatusActive,
		ReceivedFrom: receivedFrom,
	}

	if err := s.shiftRepo.Create(ctx, shift); err != nil {
		return nil, err
	}

	return shift, nil
}

func (s *ShiftService) Close(ctx context.Context, shiftID uint, finalCash float64) (*models.Shift, error) {
	shift, err := s.shiftRepo.FindByID(ctx, shiftID)
	if err != nil {
		return nil, err
	}

	if shift.Status != models.ShiftStatusActive {
		return nil, errors.New("shift is not active")
	}

	// Calculate total sales
	summary, err := s.shiftRepo.GetShiftSummary(ctx, shiftID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	shift.EndTime = &now
	shift.FinalCash = finalCash
	shift.TotalSales = summary["total_sales"].(float64)
	shift.Status = models.ShiftStatusClosed

	if err := s.shiftRepo.Update(ctx, shift); err != nil {
		return nil, err
	}

	return shift, nil
}

func (s *ShiftService) GetByID(ctx context.Context, id uint) (*models.Shift, error) {
	return s.shiftRepo.FindWithDetails(ctx, id)
}

func (s *ShiftService) GetActiveByUser(ctx context.Context, userID uint) (*models.Shift, error) {
	return s.shiftRepo.FindActiveShiftByUser(ctx, userID)
}

func (s *ShiftService) GetSummary(ctx context.Context, shiftID uint) (map[string]interface{}, error) {
	shift, err := s.shiftRepo.FindByID(ctx, shiftID)
	if err != nil {
		return nil, err
	}

	summary, err := s.shiftRepo.GetShiftSummary(ctx, shiftID)
	if err != nil {
		return nil, err
	}

	summary["shift"] = shift
	return summary, nil
}
