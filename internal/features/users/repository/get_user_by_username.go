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
		SELECT id, password
		FROM users
		WHERE username=$1;
	`

	row := r.pool.QueryRow(ctx, query, username)

	var userModel UserModel

	if err := row.Scan(&userModel.ID, &userModel.Password); err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		username,
		userModel.Password,
		"",
		"",
	)

	return userDomain, nil
}
