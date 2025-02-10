package models

import "context"

// Repository defines the interface for database operations.
//
//go:generate mockery --name Repository
type Repository interface {
	GetFunds(ctx context.Context) ([]Fund, error)
	GetFundByID(ctx context.Context, fundID string) (*Fund, error)

	CreateInvestment(ctx context.Context, investment *Investment) (*Investment, error)
	GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) ([]Investment, *string, error)
}
