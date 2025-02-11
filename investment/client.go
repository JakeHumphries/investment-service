// Package investment holds all the business logic and rules relating to investments
package investment

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/JakeHumphries/investment-service/models"
)

// Define constants for customer types
const (
	CustomerTypeRetail   = "retail"
	CustomerTypeEmployee = "employee"
)

// ClientInterface defines the business logic for investments.
//
//go:generate mockery --name ClientInterface
type ClientInterface interface {
	CreateInvestment(ctx context.Context, investment models.Investment, customerType string) (*models.Investment, error)
	GetInvestments(ctx context.Context, customerID string, encodedCursor *string, limit int) ([]models.Investment, *string, error)
	GetFunds(ctx context.Context, customerType string) ([]models.Fund, error)
}

// Client contains the business logic for managing investments.
type Client struct {
	db models.Repository
}

// NewClient returns a new instance of the investment client.
func NewClient(db models.Repository) *Client {
	return &Client{
		db: db,
	}
}

// CreateInvestment handles investment creation and delegates based on customer type.
func (c *Client) CreateInvestment(ctx context.Context, investment models.Investment, customerType string) (*models.Investment, error) {
	if investment.Amount <= 0 {
		return nil, errors.New("investment amount must be greater than zero")
	}

	fund, err := c.db.GetFundByID(ctx, investment.FundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get fund: %w", err)
	}

	if !strings.EqualFold(customerType, CustomerTypeRetail) && !strings.EqualFold(customerType, CustomerTypeEmployee) {
		return nil, fmt.Errorf("invalid customer type: %s", customerType)
	}

	if !strings.EqualFold(fund.CustomerType, customerType) {
		return nil, fmt.Errorf("%s customers cannot invest in %s funds", customerType, fund.CustomerType)
	}

	switch strings.ToLower(customerType) {
	case CustomerTypeRetail:
		return c.handleRetailInvestment(ctx, investment)
	case CustomerTypeEmployee:
		return c.handleEmployeeInvestment(ctx, investment)
	default:
		return nil, fmt.Errorf("invalid customer type: %s", customerType)
	}
}

// handleRetailInvestment processes retail investments.
func (c *Client) handleRetailInvestment(ctx context.Context, investment models.Investment) (*models.Investment, error) {
	// Add additional business logic / rules here

	return c.db.CreateInvestment(ctx, &investment)
}

// handleEmployeeInvestment processes employee investments.
func (c *Client) handleEmployeeInvestment(ctx context.Context, investment models.Investment) (*models.Investment, error) {
	// Add additional business logic / rules here

	return c.db.CreateInvestment(ctx, &investment)
}

// GetInvestments retrieves a paginated list of investments for a customer.
func (c *Client) GetInvestments(ctx context.Context, customerID string, encodedCursor *string, limit int) ([]models.Investment, *string, error) {
	var cursor *string
	if encodedCursor != nil {
		decodedCursor, err := decodeCursor(*encodedCursor)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid cursor: %w", err)
		}
		cursor = decodedCursor
	}

	investments, nextCursor, err := c.db.GetInvestments(ctx, customerID, limit, cursor)
	if err != nil {
		return nil, nil, err
	}

	var encodedNextCursor *string
	if nextCursor != nil {
		e := encodeCursor(*nextCursor)
		encodedNextCursor = &e
	}

	return investments, encodedNextCursor, nil
}

func decodeCursor(encodedCursor string) (*string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return nil, err
	}
	decodedString := string(decodedBytes)
	return &decodedString, nil
}

func encodeCursor(timestamp string) string {
	return base64.StdEncoding.EncodeToString([]byte(timestamp))
}

// GetFunds retrieves all available funds.
func (c *Client) GetFunds(ctx context.Context, customerType string) ([]models.Fund, error) {
	funds, err := c.db.GetFunds(ctx, customerType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch funds: %w", err)
	}
	return funds, nil
}
