package service

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/JakeHumphries/investment-service/graph"
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// newRouter creates a router with needed middlewares.
func (s *Service) newRouter() http.Handler {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.RequestID)
	chiRouter.Use(NewLoggerMiddleware(s.logger, []string{"/health"}))
	chiRouter.Use(middleware.Recoverer)
	chiRouter.Use(middleware.Timeout(60 * time.Second))

	chiRouter.Handle("GET /health", healthcheck.Handler())

	resolver := graph.NewResolver(s.logger, s.investmentClient)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})

	chiRouter.Handle("/graphql", otelhttp.NewHandler(srv, "graphql"))

	return chiRouter
}
