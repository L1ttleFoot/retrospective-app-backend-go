package boards_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *BoardsService) GetSections(ctx context.Context, boardId string) (string, []domain.Section, error) {

	domainBoard, err := s.boardsRepository.GetBoardById(ctx, boardId)
	if err != nil {
		return "", []domain.Section{}, fmt.Errorf("get sections: %w", err)
	}

	domainSections, err := s.boardsRepository.GetSections(ctx, boardId)
	if err != nil {
		return "", []domain.Section{}, fmt.Errorf("get sections: %w", err)
	}

	return domainBoard.Title, domainSections, nil
}
