package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/JakeHumphries/investment-service/graph/model"
)

// Invest is the resolver for the invest field.
func (r *mutationResolver) Invest(ctx context.Context, input model.InvestmentInput) (*model.Investment, error) {
	// Convert GraphQL input to DB model
	dbInvestment := MapGraphInvestmentInputToDB(&input)

	// Convert GraphQL enum to string for business logic
	customerType := strings.ToLower(input.CustomerType.String())

	// Call business logic with customerType
	createdInvestment, err := r.investmentClient.CreateInvestment(ctx, dbInvestment, customerType)
	if err != nil {
		return nil, fmt.Errorf("failed to process investment: %w", err)
	}

	// Map DB investment to GraphQL model
	return MapDBInvestmentToGraph(createdInvestment), nil
}

// GetFunds fetches all available funds.
func (r *queryResolver) GetFunds(ctx context.Context) (*model.FundList, error) {
	funds, err := r.investmentClient.GetFunds(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch funds: %w", err)
	}

	return MapDBFundsToGraphList(funds), nil
}

// GetInvestments retrieves a customer's investments with pagination.
func (r *queryResolver) GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) (*model.InvestmentList, error) {
	// Call business logic
	investments, nextCursor, err := r.investmentClient.GetInvestments(ctx, customerID, cursor, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch investments: %w", err)
	}

	// Convert to GraphQL model
	return MapDBInvestmentsToGraphList(investments, nextCursor), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

// TODO we havent done anything with jwt's yet
