package boards_repository

import (
	repository_cache "golang-clean-server/internal/core/repository/cache"
	repository_pool "golang-clean-server/internal/core/repository/pool"
)

type BoardsRepository struct {
	pool  repository_pool.Pool
	cache repository_cache.Cache
}

func NewBoardsRepository(pool repository_pool.Pool, cache repository_cache.Cache) *BoardsRepository {
	return &BoardsRepository{pool: pool, cache: cache}
}
