package core_middleware

import (
	"context"
	"fmt"
	core_errors "golang-clean-server/internal/core/errors"
	core_logger "golang-clean-server/internal/core/logger"
	core_service_tokens "golang-clean-server/internal/core/service/tokens"
	core_response "golang-clean-server/internal/core/transport/response"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

func Auth(tm core_service_tokens.TokenManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			log := core_logger.FromContext(ctx)

			authHeader := r.Header.Get("Authorization")

			responseHandler := core_response.NewResponseHandler(log, w)

			if authHeader == "" {
				responseHandler.ErrorResponse(core_errors.ErrForbidden, "missing token")
				return
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				responseHandler.ErrorResponse(core_errors.ErrForbidden, "invalid token format")
				return
			}

			claims, err := tm.Parse(parts[1], core_service_tokens.TypeAccess)
			if err != nil {
				responseHandler.ErrorResponse(core_errors.ErrUnauthorized, fmt.Errorf("token validate error: %w", err).Error())
				return
			}

			ctx = context.WithValue(r.Context(), "claims", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")

			if requestID == "" {
				requestID = uuid.NewString()

			}

			r.Header.Set("X-Request-ID", requestID)
			w.Header().Set("X-Request-ID", requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			log := core_logger.FromContext(ctx)
			rw := core_response.NewResponseWriter(w)

			before := time.Now()

			log.Debug(
				">>>>> incoming http request",
				zap.String("http method", r.Method),
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<<<< done http request",
				zap.Int("status code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(before)),
			)
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			log := core_logger.FromContext(ctx)

			responseHandler := core_response.NewResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "got unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_response_time_seconds",
	Help:    "Время ответа сервера на HTTP запросы",
	Buckets: prometheus.DefBuckets,
}, []string{"path", "method"})

var httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Общее количество HTTP запросов",
}, []string{"path", "method"})

func Metrics() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/metrics" {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(start).Seconds()

			httpDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
			httpRequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()
		})
	}
}
