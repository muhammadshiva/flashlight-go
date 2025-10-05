package repository

import (
	"context"

	"flashlight-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByRole(ctx context.Context, role models.UserRole) ([]models.User, error) {
	var users []models.User
	err := r.DB().WithContext(ctx).Where("role = ? AND is_active = ?", role, true).Find(&users).Error
	return users, err
}

func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID uint) error {
	return r.DB().WithContext(ctx).Model(&models.User{}).
		Where("id = ?", userID).
		Update("last_login_at", gorm.Expr("NOW()")).Error
}
