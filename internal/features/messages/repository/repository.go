package messages_repository

import (
	"golang-clean-server/db"
	repository_pool "golang-clean-server/internal/core/repository/pool"
)

type MessagesRepository struct {
	pool repository_pool.Pool
	q    *db.Queries
}

func NewMessagesRepository(pool repository_pool.Pool) *MessagesRepository {
	return &MessagesRepository{
		pool: pool,
		q:    db.New(pool),
	}
}
