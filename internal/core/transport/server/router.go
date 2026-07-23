package core_server

import (
	"fmt"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	"net/http"
)

type ApiRouter struct {
	*http.ServeMux
	Middleware []core_middleware.Middleware
}

func NewApiRouter(middleware ...core_middleware.Middleware) *ApiRouter {

	return &ApiRouter{ServeMux: http.NewServeMux(), Middleware: middleware}
}

func (r *ApiRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {

		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		h := core_middleware.ChainMiddleware(route.Handler, route.Middleware...)

		r.Handle(pattern, h)
	}
}
