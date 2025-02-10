package models

// Investment represents an investment, now independent of the database layer.
type Investment struct {
	ID         string
	CustomerID string
	FundID     string
	Amount     float64
	CreatedAt  string
	Fund       Fund
}
