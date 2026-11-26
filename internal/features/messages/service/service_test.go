package messages_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"

	"github.com/stretchr/testify/require"
)

type mockMessagesRepository struct {
	messages       map[string]domain.Message
	boardBySection map[string]string
}

func newMockMessagesRepository() *mockMessagesRepository {
	return &mockMessagesRepository{
		messages: make(map[string]domain.Message),
	}
}

func (m *mockMessagesRepository) CreateMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	if m.messages == nil {
		m.messages = make(map[string]domain.Message)
	}
	m.messages[message.ID] = message

	return message, nil
}

func (m *mockMessagesRepository) CreateMessageWithOutbox(ctx context.Context, message domain.Message) (domain.Message, error) {
	if m.messages == nil {
		m.messages = make(map[string]domain.Message)
	}
	m.messages[message.ID] = message

	return message, nil
}

func (m *mockMessagesRepository) GetMessageById(ctx context.Context, messageId string) (domain.Message, error) {
	msg, ok := m.messages[messageId]
	if !ok {
		return domain.Message{}, errors.New("message not found")
	}
	return msg, nil
}

func (m *mockMessagesRepository) GetBoardIdBySectionId(ctx context.Context, sectionId string) (string, error) {
	if board, ok := m.boardBySection[sectionId]; ok {
		return board, nil
	}
	return "board-1", nil
}

func (m *mockMessagesRepository) PatchMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	m.messages[message.ID] = message
	return message, nil
}

func (m *mockMessagesRepository) DeleteMessage(ctx context.Context, messageId string) error {
	delete(m.messages, messageId)
	return nil
}

func (m *mockMessagesRepository) PinMessage(ctx context.Context, messageId string) error {

	return nil
}

func (m *mockMessagesRepository) UnpinMessage(ctx context.Context, messageId string) error {

	return nil
}

func (m *mockMessagesRepository) SetReaction(ctx context.Context, reaction string, messageId string, authorId string) error {

	return nil
}

type mockMessageEventPublisher struct{}

func (m *mockMessageEventPublisher) PublishMessageCreated(ctx context.Context, queue string, msg domain.Message) error {
	return nil
}

type mockEventBroadcaster struct{}

func (m *mockEventBroadcaster) Publish(boardID, eventType string, payload any) error {
	return nil
}

func timeNow() time.Time {
	return time.Now()
}

func nullableString(s string) domain.Nullable[string] {
	return domain.Nullable[string]{Value: &s, Set: true}
}

func TestMessagesService_CreateMessage(t *testing.T) {
	msg := domain.NewMessage(
		"1",
		"test message",
		"author-1",
		timeNow(),
		timeNow(),
		nil,
		"section-1",
	)

	service := NewMessagesService(&mockMessagesRepository{}, &mockMessageEventPublisher{}, &mockEventBroadcaster{})

	result, err := service.CreateMessage(context.Background(), msg)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "1", result.ID)
}

func TestMessagesService_CreateMessage_ValidationError(t *testing.T) {
	invalidMsg := domain.Message{
		Text: "",
	}

	service := NewMessagesService(&mockMessagesRepository{}, &mockMessageEventPublisher{}, &mockEventBroadcaster{})

	_, err := service.CreateMessage(context.Background(), invalidMsg)

	require.Error(t, err)
}

func TestMessagesService_PatchMessage(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	original := domain.NewMessage(
		"1",
		"original text",
		"author-1",
		timeNow(),
		timeNow(),
		nil,
		"section-1",
	)
	mockRepo.messages["1"] = original

	patch := domain.MessagePatch{
		Text:      nullableString("updated text"),
		SectionID: domain.Nullable[string]{},
	}

	updated := domain.NewMessage(
		"1",
		"updated text",
		"author-1",
		timeNow(),
		timeNow(),
		nil,
		"section-1",
	)
	mockRepo.messages["1"] = updated

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{}, &mockEventBroadcaster{})

	result, err := service.PatchMessage(context.Background(), "1", patch)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "updated text", result.Text)
}

func TestMessagesService_PatchMessage_MoveAcrossBoardsRejected(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	mockRepo.boardBySection = map[string]string{"s1": "board-1", "s2": "board-2"}
	mockRepo.messages["1"] = domain.NewMessage("1", "text", "a1", timeNow(), timeNow(), nil, "s1")

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{}, &mockEventBroadcaster{})

	patch := domain.MessagePatch{SectionID: nullableString("s2")}

	_, err := service.PatchMessage(context.Background(), "1", patch)

	require.ErrorIs(t, err, core_errors.ErrInvalidArgument)
	require.Equal(t, "s1", mockRepo.messages["1"].SectionID)
}

func TestMessagesService_DeleteMessage(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	mockRepo.messages["1"] = domain.NewMessage("1", "text", "a1", timeNow(), timeNow(), nil, "s1")

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{}, &mockEventBroadcaster{})

	err := service.DeleteMessage(context.Background(), "1")

	require.NoError(t, err)
	require.Equal(t, 0, len(mockRepo.messages))
}
