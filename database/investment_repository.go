package database

import (
	"context"
	"fmt"
	"time"

	"github.com/JakeHumphries/investment-service/models"
	"github.com/jackc/pgx/v5"
)

// CreateInvestment inserts a new investment in a transaction.
func (p *PostgresClient) CreateInvestment(ctx context.Context, investment *models.Investment) (*models.Investment, error) {
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

	var createdAt time.Time

	query := `
		INSERT INTO investment (customer_id, fund_id, amount)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	row := tx.QueryRow(ctx, query, investment.CustomerID, investment.FundID, investment.Amount)

	err = row.Scan(&investment.ID, &createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert investment: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	investment.CreatedAt = createdAt.Format(time.RFC3339)

	fund, err := p.GetFundByID(ctx, investment.FundID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fund details: %w", err)
	}

	investment.Fund = *fund

	return investment, nil
}

// GetInvestments fetches a customer's investments with pagination.
func (p *PostgresClient) GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) ([]models.Investment, *string, error) {
	args := []interface{}{customerID}
	query := `
		SELECT i.id, i.customer_id, i.fund_id, i.amount, i.created_at,
		       f.id, f.name, f.category, f.created_at
		FROM investment i
		JOIN fund f ON i.fund_id = f.id
		WHERE i.customer_id = $1`

	argIndex := 2
	if cursor != nil {
		query += fmt.Sprintf(" AND i.created_at < $%d", argIndex)
		args = append(args, *cursor)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY i.created_at DESC LIMIT $%d", argIndex)
	args = append(args, limit+1)

	rows, err := p.Database.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch investments: %w", err)
	}
	defer rows.Close()

	var investments []models.Investment
	var nextCursor *string

	for rows.Next() {
		var inv models.Investment
		var fund models.Fund
		var createdAt time.Time
		var fundCreatedAt time.Time

		if err := rows.Scan(
			&inv.ID, &inv.CustomerID, &inv.FundID, &inv.Amount, &createdAt,
			&fund.ID, &fund.Name, &fund.Category, &fundCreatedAt,
		); err != nil {
			return nil, nil, fmt.Errorf("failed to scan investment: %w", err)
		}

		inv.CreatedAt = createdAt.Format(time.RFC3339)
		fund.CreatedAt = fundCreatedAt.Format(time.RFC3339)

		inv.Fund = fund
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
