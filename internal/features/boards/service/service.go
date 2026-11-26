package boards_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type BoardsService struct {
	boardsRepository BoardsRepository
}

type BoardsRepository interface {
	CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error)
	GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error)
	GetBoardById(ctx context.Context, boardId string) (domain.Board, error)
	DeleteBoard(ctx context.Context, boardId string) error
	GetBoardOwnerId(ctx context.Context, boardId string) (string, error)
	GetSections(ctx context.Context, boardId string) ([]domain.Section, error)
}

func NewBoardsService(
	boardsRepository BoardsRepository,
) *BoardsService {
	return &BoardsService{
		boardsRepository: boardsRepository,
	}
}
