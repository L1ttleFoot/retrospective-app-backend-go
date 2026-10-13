package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, userId string, patch domain.UserPatch) (domain.User, error) {

	if err := patch.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user patch: %w", err)
	}

	user, err := s.usersRepository.GetUser(ctx, userId)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, userId, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, err
}
