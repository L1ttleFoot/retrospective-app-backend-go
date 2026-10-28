package user_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *UsersRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		SELECT u.id, u.password,
			COALESCE(array_agg(r.role ORDER BY r.role) FILTER (WHERE r.role IS NOT NULL), '{}') AS roles
		FROM users u
		LEFT JOIN user_roles ur ON ur.user_id = u.id
		LEFT JOIN roles r ON r.id = ur.role_id
		WHERE u.username = $1
		GROUP BY u.id, u.password;
	`

	row := r.pool.QueryRow(ctx, query, username)

	var userModel UserModel
	var roles []string

	if err := row.Scan(&userModel.ID, &userModel.Password, &roles); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		username,
		userModel.Password,
		"",
		"",
		roles,
	)

	return userDomain, nil
}
