package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"fmt"

	"github.com/JakeHumphries/investment-service/graph/model"
	"go.uber.org/zap"
)

// Invest is the resolver for the invest field (Mutation).
func (r *mutationResolver) Invest(ctx context.Context, input model.InvestmentInput) (*model.Investment, error) {
	r.logger.Info("Invest mutation started", zap.String("customerID", input.CustomerID), zap.String("fundID", input.FundID))

	dbInvestment, err := r.investmentClient.CreateInvestment(ctx, MapGraphInvestmentInputToDB(&input), input.CustomerType.String())
	if err != nil {
		r.logger.Error("Failed to process investment", zap.Error(err), zap.String("customerID", input.CustomerID), zap.String("fundID", input.FundID))
		return nil, fmt.Errorf("failed to process investment: %w", err)
	}

	r.logger.Info("Investment successfully created", zap.String("investmentID", dbInvestment.ID), zap.String("customerID", input.CustomerID), zap.String("fundID", input.FundID))
	return MapDBInvestmentToGraph(dbInvestment), nil
}

// GetFunds fetches all available funds (Query).
func (r *queryResolver) GetFunds(ctx context.Context, customerType model.CustomerType) (*model.FundList, error) {
	r.logger.Info("Fetching funds")

	funds, err := r.investmentClient.GetFunds(ctx, customerType.String())
	if err != nil {
		r.logger.Error("Failed to fetch funds", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch funds: %w", err)
	}

	r.logger.Info("Successfully fetched funds", zap.Int("count", len(funds)))
	return MapDBFundsToGraphList(funds), nil
}

// GetInvestments retrieves a customer's investments with pagination (Query).
func (r *queryResolver) GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) (*model.InvestmentList, error) {
	r.logger.Info("Fetching investments", zap.String("customerID", customerID), zap.Int("limit", limit))

	investments, nextCursor, err := r.investmentClient.GetInvestments(ctx, customerID, cursor, limit)
	if err != nil {
		r.logger.Error("Failed to fetch investments", zap.Error(err), zap.String("customerID", customerID))
		return nil, fmt.Errorf("failed to fetch investments: %w", err)
	}

	r.logger.Info("Successfully fetched investments", zap.String("customerID", customerID), zap.Int("investmentCount", len(investments)))
	return MapDBInvestmentsToGraphList(investments, nextCursor), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
