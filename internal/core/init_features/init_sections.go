package core_init_features

import (
	repository_pool "golang-clean-server/internal/core/repository/pool"
	sections_repository "golang-clean-server/internal/features/sections/repository"
	sections_service "golang-clean-server/internal/features/sections/service"
	sections_transport "golang-clean-server/internal/features/sections/transport"
)

func InitSections(
	pool *repository_pool.ConnectionPool,
) *sections_transport.SectionsHandler {

	sectionsRepository := sections_repository.NewSectionsRepository(pool)
	sectionsService := sections_service.NewSectionsService(sectionsRepository)
	sectionsTransport := sections_transport.NewSectionsHandler(sectionsService)

	return sectionsTransport
}
