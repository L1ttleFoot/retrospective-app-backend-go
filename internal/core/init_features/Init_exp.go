package core_init_features

import (
	exp_gateway "golang-clean-server/internal/features/experiments/gateway"
	exp_service "golang-clean-server/internal/features/experiments/service"
	exp_transport "golang-clean-server/internal/features/experiments/transport"
	"net/http"
	"time"
)

func InitExp() *exp_transport.ExpHandler {

	baseUrl := "http://exp:8090/http"
	client := &http.Client{Timeout: 10 * time.Second}

	expRepository := exp_gateway.NewExpRepository(baseUrl, client)
	expService := exp_service.NewExpService(expRepository)
	expTransport := exp_transport.NewExpTransport(expService)

	return expTransport
}
