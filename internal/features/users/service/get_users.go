package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_errors "golang-clean-server/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {

	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non negative : %w", core_errors.ErrInvalideArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non negative : %w", core_errors.ErrInvalideArgument)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)

	if err != nil {
		return []domain.User{}, fmt.Errorf("get users : %w", err)
	}

	return users, nil
}
