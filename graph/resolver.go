package graph

import "go.uber.org/zap"

// Resolver is a struct for dependency injection in graphql resolvers
type Resolver struct {
	logger *zap.Logger
}

// NewResolver is a function for dependency injection in graphql resolvers
func NewResolver(logger *zap.Logger) *Resolver {
	return &Resolver{
		logger: logger,
	}
}
