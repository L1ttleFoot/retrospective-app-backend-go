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
		return domain.User{}, fmt.Errorf("Invalid or expired refresh token: %w", err)
	}

	domainUser, err := s.usersRepository.GetUser(ctx, claims.UserID)
	if err != nil {
		return domain.User{}, fmt.Errorf("refresh tokens: %w", err)
	}

	accessClaims := core_service_tokens.TokenClaims{UserID: claims.UserID, Roles: domainUser.Roles}

	newAccessToken, err := s.tokenManager.Generate(accessClaims, core_service_tokens.TypeAccess, 15*time.Minute)
	if err != nil {
		return domain.User{}, fmt.Errorf("Unable to generate access token: %w", err)
	}

	refreshClaims := core_service_tokens.TokenClaims{UserID: claims.UserID}

	newRefreshToken, err := s.tokenManager.Generate(refreshClaims, core_service_tokens.TypeRefresh, 15*24*time.Hour)
	if err != nil {
		return domain.User{}, fmt.Errorf("Unable to generate refresh token: %w", err)
	}

	domainUser = domain.NewUser(
		domainUser.ID,
		domainUser.Username,
		"",
		newAccessToken,
		newRefreshToken,
		domainUser.Roles,
	)

	return domainUser, nil
}
