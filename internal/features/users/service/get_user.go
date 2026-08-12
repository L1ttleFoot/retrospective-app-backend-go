package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, userId string) (domain.User, error) {

	user, err := s.usersRepository.GetUser(ctx, userId)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
