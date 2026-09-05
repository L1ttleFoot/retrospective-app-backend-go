package boards_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockBoardsRepository struct {
	boards map[string]domain.Board
}

func newMockBoardsRepository() *mockBoardsRepository {
	return &mockBoardsRepository{
		boards: make(map[string]domain.Board),
	}
}

func (b *mockBoardsRepository) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {
	if b.boards == nil {
		b.boards = make(map[string]domain.Board)
	}
	b.boards[board.ID] = board
	return board, nil
}

func (b *mockBoardsRepository) GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error) {
	var result []domain.Board
	for _, b := range b.boards {
		if b.OwnerID == ownerId {
			result = append(result, b)
		}
	}
	return result, nil
}

func (b *mockBoardsRepository) DeleteBoard(ctx context.Context, boardId string) error {
	delete(b.boards, boardId)
	return nil
}

func timeNow() time.Time {
	return time.Now()
}

func TestBoardsService_CreateBoard(t *testing.T) {
	board := domain.NewBoard("1", "test board", "owner-1", timeNow())

	service := NewBoardsService(&mockBoardsRepository{})

	result, err := service.CreateBoard(context.Background(), board)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, "1", result.ID)
}

func TestBoardsService_CreateBoard_ValidationError(t *testing.T) {
	invalidBoard := domain.NewBoard("1", "", "owner-1", timeNow())

	service := NewBoardsService(&mockBoardsRepository{})

	_, err := service.CreateBoard(context.Background(), invalidBoard)

	require.Error(t, err)
}

func TestBoardsService_GetBoardsByOwnderId(t *testing.T) {
	mockRepo := newMockBoardsRepository()

	mockRepo.boards["1"] = domain.NewBoard("1", "test-title-1", "o1", timeNow())
	mockRepo.boards["2"] = domain.NewBoard("2", "test-title-2", "o1", timeNow())
	mockRepo.boards["3"] = domain.NewBoard("3", "test-title-3", "o2", timeNow())

	service := NewBoardsService(mockRepo)

	result, err := service.GetBoardsByOwnerId(context.Background(), "o1")

	require.NoError(t, err)
	require.Equal(t, 2, len(result))
	require.Contains(t, []string{"test-title-1", "test-title-2"}, result[0].Title)
	require.Contains(t, []string{"test-title-1", "test-title-2"}, result[1].Title)
}
