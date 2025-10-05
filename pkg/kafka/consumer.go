package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/go-skeleton/pkg/config"
)

type ConsumerHandler interface {
	Handle(ctx context.Context, message *sarama.ConsumerMessage) error
}

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	topics        []string
	handler       ConsumerHandler
	logger        *logrus.Logger
}

func NewConsumer(cfg *config.KafkaConfig, topics []string, handler ConsumerHandler, logger *logrus.Logger) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumerGroup, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.ConsumerGroup, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumerGroup: consumerGroup,
		topics:        topics,
		handler:       handler,
		logger:        logger,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	handler := &consumerGroupHandler{
		handler: c.handler,
		logger:  c.logger,
	}

	for {
		if err := c.consumerGroup.Consume(ctx, c.topics, handler); err != nil {
			c.logger.Errorf("Error from consumer: %v", err)
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *Consumer) Close() error {
	return c.consumerGroup.Close()
}

type consumerGroupHandler struct {
	handler ConsumerHandler
	logger  *logrus.Logger
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.logger.Infof("Message claimed: topic = %s, partition = %d, offset = %d", message.Topic, message.Partition, message.Offset)

		if err := h.handler.Handle(session.Context(), message); err != nil {
			h.logger.Errorf("Error handling message: %v", err)
			continue
		}

		session.MarkMessage(message, "")
	}
	return nil
}
