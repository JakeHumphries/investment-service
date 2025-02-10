package investment

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/JakeHumphries/investment-service/database"
)

// Define constants for customer types
const (
	CustomerTypeRetail   = "retail"
	CustomerTypeEmployee = "employee"
)

// Define constants for fund categories
const (
	FundCategoryRetailISA       = "retail_isa"
	FundCategoryEmployeeISA     = "employee_isa"
	FundCategoryEmployeePension = "employee_pension"
)

// ClientInterface defines the business logic for investments.
//
//go:generate mockery --name ClientInterface
type ClientInterface interface {
	CreateInvestment(ctx context.Context, investment database.Investment, customerType string) (*database.Investment, error)
	GetInvestments(ctx context.Context, customerID string, encodedCursor *string, limit int) ([]database.Investment, *string, error)
}

// Client contains the business logic for managing investments.
type Client struct {
	db database.Client
}

// NewClient returns a new instance of the investment client.
func NewClient(db database.Client) *Client {
	return &Client{
		db: db,
	}
}

// CreateInvestment handles investment creation and delegates based on customer type.
func (c *Client) CreateInvestment(ctx context.Context, investment database.Investment, customerType string) (*database.Investment, error) {
	fund, err := c.db.GetFundByID(ctx, investment.FundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get fund: %w", err)
	}

	if investment.Amount <= 0 {
		return nil, fmt.Errorf("investment amount must be greater than zero")
	}

	switch strings.ToLower(customerType) {
	case CustomerTypeRetail:
		return c.handleRetailInvestment(ctx, investment, fund)
	case CustomerTypeEmployee:
		return c.handleEmployeeInvestment(ctx, investment, fund)
	default:
		return nil, fmt.Errorf("invalid customer type: %s", customerType)
	}
}

// handleRetailInvestment processes retail investments.
func (c *Client) handleRetailInvestment(ctx context.Context, investment database.Investment, fund *database.Fund) (*database.Investment, error) {
	if strings.ToLower(fund.Category) != FundCategoryRetailISA {
		return nil, fmt.Errorf("retail customers can only invest in retail ISAs")
	}

	return c.db.CreateInvestment(ctx, &investment)
}

// handleEmployeeInvestment processes employee investments.
func (c *Client) handleEmployeeInvestment(ctx context.Context, investment database.Investment, fund *database.Fund) (*database.Investment, error) {
	if strings.ToLower(fund.Category) != FundCategoryEmployeeISA && strings.ToLower(fund.Category) != FundCategoryEmployeePension {
		return nil, fmt.Errorf("employees can only invest in employer ISAs or pensions")
	}

	return c.db.CreateInvestment(ctx, &investment)
}

// GetInvestments retrieves a paginated list of investments for a customer.
func (c *Client) GetInvestments(ctx context.Context, customerID string, encodedCursor *string, limit int) ([]database.Investment, *string, error) {
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

// Encode and decode cursors for pagination.
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
