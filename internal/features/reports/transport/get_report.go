package reports_transport

import (
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_request "golang-clean-server/internal/core/transport/request"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"time"
)

type GetReportResponse struct {
	Board        ReportBoardInfoResponse `json:"board"`
	Summary      ReportSummaryResponse   `json:"summary"`
	Sections     []SectionStatResponse   `json:"sections"`
	TopReactions []ReactionStatResponse  `json:"topReactions"`
	TopMessages  []TopMessageResponse    `json:"topMessages"`
}

type ReportBoardInfoResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ReportSummaryResponse struct {
	TotalMessages  int `json:"totalMessages"`
	PinnedMessages int `json:"pinnedMessages"`
	TotalReactions int `json:"totalReactions"`
}

type SectionStatResponse struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Color         string  `json:"color"`
	MessagesCount int     `json:"messagesCount"`
	SharePercent  float64 `json:"sharePercent"`
	PinnedCount   int     `json:"pinnedCount"`
}

type ReactionStatResponse struct {
	Char  string `json:"char"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type TopMessageResponse struct {
	MessageID      string `json:"messageId"`
	SectionTitle   string `json:"sectionTitle"`
	Text           string `json:"text"`
	ReactionsCount int    `json:"reactionsCount"`
}

func (h *ReportsHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_response.NewResponseHandler(log, w)

	boardId, err := core_request.GetPathValue(r, "boardId")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid boardId format")
		return
	}

	claims := core_service_tokens.FromContext(ctx)

	report, err := h.reportsService.GetReport(ctx, boardId, claims.UserID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get report")
		return
	}

	sections := make([]SectionStatResponse, len(report.Sections))
	for i, s := range report.Sections {
		sections[i] = SectionStatResponse{
			ID:            s.ID,
			Title:         s.Title,
			Color:         s.Color,
			MessagesCount: s.MessagesCount,
			SharePercent:  s.SharePercent,
			PinnedCount:   s.PinnedCount,
		}
	}

	reactions := make([]ReactionStatResponse, len(report.TopReactions))
	for i, rc := range report.TopReactions {
		reactions[i] = ReactionStatResponse{
			Char:  rc.Char,
			Name:  rc.Name,
			Count: rc.Count,
		}
	}

	topMessages := make([]TopMessageResponse, len(report.TopMessages))
	for i, m := range report.TopMessages {
		topMessages[i] = TopMessageResponse{
			MessageID:      m.MessageID,
			SectionTitle:   m.SectionTitle,
			Text:           m.Text,
			ReactionsCount: m.ReactionsCount,
		}
	}

	response := GetReportResponse{
		Board: ReportBoardInfoResponse{
			ID:        report.Board.ID,
			Title:     report.Board.Title,
			OwnerID:   report.Board.OwnerID,
			CreatedAt: report.Board.CreatedAt,
		},
		Summary: ReportSummaryResponse{
			TotalMessages:  report.Summary.TotalMessages,
			PinnedMessages: report.Summary.PinnedMessages,
			TotalReactions: report.Summary.TotalReactions,
		},
		Sections:     sections,
		TopReactions: reactions,
		TopMessages:  topMessages,
	}

	responseHandler.JsonResponse(response, http.StatusOK)
}
