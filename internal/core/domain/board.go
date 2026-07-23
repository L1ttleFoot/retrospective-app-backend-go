package domain

import "time"

type Board struct {
	ID        string
	Title     string
	OwnerId   string
	CreatedAt time.Time
}

func NewBoard(
	id string,
	title string,
	ownerId string,
	createdAt time.Time,
) Board {
	return Board{
		ID:        id,
		Title:     title,
		OwnerId:   ownerId,
		CreatedAt: createdAt,
	}
}
