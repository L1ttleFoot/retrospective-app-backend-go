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
	INSERT INTO users (username, password)
	VALUES ($1, $2)
	RETURNING id;
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
	)

	return userDomain, nil
}
