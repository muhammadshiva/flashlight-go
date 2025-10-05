package grpc

import (
	"context"

	"github.com/yourusername/go-skeleton/internal/model"
	"github.com/yourusername/go-skeleton/internal/usecase"
)

// ExampleService implements gRPC service
// Note: You need to generate protobuf files using protoc
type ExampleService struct {
	useCase usecase.ExampleUseCase
}

func NewExampleService(useCase usecase.ExampleUseCase) *ExampleService {
	return &ExampleService{
		useCase: useCase,
	}
}

// Example method - implement your gRPC methods here
func (s *ExampleService) CreateExample(ctx context.Context, req *model.ExampleRequest) (*model.ExampleResponse, error) {
	return s.useCase.Create(ctx, req)
}

// Add more gRPC methods as needed
