package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/yourusername/go-skeleton/internal/entity"
)

type ExampleRepository interface {
	Create(ctx context.Context, example *entity.Example) error
	FindByID(ctx context.Context, id uint) (*entity.Example, error)
	FindAll(ctx context.Context, filter map[string]interface{}, page, size int) ([]*entity.Example, int64, error)
	Update(ctx context.Context, example *entity.Example) error
	Delete(ctx context.Context, id uint) error
}

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}

func (r *exampleRepository) Create(ctx context.Context, example *entity.Example) error {
	return r.db.WithContext(ctx).Create(example).Error
}

func (r *exampleRepository) FindByID(ctx context.Context, id uint) (*entity.Example, error) {
	var example entity.Example
	if err := r.db.WithContext(ctx).First(&example, id).Error; err != nil {
		return nil, err
	}
	return &example, nil
}

func (r *exampleRepository) FindAll(ctx context.Context, filter map[string]interface{}, page, size int) ([]*entity.Example, int64, error) {
	var examples []*entity.Example
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Example{})

	// Apply filters
	for key, value := range filter {
		if value != "" && value != nil {
			query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
		}
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Find(&examples).Error; err != nil {
		return nil, 0, err
	}

	return examples, total, nil
}

func (r *exampleRepository) Update(ctx context.Context, example *entity.Example) error {
	return r.db.WithContext(ctx).Save(example).Error
}

func (r *exampleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Example{}, id).Error
}
