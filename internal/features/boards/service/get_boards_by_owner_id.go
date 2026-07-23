package boards_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *BoardsService) GetBoardsByOwnerId(ctx context.Context, ownerId string) ([]domain.Board, error) {

	boards, err := s.boardsRepository.GetBoardsByOwnerId(ctx, ownerId)
	if err != nil {
		return []domain.Board{}, fmt.Errorf("get boards : %w", err)
	}

	return boards, nil
}
