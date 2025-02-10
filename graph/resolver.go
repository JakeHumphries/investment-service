package graph

import (
	"github.com/JakeHumphries/investment-service/investment"
	"go.uber.org/zap"
)

// Resolver is a struct for dependency injection in graphql resolvers
type Resolver struct {
	logger           *zap.Logger
	investmentClient investment.ClientInterface
}

// NewResolver is a function for dependency injection in graphql resolvers
func NewResolver(logger *zap.Logger, investmentClient investment.ClientInterface) *Resolver {
	return &Resolver{
		logger:           logger,
		investmentClient: investmentClient,
	}
}
