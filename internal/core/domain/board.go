package domain

import "time"

type Board struct {
	ID        string
	Title     string
	OwnerID   string
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
		OwnerID:   ownerId,
		CreatedAt: createdAt,
	}
}
