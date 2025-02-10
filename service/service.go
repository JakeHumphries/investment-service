// Package service contains definitions of the basic objects that make up the service
package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/JakeHumphries/investment-service/config"
)

// HTTPServer encapsulates two http server operations  that we need to execute in the service
type HTTPServer interface {
	Shutdown(ctx context.Context) error
	ListenAndServe() error
}

// Service holds configurations and set up needed to run the service
type Service struct {
	config      config.Config
	stopChannel chan bool
	httpServer  HTTPServer
	logger      *zap.Logger
}

// NewService creates a new instance of the service struct that instanciates the service dependencies
func NewService(ctx context.Context) (*Service, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	cfg := config.NewConfig()

	s := &Service{
		config:      cfg,
		logger:      logger,
		stopChannel: make(chan bool),
	}

	s.httpServer = &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
	}

	return s, nil
}

// Start the service will kick-start http server
func (s *Service) Start() error {
	go func() {
		s.logger.Info("Starting HTTP server", zap.Int("port", s.config.Port))

		if err := s.httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				s.logger.Error("Failed to start http listener. Stopping the service", zap.Error(err))

				s.stopChannel <- true
			}
		}
	}()

	return nil
}

// Wait will keep the service in a wait state until a message is received in the stopChannel
func (s *Service) Wait() {
	s.logger.Info("Waiting on the service")
	<-s.stopChannel
}

// Close will wait for error in error channel or signal interrupt in signal channel
func (s *Service) Close(ctx context.Context) {
	s.logger.Info("Closing the service")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("Failed to shutdown http server", zap.Error(err))
	}

	close(s.stopChannel)
}
