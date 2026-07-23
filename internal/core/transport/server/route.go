package core_server

import (
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	"net/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_middleware.Middleware
}
