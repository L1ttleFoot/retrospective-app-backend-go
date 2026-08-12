package messages_repository

import repository_pool "golang-clean-server/internal/core/repository/pool"

type MessagesRepository struct {
	pool repository_pool.Pool
}

func NewMessagesRepository(pool repository_pool.Pool) *MessagesRepository {
	return &MessagesRepository{pool: pool}
}
