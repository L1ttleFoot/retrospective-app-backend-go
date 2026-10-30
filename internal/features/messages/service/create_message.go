package messages_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_logger "golang-clean-server/internal/core/logger"
	core_rabbitmq "golang-clean-server/internal/core/rabbitmq"

	"go.uber.org/zap"
)

func (s *MessagesService) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {

	if err := message.ValidateMessage(); err != nil {
		return domain.Message{}, fmt.Errorf("validate message: %w", err)
	}

	message, err := s.messagesRepository.CreateMessage(ctx, message)
	if err != nil {
		return domain.Message{}, fmt.Errorf("create message: %w", err)
	}

	if err := s.rmqPublisher.PublishMessageCreated(ctx, core_rabbitmq.QueueMessageCreated, message); err != nil {
		core_logger.FromContext(ctx).Error("publish messages.created", zap.Error(err))
	}

	s.publishMessageEvent(ctx, "message.created", messageEvent{SectionID: message.SectionID})

	return message, nil
}
