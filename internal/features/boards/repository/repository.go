package boards_repository

import repository_pool "golang-clean-server/internal/core/repository/pool"

type BoardsRepository struct {
	pool repository_pool.Pool
}

func NewBoardsRepository(pool repository_pool.Pool) *BoardsRepository {
	return &BoardsRepository{pool: pool}
}
