package usecase

import (
	"context"
	"math"

	"github.com/yourusername/go-skeleton/internal/entity"
	"github.com/yourusername/go-skeleton/internal/gateway"
	"github.com/yourusername/go-skeleton/internal/model"
	"github.com/yourusername/go-skeleton/internal/repository"
)

type ExampleUseCase interface {
	Create(ctx context.Context, req *model.ExampleRequest) (*model.ExampleResponse, error)
	GetByID(ctx context.Context, id uint) (*model.ExampleResponse, error)
	GetAll(ctx context.Context, filter *model.ExampleFilterRequest) ([]*model.ExampleResponse, *model.MetaData, error)
	Update(ctx context.Context, id uint, req *model.ExampleRequest) (*model.ExampleResponse, error)
	Delete(ctx context.Context, id uint) error
}

type exampleUseCase struct {
	repo    repository.ExampleRepository
	gateway gateway.ExampleGateway
}

func NewExampleUseCase(repo repository.ExampleRepository, gateway gateway.ExampleGateway) ExampleUseCase {
	return &exampleUseCase{
		repo:    repo,
		gateway: gateway,
	}
}

func (u *exampleUseCase) Create(ctx context.Context, req *model.ExampleRequest) (*model.ExampleResponse, error) {
	// Convert model to entity
	example := &entity.Example{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}

	// Set default status if not provided
	if example.Status == "" {
		example.Status = "active"
	}

	// Save to database via repository
	if err := u.repo.Create(ctx, example); err != nil {
		return nil, err
	}

	// Optional: Call external gateway if needed
	// _, err := u.gateway.SendToExternalAPI(ctx, req)
	// if err != nil {
	//     return nil, err
	// }

	// Convert entity to response model
	return u.entityToResponse(example), nil
}

func (u *exampleUseCase) GetByID(ctx context.Context, id uint) (*model.ExampleResponse, error) {
	example, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.entityToResponse(example), nil
}

func (u *exampleUseCase) GetAll(ctx context.Context, filter *model.ExampleFilterRequest) ([]*model.ExampleResponse, *model.MetaData, error) {
	// Set default pagination
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Size == 0 {
		filter.Size = 10
	}

	// Build filter map
	filterMap := make(map[string]interface{})
	if filter.Name != "" {
		filterMap["name"] = filter.Name
	}
	if filter.Status != "" {
		filterMap["status"] = filter.Status
	}

	// Get data from repository
	examples, total, err := u.repo.FindAll(ctx, filterMap, filter.Page, filter.Size)
	if err != nil {
		return nil, nil, err
	}

	// Convert entities to response models
	var responses []*model.ExampleResponse
	for _, example := range examples {
		responses = append(responses, u.entityToResponse(example))
	}

	// Build metadata
	totalPage := int(math.Ceil(float64(total) / float64(filter.Size)))
	meta := &model.MetaData{
		Page:      filter.Page,
		PageSize:  filter.Size,
		Total:     total,
		TotalPage: totalPage,
	}

	return responses, meta, nil
}

func (u *exampleUseCase) Update(ctx context.Context, id uint, req *model.ExampleRequest) (*model.ExampleResponse, error) {
	// Get existing data
	example, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update entity fields
	example.Name = req.Name
	example.Description = req.Description
	if req.Status != "" {
		example.Status = req.Status
	}

	// Update in database
	if err := u.repo.Update(ctx, example); err != nil {
		return nil, err
	}

	return u.entityToResponse(example), nil
}

func (u *exampleUseCase) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

// Helper function to convert entity to response model
func (u *exampleUseCase) entityToResponse(example *entity.Example) *model.ExampleResponse {
	return &model.ExampleResponse{
		ID:          example.ID,
		Name:        example.Name,
		Description: example.Description,
		Status:      example.Status,
		CreatedAt:   example.CreatedAt,
		UpdatedAt:   example.UpdatedAt,
	}
}
