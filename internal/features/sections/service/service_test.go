package sections_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockServiceRepository struct {
	messages map[string]domain.Message
	sections map[string]domain.Section
}

func newMockBoardsRepository() *mockServiceRepository {
	return &mockServiceRepository{
		sections: make(map[string]domain.Section),
		messages: make(map[string]domain.Message),
	}
}

func (r *mockServiceRepository) CreateSections(ctx context.Context, sections []domain.Section) ([]domain.Section, error) {
	if r.sections == nil {
		r.sections = make(map[string]domain.Section)
	}

	for _, s := range sections {
		r.sections[s.ID] = s
	}

	return sections, nil
}

func (m *mockServiceRepository) GetMessagesBySectionId(ctx context.Context, sectionId string, userId *string) ([]domain.Message, error) {
	var result []domain.Message
	for _, msg := range m.messages {
		if msg.SectionID == sectionId {
			result = append(result, msg)
		}
	}
	return result, nil
}

func timeNow() time.Time {
	return time.Now()
}

func TestSectionsService_CreateSections(t *testing.T) {

	mockRepo := newMockBoardsRepository()

	sections := []domain.Section{
		domain.NewSection("1", "test-title-1", "red", "1", timeNow()),
		domain.NewSection("2", "test-title-2", "green", "1", timeNow()),
		domain.NewSection("3", "test-title-3", "blue", "1", timeNow()),
	}

	service := NewSectionsService(mockRepo)

	result, err := service.CreateSections(context.Background(), sections)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result, 3)
	require.Len(t, mockRepo.sections, 3)
}

func TestSectionsService_GetMessagesBySectionId(t *testing.T) {
	mockRepo := newMockBoardsRepository()
	mockRepo.messages["1"] = domain.NewMessage("1", "msg1", "a1", timeNow(), timeNow(), nil, "s1")
	mockRepo.messages["2"] = domain.NewMessage("2", "msg2", "a1", timeNow(), timeNow(), nil, "s1")
	mockRepo.messages["3"] = domain.NewMessage("3", "msg3", "a1", timeNow(), timeNow(), nil, "s2")

	service := NewSectionsService(mockRepo)

	result, err := service.GetMessagesBySectionId(context.Background(), "s1", nil)

	require.NoError(t, err)
	require.Equal(t, 2, len(result))
	require.Contains(t, []string{"msg1", "msg2"}, result[0].Text)
	require.Contains(t, []string{"msg1", "msg2"}, result[1].Text)
}
