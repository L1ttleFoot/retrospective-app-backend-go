package reports_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"math"
)

func (s *ReportsService) GetReport(ctx context.Context, boardId string, ownerId string) (domain.BoardReport, error) {
	board, err := s.reportsRepository.GetBoardInfo(ctx, boardId)
	if err != nil {
		return domain.BoardReport{}, fmt.Errorf("get board info: %w", err)
	}

	if board.OwnerID != ownerId {
		return domain.BoardReport{}, fmt.Errorf("get report: %w", core_errors.ErrForbidden)
	}

	sections, err := s.reportsRepository.GetSectionsStats(ctx, boardId)
	if err != nil {
		return domain.BoardReport{}, fmt.Errorf("get sections stats: %w", err)
	}

	reactions, err := s.reportsRepository.GetReactionsStats(ctx, boardId)
	if err != nil {
		return domain.BoardReport{}, fmt.Errorf("get reactions stats: %w", err)
	}

	topMessages, err := s.reportsRepository.GetTopMessages(ctx, boardId)
	if err != nil {
		return domain.BoardReport{}, fmt.Errorf("get top messages: %w", err)
	}

	var totalMessages, pinnedMessages, totalReactions int
	for _, sec := range sections {
		totalMessages += sec.MessagesCount
		pinnedMessages += sec.PinnedCount
	}
	if totalMessages > 0 {
		for i := range sections {
			pct := (float64(sections[i].MessagesCount) / float64(totalMessages)) * 100
			sections[i].SharePercent = math.Round(pct*10) / 10
		}
	}
	for _, r := range reactions {
		totalReactions += r.Count
	}

	return domain.BoardReport{
		Board: board,
		Summary: domain.ReportSummary{
			TotalMessages:  totalMessages,
			PinnedMessages: pinnedMessages,
			TotalReactions: totalReactions,
		},
		Sections:     sections,
		TopReactions: reactions,
		TopMessages:  topMessages,
	}, nil
}
