package database

import (
	"context"
	"fmt"
)

// Fund represents an investment fund.
type Fund struct {
	ID        string
	Name      string
	Category  string
	CreatedAt string
}

// GetFunds retrieves all available funds.
func (p *PostgresClient) GetFunds(ctx context.Context) ([]Fund, error) {
	query := `SELECT id, name, category, created_at FROM fund ORDER BY created_at DESC`
	rows, err := p.Database.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch funds: %w", err)
	}
	defer rows.Close()

	var funds []Fund
	for rows.Next() {
		var fund Fund
		if err := rows.Scan(&fund.ID, &fund.Name, &fund.Category, &fund.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan fund: %w", err)
		}
		funds = append(funds, fund)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through funds: %w", err)
	}

	return funds, nil
}

// GetFundByID retrieves a specific fund by its ID.
func (p *PostgresClient) GetFundByID(ctx context.Context, fundID string) (*Fund, error) {
	query := `SELECT id, name, category, created_at FROM fund WHERE id = $1`
	row := p.Database.QueryRow(ctx, query, fundID)

	var fund Fund
	err := row.Scan(&fund.ID, &fund.Name, &fund.Category, &fund.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fund: %w", err)
	}

	return &fund, nil
}
