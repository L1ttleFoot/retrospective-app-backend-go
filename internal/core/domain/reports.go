package domain

import "time"

type BoardReport struct {
	Board        ReportBoardInfo
	Summary      ReportSummary
	Sections     []SectionStat
	TopReactions []ReactionStat
	TopMessages  []TopMessage
}

type ReportBoardInfo struct {
	ID        string
	Title     string
	OwnerID   string
	CreatedAt time.Time
}

type ReportSummary struct {
	TotalMessages  int
	PinnedMessages int
	TotalReactions int
}

type SectionStat struct {
	ID            string
	Title         string
	Color         string
	MessagesCount int
	SharePercent  float64
	PinnedCount   int
}

type ReactionStat struct {
	Char  string
	Name  string
	Count int
}

type TopMessage struct {
	MessageID      string
	SectionTitle   string
	Text           string
	ReactionsCount int
}
