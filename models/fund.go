// Package models holds the struct definitions for the service
package models

// Fund represents a financial fund available for investment.
type Fund struct {
	ID           string
	Name         string
	Category     string
	CustomerType string
	CreatedAt    string
}
