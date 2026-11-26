package reports_repository

import (
	"context"
	"errors"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *ReportsRepository) GetBoardInfo(ctx context.Context, boardId string) (domain.ReportBoardInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	row, err := r.q.GetBoardById(ctx, boardId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ReportBoardInfo{}, fmt.Errorf("board with id %s not found: %w", boardId, core_errors.ErrNotFound)
		}
		return domain.ReportBoardInfo{}, fmt.Errorf("get board info: %w", err)
	}

	return domain.ReportBoardInfo{
		ID:        row.ID,
		Title:     row.Title,
		OwnerID:   row.OwnerID,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (r *ReportsRepository) GetSectionsStats(ctx context.Context, boardId string) ([]domain.SectionStat, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := r.q.GetReportSections(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("get report sections: %w", err)
	}

	result := make([]domain.SectionStat, 0, len(rows))
	for _, row := range rows {
		result = append(result, domain.SectionStat{
			ID:            row.ID,
			Title:         row.Title,
			Color:         row.Color,
			MessagesCount: int(row.MessagesCount),
			PinnedCount:   int(row.PinnedCount),
		})
	}

	return result, nil
}

func (r *ReportsRepository) GetReactionsStats(ctx context.Context, boardId string) ([]domain.ReactionStat, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := r.q.GetReportReactions(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("get report reactions: %w", err)
	}

	result := make([]domain.ReactionStat, 0, len(rows))
	for _, row := range rows {
		result = append(result, domain.ReactionStat{
			Char:  row.Char,
			Name:  row.Name,
			Count: int(row.UsageCount),
		})
	}

	return result, nil
}

func (r *ReportsRepository) GetTopMessages(ctx context.Context, boardId string) ([]domain.TopMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	rows, err := r.q.GetReportTopMessages(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("get report top messages: %w", err)
	}

	result := make([]domain.TopMessage, 0, len(rows))
	for _, row := range rows {
		result = append(result, domain.TopMessage{
			MessageID:      row.MessageID,
			SectionTitle:   row.SectionTitle,
			Text:           row.Text,
			ReactionsCount: int(row.ReactionsCount),
		})
	}

	return result, nil
}
