package boards_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *BoardsService) CreateBoard(ctx context.Context, board domain.Board) (domain.Board, error) {

	createdBoard, err := s.boardsRepository.CreateBoard(ctx, board)
	if err != nil {
		return domain.Board{}, fmt.Errorf("create board: %w", err)
	}

	return createdBoard, nil
}
