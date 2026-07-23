package boards_service

import (
	"context"
	"fmt"
)

func (s *BoardsService) DeleteBoard(ctx context.Context, boardId string) error {

	if err := s.boardsRepository.DeleteBoard(ctx, boardId); err != nil {
		return fmt.Errorf("delete board : %w", err)
	}

	return nil
}
