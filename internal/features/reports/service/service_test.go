package reports_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeReportsRepo struct {
	board       domain.ReportBoardInfo
	boardErr    error
	sections    []domain.SectionStat
	reactions   []domain.ReactionStat
	topMessages []domain.TopMessage
}

func (f *fakeReportsRepo) GetBoardInfo(ctx context.Context, boardId string) (domain.ReportBoardInfo, error) {
	return f.board, f.boardErr
}
func (f *fakeReportsRepo) GetSectionsStats(ctx context.Context, boardId string) ([]domain.SectionStat, error) {
	return f.sections, nil
}
func (f *fakeReportsRepo) GetReactionsStats(ctx context.Context, boardId string) ([]domain.ReactionStat, error) {
	return f.reactions, nil
}
func (f *fakeReportsRepo) GetTopMessages(ctx context.Context, boardId string) ([]domain.TopMessage, error) {
	return f.topMessages, nil
}

func TestGetReport_ForbiddenForNonOwner(t *testing.T) {
	svc := NewReportsService(&fakeReportsRepo{
		board: domain.ReportBoardInfo{ID: "b1", OwnerID: "owner-1"},
	})

	_, err := svc.GetReport(context.Background(), "b1", "other-user")
	require.ErrorIs(t, err, core_errors.ErrForbidden)
}

func TestGetReport_AggregatesAndEmptyBoard(t *testing.T) {
	svc := NewReportsService(&fakeReportsRepo{
		board: domain.ReportBoardInfo{ID: "b1", OwnerID: "owner-1"},
		sections: []domain.SectionStat{
			{ID: "s1", MessagesCount: 3, PinnedCount: 1},
			{ID: "s2", MessagesCount: 1, PinnedCount: 1},
		},
		reactions: []domain.ReactionStat{
			{Name: "thumbsUp", Count: 5},
			{Name: "fire", Count: 2},
		},
		topMessages: []domain.TopMessage{},
	})

	rep, err := svc.GetReport(context.Background(), "b1", "owner-1")
	require.NoError(t, err)
	require.Equal(t, 4, rep.Summary.TotalMessages)
	require.Equal(t, 2, rep.Summary.PinnedMessages)
	require.Equal(t, 7, rep.Summary.TotalReactions)
	require.Equal(t, 75.0, rep.Sections[0].SharePercent)
	require.Equal(t, 25.0, rep.Sections[1].SharePercent)

	emptySvc := NewReportsService(&fakeReportsRepo{
		board:       domain.ReportBoardInfo{ID: "b1", OwnerID: "owner-1"},
		sections:    []domain.SectionStat{{ID: "s1", MessagesCount: 0}},
		reactions:   []domain.ReactionStat{},
		topMessages: []domain.TopMessage{},
	})
	emptyRep, err := emptySvc.GetReport(context.Background(), "b1", "owner-1")
	require.NoError(t, err)
	require.Equal(t, 0.0, emptyRep.Sections[0].SharePercent)
}
