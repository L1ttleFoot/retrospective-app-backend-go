package reports_repository

import (
	"golang-clean-server/db"
	repository_pool "golang-clean-server/internal/core/repository/pool"
)

type ReportsRepository struct {
	pool repository_pool.Pool
	q    *db.Queries
}

func NewReportsRepository(pool repository_pool.Pool) *ReportsRepository {
	return &ReportsRepository{
		pool: pool,
		q:    db.New(pool),
	}
}
