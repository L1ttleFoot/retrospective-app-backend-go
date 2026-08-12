package domain

import "time"

type Section struct {
	ID        string
	Title     string
	Color     string
	BoardID   string
	CreatedAt time.Time
}

func NewSection(
	id string,
	title string,
	color string,
	boardId string,
	createdAt time.Time,
) Section {
	return Section{
		ID:        id,
		Title:     title,
		Color:     color,
		BoardID:   boardId,
		CreatedAt: createdAt,
	}
}
