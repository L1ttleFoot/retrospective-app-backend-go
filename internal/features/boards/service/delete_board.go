package boards_service

import (
	"context"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
)

func (s *BoardsService) DeleteBoard(ctx context.Context, boardId string, userId string) error {

	ownerId, err := s.boardsRepository.GetBoardOwnerId(ctx, boardId)
	if err != nil {
		return fmt.Errorf("get board owner: %w", err)
	}

	if ownerId != userId {
		return fmt.Errorf("delete board: %w", core_errors.ErrForbidden)
	}

	if err := s.boardsRepository.DeleteBoard(ctx, boardId); err != nil {
		return fmt.Errorf("delete board: %w", err)
	}

	return nil
}
