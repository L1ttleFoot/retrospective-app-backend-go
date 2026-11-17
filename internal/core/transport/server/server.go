package core_server

import (
	"context"
	"errors"
	"fmt"
	core_logger "golang-clean-server/internal/core/logger"
	core_middleware "golang-clean-server/internal/core/transport/middleware"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_middleware.Middleware
}

func NewServer(config Config, log *core_logger.Logger, middleware ...core_middleware.Middleware) *Server {
	return &Server{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (s *Server) RegisterGroups(groups ...RouteGroup) {
	for _, group := range groups {
		for _, route := range group.Routes {
			pattern := fmt.Sprintf("%s %s%s", route.Method, group.Prefix, route.Path)
			handler := core_middleware.ChainMiddleware(route.WithMiddleware(), group.Middleware...)
			s.mux.Handle(pattern, handler)
		}
	}
}

func (s *Server) Run(ctx context.Context) error {

	mux := core_middleware.ChainMiddleware(s.mux, s.middleware...)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("start http server", zap.String("addr", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutdown http server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		s.log.Warn("http server stopped")
	}

	return nil
}
