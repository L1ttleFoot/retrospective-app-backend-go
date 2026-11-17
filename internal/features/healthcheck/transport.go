package healthcheck

import (
	core_server "golang-clean-server/internal/core/transport/server"
	"net/http"
)

func Routes() []core_server.Route {
	return []core_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/health",
			Handler: Healthcheck,
		},
	}
}
