package messages_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type MessagesService struct {
	messagesRepository MessagesRepository
	rmqPublisher       MessageEventPublisher
}

type MessageEventPublisher interface {
	PublishMessageCreated(ctx context.Context, queue string, msg domain.Message) error
}

type MessagesRepository interface {
	CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	GetMessagesBySectionId(ctx context.Context, sectionId string) ([]domain.Message, error)
	GetMessageById(ctx context.Context, messageId string) (domain.Message, error)
	PatchMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	DeleteMessage(ctx context.Context, messageId string) error
	PinMessage(ctx context.Context, messageId string) error
	UnpinMessage(ctx context.Context, messageId string) error
	SetReaction(ctx context.Context, reaction, messageId, authorId string) error
}

func NewMessagesService(
	messagesRepository MessagesRepository,
	rmqPublisher MessageEventPublisher,
) *MessagesService {
	return &MessagesService{
		messagesRepository: messagesRepository,
		rmqPublisher:       rmqPublisher,
	}
}
