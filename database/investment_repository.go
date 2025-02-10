package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Investment represents an investment record.
type Investment struct {
	ID         string
	CustomerID string
	FundID     string
	Amount     float64
	CreatedAt  string
}

// TODO - SHOULD THESE STRUCT DEFINITIONS BE IN A TYPES PACKAGE TO AVOID DEPENDENCY IN SAY THE GRAPH PACKAGE

// CreateInvestment inserts a new investment in a transaction.
func (p *PostgresClient) CreateInvestment(ctx context.Context, investment *Investment) (*Investment, error) {
	tx, err := p.Database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				fmt.Printf("failed to rollback transaction: %v\n", rbErr)
			}
		}
	}()

	query := `
		INSERT INTO investment (customer_id, fund_id, amount)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	row := tx.QueryRow(ctx, query, investment.CustomerID, investment.FundID, investment.Amount)

	err = row.Scan(&investment.ID, &investment.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert investment: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return investment, nil
}

// GetInvestments fetches a customer's investments with pagination.
func (p *PostgresClient) GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) ([]Investment, *string, error) {
	args := []interface{}{customerID}
	query := `
		SELECT id, fund_id, amount, created_at
		FROM investment
		WHERE customer_id = $1`

	argIndex := 2

	if cursor != nil {
		query += fmt.Sprintf(" AND created_at < $%d", argIndex)
		args = append(args, *cursor)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d", argIndex)
	args = append(args, limit+1)

	rows, err := p.Database.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch investments: %w", err)
	}
	defer rows.Close()

	var investments []Investment
	var nextCursor *string

	for rows.Next() {
		var inv Investment
		if err := rows.Scan(&inv.ID, &inv.FundID, &inv.Amount, &inv.CreatedAt); err != nil {
			return nil, nil, fmt.Errorf("failed to scan investment: %w", err)
		}
		investments = append(investments, inv)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, fmt.Errorf("error iterating through investments: %w", err)
	}

	if len(investments) > limit {
		lastInvestment := investments[limit-1]
		nextCursor = &lastInvestment.CreatedAt
		investments = investments[:limit]
	}

	return investments, nextCursor, nil
}
