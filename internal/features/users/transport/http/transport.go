package users_transport_http

import (
	"context"
	"golang-clean-server/internal/core/domain"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

type UsersHandler struct {
	usersService UsersService
	tokenManager core_service_tokens.TokenManager
}

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userId string) (domain.User, error)
	PatchUser(ctx context.Context, id string, user domain.UserPatch) (domain.User, error)
	DeleteUser(ctx context.Context, userId string) error
	LoginUser(ctx context.Context, user domain.User) (domain.User, error)
	RefreshTokens(ctx context.Context, refreshToken string) (domain.User, error)
}

func NewUsersHandlers(
	usersService UsersService,
	tokenManager core_service_tokens.TokenManager,
) *UsersHandler {
	return &UsersHandler{usersService: usersService, tokenManager: tokenManager}
}

func (h *UsersHandler) Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetAllUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.LoginUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/auth/refresh",
			Handler: h.RefreshTokens,
		},
	}
}
