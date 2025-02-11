package database

import (
	"context"
	"fmt"
	"time"

	"github.com/JakeHumphries/investment-service/models"
)

// GetFunds retrieves all available funds.
func (p *PostgresClient) GetFunds(ctx context.Context, customerType string) ([]models.Fund, error) {
	query := `SELECT id, name, category, customer_type, created_at FROM fund WHERE customer_type = $1`
	rows, err := p.Database.Query(ctx, query, customerType)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch funds: %w", err)
	}
	defer rows.Close()

	var funds []models.Fund
	for rows.Next() {
		var fund models.Fund
		var createdAt time.Time

		if err := rows.Scan(&fund.ID, &fund.Name, &fund.Category, &fund.CustomerType, &createdAt); err != nil {
			return nil, fmt.Errorf("failed to scan fund: %w", err)
		}

		fund.CreatedAt = createdAt.Format(time.RFC3339)
		funds = append(funds, fund)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through funds: %w", err)
	}

	return funds, nil
}

// GetFundByID retrieves a specific fund by its ID.
func (p *PostgresClient) GetFundByID(ctx context.Context, fundID string) (*models.Fund, error) {
	query := `SELECT id, name, category, customer_type, created_at FROM fund WHERE id = $1`
	row := p.Database.QueryRow(ctx, query, fundID)

	var fund models.Fund
	var createdAt time.Time

	err := row.Scan(&fund.ID, &fund.Name, &fund.Category, &fund.CustomerType, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fund: %w", err)
	}

	fund.CreatedAt = createdAt.Format(time.RFC3339)

	return &fund, nil
}
