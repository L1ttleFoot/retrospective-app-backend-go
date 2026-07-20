package user_repository

import (
	"context"
	"errors"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(ctx context.Context, userId string) (domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT username
		FROM users
		WHERE id=$1
	`

	row := r.pool.QueryRow(ctx, query, userId)

	var userModel UserModel

	if err := row.Scan(&userModel.Username); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id=%s : %w",
				userId,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error : %w", err)
	}

	userDomain := domain.NewUser(
		userId,
		userModel.Username,
		"",
	)

	return userDomain, nil
}
