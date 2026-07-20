package user_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT id, username
		FROM users
		LIMIT $1
		OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)

	if err != nil {
		return []domain.User{}, fmt.Errorf("pool error : %w", err)
	}
	defer rows.Close()

	var userModels []UserModel

	for rows.Next() {
		var u UserModel

		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, fmt.Errorf("scan user row: %w", err)
		}

		userModels = append(userModels, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows : %w", err)
	}

	userDomains := userDomainsFromModels(userModels)

	return userDomains, nil
}
