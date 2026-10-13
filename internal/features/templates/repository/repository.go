package templates_repository

import repository_pool "golang-clean-server/internal/core/repository/pool"

type TemplatesRepository struct {
	pool repository_pool.Pool
}

func NewTemplatesRepository(pool repository_pool.Pool) *TemplatesRepository {
	return &TemplatesRepository{pool: pool}
}
