package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain : %w", err)
	}

	hashedPassword, err := s.hasher.Hash(user.Password)

	if err != nil {
		return domain.User{}, fmt.Errorf("hash password : %w", err)
	}

	user.Password = hashedPassword

	createdUser, err := s.usersRepository.CreateUser(ctx, user)

	if err != nil {
		return domain.User{}, fmt.Errorf("create user : %w", err)
	}

	return createdUser, nil
}
