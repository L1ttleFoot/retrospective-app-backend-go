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
	DeleteBoard(ctx context.Context, boardId string) error
}

func NewBoardsService(
	boardsRepository BoardsRepository,
) *BoardsService {
	return &BoardsService{
		boardsRepository: boardsRepository,
	}
}
