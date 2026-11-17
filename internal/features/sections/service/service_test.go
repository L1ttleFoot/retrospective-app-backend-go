package sections_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockServiceRepository struct {
	sections map[string]domain.Section
}

func newMockBoardsRepository() *mockServiceRepository {
	return &mockServiceRepository{
		sections: make(map[string]domain.Section),
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

func (r *mockServiceRepository) GetSections(ctx context.Context, boardId string) ([]domain.Section, error) {

	var result []domain.Section

	for _, s := range r.sections {
		if s.BoardID == boardId {
			result = append(result, s)
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

func TestSectionsService_GetSections(t *testing.T) {
	mockRepo := newMockBoardsRepository()

	mockRepo.sections["1"] = domain.NewSection("1", "test-title-1", "red", "1", timeNow())
	mockRepo.sections["2"] = domain.NewSection("2", "test-title-2", "green", "1", timeNow())
	mockRepo.sections["3"] = domain.NewSection("1", "test-title-1", "blue", "2", timeNow())

	service := NewSectionsService(mockRepo)

	result, err := service.GetSections(context.Background(), "1")

	require.NoError(t, err)
	require.Equal(t, 2, len(result))
	require.Contains(t, []string{"test-title-1", "test-title-2"}, result[0].Title)
	require.Contains(t, []string{"test-title-1", "test-title-2"}, result[1].Title)
}
