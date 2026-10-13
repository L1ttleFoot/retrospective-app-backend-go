package core_init_features

import (
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	templates_repository "golang-clean-server/internal/features/templates/repository"
	templates_service "golang-clean-server/internal/features/templates/service"
	templates_transport "golang-clean-server/internal/features/templates/transport"
)

func InitTemplates(
	pool *repository_pool.ConnectionPool,
	tokenManager *core_service_tokens.Tokens,
) *templates_transport.TemplatesHandler {

	templatesRepository := templates_repository.NewTemplatesRepository(pool)
	templatesService := templates_service.NewTemplateService(templatesRepository)
	templatesTransport := templates_transport.NewTemplatesHandler(templatesService, tokenManager)

	return templatesTransport
}
