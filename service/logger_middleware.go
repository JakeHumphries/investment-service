package service

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// NewLoggerMiddleware will create a middleware that logs every request except the ones in the excludedPaths
func NewLoggerMiddleware(logger *zap.Logger, excludedPaths []string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Skip execution for excluded paths
			for _, path := range excludedPaths {
				if r.URL.Path == path {
					next.ServeHTTP(w, r)
					return
				}
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			// defer ensures the log happens after the request has happened
			defer func() {
				logger.Info("request",
					zap.String("method", r.Method),
					zap.String("proto", r.Proto),
					zap.Duration("duration", time.Since(t1)),
					zap.String("path", r.URL.Path),
					zap.Int("status", ww.Status()),
					zap.Int("responseSize", ww.BytesWritten()),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
