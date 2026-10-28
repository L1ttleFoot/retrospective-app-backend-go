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
		SELECT u.username,
			COALESCE(array_agg(r.role ORDER BY r.role) FILTER (WHERE r.role IS NOT NULL), '{}') AS roles
		FROM users u
		LEFT JOIN user_roles ur ON ur.user_id = u.id
		LEFT JOIN roles r ON r.id = ur.role_id
		WHERE u.id = $1
		GROUP BY u.id, u.username;
	`

	row := r.pool.QueryRow(ctx, query, userId)

	var userModel UserModel
	var roles []string

	if err := row.Scan(&userModel.Username, &roles); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id=%s: %w",
				userId,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userId,
		userModel.Username,
		"",
		"",
		"",
		roles,
	)

	return userDomain, nil
}
