package user_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *UsersRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
	WITH new_user AS (
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id
	),
	_ AS (
		INSERT INTO user_roles (user_id, role_id)
		SELECT new_user.id, roles.id
		FROM new_user
		JOIN roles ON roles.role = 'USER'
	)
	SELECT id FROM new_user;
	`

	row := r.pool.QueryRow(ctx, query, user.Username, user.Password)

	var userModel UserModel

	if err := row.Scan(&userModel.ID); err != nil {
		return domain.User{}, fmt.Errorf("user scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		user.Username,
		"",
		"",
		"",
		[]string{"USER"},
	)

	return userDomain, nil
}
