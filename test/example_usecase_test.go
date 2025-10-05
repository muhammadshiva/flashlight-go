package test

import (
	"context"
	"testing"

	"github.com/yourusername/go-skeleton/internal/entity"
	"github.com/yourusername/go-skeleton/internal/model"
)

// Mock repository
type mockExampleRepository struct {
	createFunc  func(ctx context.Context, example *entity.Example) error
	findByIDFunc func(ctx context.Context, id uint) (*entity.Example, error)
}

func (m *mockExampleRepository) Create(ctx context.Context, example *entity.Example) error {
	return m.createFunc(ctx, example)
}

func (m *mockExampleRepository) FindByID(ctx context.Context, id uint) (*entity.Example, error) {
	return m.findByIDFunc(ctx, id)
}

func (m *mockExampleRepository) FindAll(ctx context.Context, filter map[string]interface{}, page, size int) ([]*entity.Example, int64, error) {
	return nil, 0, nil
}

func (m *mockExampleRepository) Update(ctx context.Context, example *entity.Example) error {
	return nil
}

func (m *mockExampleRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

// Mock gateway
type mockExampleGateway struct {
	sendFunc func(ctx context.Context, data *model.ExampleRequest) (*model.ExampleResponse, error)
}

func (m *mockExampleGateway) SendToExternalAPI(ctx context.Context, data *model.ExampleRequest) (*model.ExampleResponse, error) {
	return m.sendFunc(ctx, data)
}

// Test example
func TestExampleUseCase_Create(t *testing.T) {
	// TODO: Implement unit tests
	t.Skip("Example test - implement your tests here")
}
