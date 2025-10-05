package service

import (
	"context"
	"errors"
	"time"

	"flashlight-go/internal/dto"
	"flashlight-go/internal/models"
	"flashlight-go/internal/repository"
	"flashlight-go/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:             req.Name,
		Email:            req.Email,
		PhoneNumber:      req.PhoneNumber,
		Password:         string(hashedPassword),
		Role:             models.UserRole(req.Role),
		MembershipTypeID: req.MembershipTypeID,
		Address:          req.Address,
		City:             req.City,
		State:            req.State,
		PostalCode:       req.PostalCode,
		Country:          req.Country,
		IsActive:         true,
	}

	// Parse membership expiration if provided
	if req.MembershipExpiresAt != nil {
		t, err := time.Parse(time.RFC3339, *req.MembershipExpiresAt)
		if err == nil {
			user.MembershipExpiresAt = &t
		}
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return s.toResponse(user), nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(user), nil
}

func (s *UserService) GetAll(ctx context.Context, page, perPage int) ([]dto.UserResponse, *dto.PaginationMeta, error) {
	users, total, err := s.userRepo.FindAll(ctx, page, perPage)
	if err != nil {
		return nil, nil, err
	}

	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *s.toResponse(&user)
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

func (s *UserService) Update(ctx context.Context, id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		// Check if new email already exists
		existing, _ := s.userRepo.FindByEmail(ctx, *req.Email)
		if existing != nil && existing.ID != id {
			return nil, errors.New("email already exists")
		}
		user.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		user.PhoneNumber = *req.PhoneNumber
	}
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	if req.Role != nil {
		user.Role = models.UserRole(*req.Role)
	}
	if req.MembershipTypeID != nil {
		user.MembershipTypeID = req.MembershipTypeID
	}
	if req.Address != nil {
		user.Address = req.Address
	}
	if req.City != nil {
		user.City = req.City
	}
	if req.State != nil {
		user.State = req.State
	}
	if req.PostalCode != nil {
		user.PostalCode = req.PostalCode
	}
	if req.Country != nil {
		user.Country = req.Country
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return s.toResponse(user), nil
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	_ = s.userRepo.UpdateLastLogin(ctx, user.ID)

	// Generate token
	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User:  *s.toResponse(user),
	}, nil
}

func (s *UserService) toResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:                  user.ID,
		Name:                user.Name,
		Email:               user.Email,
		PhoneNumber:         user.PhoneNumber,
		Role:                string(user.Role),
		MembershipTypeID:    user.MembershipTypeID,
		MembershipExpiresAt: user.MembershipExpiresAt,
		Address:             user.Address,
		City:                user.City,
		State:               user.State,
		PostalCode:          user.PostalCode,
		Country:             user.Country,
		ProfileImage:        user.ProfileImage,
		IsActive:            user.IsActive,
		LastLoginAt:         user.LastLoginAt,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}
}
