package sections_repository

import repository_pool "golang-clean-server/internal/core/repository/pool"

type SectionsRepository struct {
	pool repository_pool.Pool
}

func NewSectionsRepository(pool repository_pool.Pool) *SectionsRepository {
	return &SectionsRepository{pool: pool}
}
