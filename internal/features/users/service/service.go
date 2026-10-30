package users_service

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_hasher "golang-clean-server/internal/core/service/hasher"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
)

type UsersService struct {
	usersRepository UsersRepository
	hasher          core_service_hasher.PasswordHasher
	tokenManager    core_service_tokens.TokenManager
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, userId string) (domain.User, error)
	DeleteUser(ctx context.Context, userId string) error
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
	PatchUser(ctx context.Context, id string, patch domain.User) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
	hasher core_service_hasher.PasswordHasher,
	tokenManager core_service_tokens.TokenManager,
) *UsersService {
	return &UsersService{usersRepository: usersRepository, hasher: hasher, tokenManager: tokenManager}
}
