package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	"time"
)

type LoginUserDto struct {
	ID       string
	Username string
}

func (s *UsersService) LoginUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain : %w", err)
	}

	domainUser, err := s.usersRepository.GetUserByUsername(ctx, user.Username)

	if err != nil {
		return domain.User{}, fmt.Errorf("get user : %w", err)
	}

	if err := s.hasher.Compare(domainUser.Password, user.Password); err != nil {
		return domain.User{}, fmt.Errorf("incorrect user data : %w", err)
	}

	claims := core_service_tokens.TokenClaims{
		UserID: domainUser.ID,
	}

	accessToken, err := s.tokenManager.Generate(claims, core_service_tokens.TypeAccess, 15*time.Minute)
	if err != nil {
		return domain.User{}, fmt.Errorf("generate token : %w", err)
	}

	refreshToken, err := s.tokenManager.Generate(claims, core_service_tokens.TypeRefresh, 15*24*time.Hour)
	if err != nil {
		return domain.User{}, fmt.Errorf("generate token : %w", err)
	}

	domainUser = domain.NewUser(
		domainUser.ID,
		domainUser.Username,
		"",
		accessToken,
		refreshToken,
	)

	return domainUser, nil
}
