package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

type LoginUserDto struct {
	ID       string
	Username string
}

func (s *UsersService) LoginUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain : %w", err)
	}

	gettedUser, err := s.usersRepository.GetUserByUsername(ctx, user.Username)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user : %w", err)
	}

	if err := s.hasher.Compare(gettedUser.Password, user.Password); err != nil {
		return domain.User{}, fmt.Errorf("incorrect user data : %w", err)
	}

	return gettedUser, nil
}
