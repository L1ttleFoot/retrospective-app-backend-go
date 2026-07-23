package users_service

import (
	"context"
	"fmt"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	"time"
)

func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (domain.User, error) {

	claims, err := s.tokenManager.Parse(refreshToken, core_service_tokens.TypeRefresh)
	if err != nil {
		return domain.User{}, fmt.Errorf("Invalid or expired refresh token : %w", err)
	}

	newClaims := core_service_tokens.TokenClaims{UserID: claims.UserID}

	domainUser, err := s.usersRepository.GetUser(ctx, claims.UserID)

	newAccessToken, err := s.tokenManager.Generate(newClaims, core_service_tokens.TypeAccess, 15*time.Minute)
	if err != nil {

		return domain.User{}, fmt.Errorf("Unable to generate token : %w", err)
	}

	newRefreshToken, err := s.tokenManager.Generate(newClaims, core_service_tokens.TypeRefresh, 15*24*time.Hour)
	if err != nil {
		return domain.User{}, fmt.Errorf("Unable to generate token : %w", err)
	}

	domainUser = domain.NewUser(
		domainUser.ID,
		domainUser.Username,
		"",
		newAccessToken,
		newRefreshToken,
	)

	return domainUser, nil
}
