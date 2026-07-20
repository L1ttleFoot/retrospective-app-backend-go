package user_repository

import repository_pool "golang-clean-server/internal/core/repository/pool"

type UsersRepository struct {
	pool repository_pool.Pool
}

func NewUsersRepository(pool repository_pool.Pool) *UsersRepository {
	return &UsersRepository{pool: pool}
}
