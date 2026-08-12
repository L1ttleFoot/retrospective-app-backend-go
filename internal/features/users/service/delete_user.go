package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) DeleteUser(ctx context.Context, userId string) error {

	if err := s.usersRepository.DeleteUser(ctx, userId); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}
