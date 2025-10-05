package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"

	"github.com/yourusername/go-skeleton/internal/model"
	"github.com/yourusername/go-skeleton/internal/usecase"
)

type ExampleConsumerHandler struct {
	useCase usecase.ExampleUseCase
	logger  *logrus.Logger
}

func NewExampleConsumerHandler(useCase usecase.ExampleUseCase, logger *logrus.Logger) *ExampleConsumerHandler {
	return &ExampleConsumerHandler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *ExampleConsumerHandler) Handle(ctx context.Context, message *sarama.ConsumerMessage) error {
	h.logger.Infof("Processing message from topic %s: %s", message.Topic, string(message.Value))

	var req model.ExampleRequest
	if err := json.Unmarshal(message.Value, &req); err != nil {
		h.logger.Errorf("Failed to unmarshal message: %v", err)
		return err
	}

	// Process the message using use case
	result, err := h.useCase.Create(ctx, &req)
	if err != nil {
		h.logger.Errorf("Failed to process message: %v", err)
		return err
	}

	h.logger.Infof("Successfully processed message: %+v", result)
	return nil
}
