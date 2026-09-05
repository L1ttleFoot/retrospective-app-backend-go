package messages_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang-clean-server/internal/core/domain"

	"github.com/stretchr/testify/require"
)

type mockMessagesRepository struct {
	messages map[string]domain.Message
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

func (m *mockMessagesRepository) GetMessagesBySectionId(ctx context.Context, sectionId string) ([]domain.Message, error) {
	var result []domain.Message
	for _, msg := range m.messages {
		if msg.SectionID == sectionId {
			result = append(result, msg)
		}
	}
	return result, nil
}

func (m *mockMessagesRepository) GetMessageById(ctx context.Context, messageId string) (domain.Message, error) {
	msg, ok := m.messages[messageId]
	if !ok {
		return domain.Message{}, errors.New("message not found")
	}
	return msg, nil
}

func (m *mockMessagesRepository) PatchMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	m.messages[message.ID] = message
	return message, nil
}

func (m *mockMessagesRepository) DeleteMessage(ctx context.Context, messageId string) error {
	delete(m.messages, messageId)
	return nil
}

type mockMessageEventPublisher struct{}

func (m *mockMessageEventPublisher) PublishMessageCreated(ctx context.Context, msg domain.Message) error {
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
		"section-1",
	)

	service := NewMessagesService(&mockMessagesRepository{}, &mockMessageEventPublisher{})

	result, err := service.CreateMessage(context.Background(), msg)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "1", result.ID)
}

func TestMessagesService_CreateMessage_ValidationError(t *testing.T) {
	invalidMsg := domain.Message{
		Text: "",
	}

	service := NewMessagesService(&mockMessagesRepository{}, &mockMessageEventPublisher{})

	_, err := service.CreateMessage(context.Background(), invalidMsg)

	require.Error(t, err)
}

func TestMessagesService_GetMessagesBySectionId(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	mockRepo.messages["1"] = domain.NewMessage("1", "msg1", "a1", timeNow(), timeNow(), "s1")
	mockRepo.messages["2"] = domain.NewMessage("2", "msg2", "a1", timeNow(), timeNow(), "s1")
	mockRepo.messages["3"] = domain.NewMessage("3", "msg3", "a1", timeNow(), timeNow(), "s2")

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{})

	result, err := service.GetMessagesBySectionId(context.Background(), "s1")

	require.NoError(t, err)
	require.Equal(t, 2, len(result))
	require.Contains(t, []string{"msg1", "msg2"}, result[0].Text)
	require.Contains(t, []string{"msg1", "msg2"}, result[1].Text)
}

func TestMessagesService_PatchMessage(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	original := domain.NewMessage(
		"1",
		"original text",
		"author-1",
		timeNow(),
		timeNow(),
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
		"section-1",
	)
	mockRepo.messages["1"] = updated

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{})

	result, err := service.PatchMessage(context.Background(), "1", patch)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "updated text", result.Text)
}

func TestMessagesService_DeleteMessage(t *testing.T) {
	mockRepo := newMockMessagesRepository()
	mockRepo.messages["1"] = domain.NewMessage("1", "text", "a1", timeNow(), timeNow(), "s1")

	service := NewMessagesService(mockRepo, &mockMessageEventPublisher{})

	err := service.DeleteMessage(context.Background(), "1")

	require.NoError(t, err)
	require.Equal(t, 0, len(mockRepo.messages))
}
