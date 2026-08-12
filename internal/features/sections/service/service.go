package sections_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
)

type SectionsService struct {
	sectionsRepository SectionsRepository
}

type SectionsRepository interface {
	CreateSections(ctx context.Context, sections []domain.Section) ([]domain.Section, error)
	GetSections(ctx context.Context, boardId string) ([]domain.Section, error)
}

func NewSectionsService(sectionsRepository SectionsRepository) *SectionsService {
	return &SectionsService{sectionsRepository: sectionsRepository}
}
