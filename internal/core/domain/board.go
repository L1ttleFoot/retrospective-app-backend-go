package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

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

func (b *Board) ValidateBoard() error {

	if b.Title == "" {
		return fmt.Errorf("%w: board title cannot be empty", core_errors.ErrInvalideArgument)
	}

	return nil
}
