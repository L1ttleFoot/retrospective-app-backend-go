package prometheus_transport

import (
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusHandler struct {
}

func NewPrometheusHandler() *PrometheusHandler {
	return &PrometheusHandler{}
}

func (h *PrometheusHandler) Routes() []core_server.Route {
	return []core_server.Route{{
		Method:  http.MethodGet,
		Path:    "/metrics",
		Handler: promhttp.Handler().ServeHTTP,
	}}
}
