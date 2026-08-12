package messages_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type MessagesService struct {
	messagesRepository MessagesRepository
}

type MessagesRepository interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	GetMessagesBySectionId(ctx context.Context, sectionId string) ([]domain.Message, error)
	DeleteMessage(ctx context.Context, messageId string) error
}

func NewMessagesService(messagesRepository MessagesRepository) *MessagesService {
	return &MessagesService{messagesRepository: messagesRepository}
}
