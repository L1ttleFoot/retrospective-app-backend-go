package core_server

import (
	"fmt"
	"net/http"
)

type ApiRouter struct {
	*http.ServeMux
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{ServeMux: http.NewServeMux()}
}

func (r *ApiRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {

		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.Handler)
	}
}
