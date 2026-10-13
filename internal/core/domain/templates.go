package domain

import (
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	"time"
)

type Template struct {
	ID        string
	Title     string
	AuthorID  string
	CreatedAt time.Time
	Sections  []TemplateSection
}

type TemplateSection struct {
	ID    string
	Title string
	Color string
}

func NewTemplate(
	id string,
	title string,
	authorID string,
	createdAt time.Time,
) *Template {
	return &Template{
		ID:        id,
		Title:     title,
		AuthorID:  authorID,
		CreatedAt: createdAt,
	}
}

func (t *Template) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("%w: template title cannot be empty", core_errors.ErrInvalidArgument)
	}

	if len(t.Sections) == 0 || t.Sections == nil {
		return fmt.Errorf("%w: template sections cannot be empty", core_errors.ErrInvalidArgument)
	}

	return nil
}
