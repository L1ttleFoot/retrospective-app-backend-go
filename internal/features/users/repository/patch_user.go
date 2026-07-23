package user_repository

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	"time"
)

func (r *UsersRepository) PatchUser(ctx context.Context, userId string, user domain.User) (domain.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	query := `
		UPDATE users 
		SET
			username=$1
		WHERE id=$2
		RETURNING
			id,
			username;
	`

	row := r.pool.QueryRow(ctx, query, user.Username, userId)

	var userModel UserModel

	if err := row.Scan(&userModel.ID, userModel.Username); err != nil {
		return domain.User{}, fmt.Errorf("scarn error : %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Username,
		"",
		"",
		"",
	)

	return userDomain, nil

}
