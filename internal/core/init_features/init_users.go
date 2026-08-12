package core_init_features

import (
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_hasher "golang-clean-server/internal/core/service/hasher"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	user_repository "golang-clean-server/internal/features/users/repository"
	users_service "golang-clean-server/internal/features/users/service"
	users_transport_http "golang-clean-server/internal/features/users/transport/http"
)

func InitUsers(
	pool *repository_pool.ConnectionPool,
	hasher *core_service_hasher.Hasher,
	tokenManager *core_service_tokens.Tokens,
) *users_transport_http.UsersHandler {

	usersRepository := user_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, hasher, tokenManager)
	usersTransport := users_transport_http.NewUsersHandler(usersService, tokenManager)

	return usersTransport
}
