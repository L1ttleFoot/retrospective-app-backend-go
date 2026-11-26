package reports_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type ReportsService struct {
	reportsRepository ReportsRepository
}

type ReportsRepository interface {
	GetBoardInfo(ctx context.Context, boardId string) (domain.ReportBoardInfo, error)
	GetSectionsStats(ctx context.Context, boardId string) ([]domain.SectionStat, error)
	GetReactionsStats(ctx context.Context, boardId string) ([]domain.ReactionStat, error)
	GetTopMessages(ctx context.Context, boardId string) ([]domain.TopMessage, error)
}

func NewReportsService(
	reportsRepository ReportsRepository,
) *ReportsService {
	return &ReportsService{
		reportsRepository: reportsRepository,
	}
}
