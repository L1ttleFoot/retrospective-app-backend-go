package core_init_features

import (
	repository_pool "golang-clean-server/internal/core/repository/pool"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	reports_repository "golang-clean-server/internal/features/reports/repository"
	reports_service "golang-clean-server/internal/features/reports/service"
	reports_transport "golang-clean-server/internal/features/reports/transport"
)

func InitReports(
	pool *repository_pool.ConnectionPool,
	tokenManager core_service_tokens.TokenManager,
) *reports_transport.ReportsHandler {

	reportsRepository := reports_repository.NewReportsRepository(pool)
	reportsService := reports_service.NewReportsService(reportsRepository)
	reportsTransport := reports_transport.NewReportsHandler(reportsService, tokenManager)

	return reportsTransport
}
