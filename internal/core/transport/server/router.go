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

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *ApiRouter) WithMiddleware() http.Handler {
	return core_middleware.ChainMiddleware(
		r,
		r.Middleware...,
	)
}

func RoutesMerge(slices ...[]Route) []Route {

	totalLen := 0

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]Route, 0, totalLen)

	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}
