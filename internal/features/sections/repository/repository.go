package sections_repository

import (
	db "golang-clean-server/db"
	repository_pool "golang-clean-server/internal/core/repository/pool"
)

type SectionsRepository struct {
	pool repository_pool.Pool
	q    *db.Queries
}

func NewSectionsRepository(pool repository_pool.Pool) *SectionsRepository {
	return &SectionsRepository{
		pool: pool,
		q:    db.New(pool),
	}
}
