package core_server

import core_middleware "golang-clean-server/internal/core/transport/middleware"

type RouteGroup struct {
	Prefix     string
	Middleware []core_middleware.Middleware
	Routes     []Route
}
