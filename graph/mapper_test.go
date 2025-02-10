package graph

import (
	"testing"

	"github.com/JakeHumphries/investment-service/graph/model"
	"github.com/JakeHumphries/investment-service/models"
	"github.com/stretchr/testify/assert"
)

// Test MapGraphInvestmentInputToDB
func TestMapGraphInvestmentInputToDB(t *testing.T) {
	t.Run("success: should map GraphQL input to DB model", func(t *testing.T) {
		input := &model.InvestmentInput{
			CustomerID: "customer-123",
			FundID:     "fund-456",
			Amount:     1000,
		}

		expected := models.Investment{
			CustomerID: "customer-123",
			FundID:     "fund-456",
			Amount:     1000,
		}

		result := MapGraphInvestmentInputToDB(input)
		assert.Equal(t, expected, result)
	})
}

// Test MapDBInvestmentToGraph
func TestMapDBInvestmentToGraph(t *testing.T) {
	t.Run("success: should map DB investment to GraphQL model", func(t *testing.T) {
		dbInvestment := &models.Investment{
			ID:         "inv-123",
			CustomerID: "customer-123",
			FundID:     "fund-456",
			Amount:     5000,
			CreatedAt:  "2025-02-10T18:00:00Z",
			Fund: models.Fund{
				ID:        "fund-456",
				Name:      "Cushon Equities Fund",
				Category:  "RETAIL_ISA",
				CreatedAt: "2025-02-10T18:00:00Z",
			},
		}

		expected := &model.Investment{
			ID:        "inv-123",
			Amount:    5000,
			CreatedAt: "2025-02-10T18:00:00Z",
			Fund: &model.Fund{
				ID:        "fund-456",
				Name:      "Cushon Equities Fund",
				Category:  "RETAIL_ISA",
				CreatedAt: "2025-02-10T18:00:00Z",
			},
		}

		result := MapDBInvestmentToGraph(dbInvestment)
		assert.Equal(t, expected, result)
	})
}

// Test MapDBFundsToGraphList
func TestMapDBFundsToGraphList(t *testing.T) {
	t.Run("success: should map DB funds list to GraphQL model", func(t *testing.T) {
		dbFunds := []models.Fund{
			{ID: "fund-1", Name: "Fund One", Category: "RETAIL_ISA", CreatedAt: "2025-02-10T18:00:00Z"},
			{ID: "fund-2", Name: "Fund Two", Category: "EMPLOYEE_PENSION", CreatedAt: "2025-02-10T18:00:00Z"},
		}

		expected := &model.FundList{
			Funds: []*model.Fund{
				{ID: "fund-1", Name: "Fund One", Category: "RETAIL_ISA", CreatedAt: "2025-02-10T18:00:00Z"},
				{ID: "fund-2", Name: "Fund Two", Category: "EMPLOYEE_PENSION", CreatedAt: "2025-02-10T18:00:00Z"},
			},
		}

		result := MapDBFundsToGraphList(dbFunds)
		assert.Equal(t, expected, result)
	})

	t.Run("edge case: should handle empty list", func(t *testing.T) {
		result := MapDBFundsToGraphList([]models.Fund{})
		assert.NotNil(t, result)
		assert.Empty(t, result.Funds)
	})
}

// Test MapDBInvestmentsToGraphList
func TestMapDBInvestmentsToGraphList(t *testing.T) {
	t.Run("success: should map DB investments list to GraphQL model", func(t *testing.T) {
		dbInvestments := []models.Investment{
			{
				ID:         "inv-1",
				CustomerID: "customer-123",
				FundID:     "fund-456",
				Amount:     5000,
				CreatedAt:  "2025-02-10T18:00:00Z",
				Fund: models.Fund{
					ID:        "fund-456",
					Name:      "Cushon Equities Fund",
					Category:  "RETAIL_ISA",
					CreatedAt: "2025-02-10T18:00:00Z",
				},
			},
			{
				ID:         "inv-2",
				CustomerID: "customer-456",
				FundID:     "fund-789",
				Amount:     7000,
				CreatedAt:  "2025-02-10T18:10:00Z",
				Fund: models.Fund{
					ID:        "fund-789",
					Name:      "Cushon Pension Growth Fund",
					Category:  "EMPLOYEE_PENSION",
					CreatedAt: "2025-02-10T18:10:00Z",
				},
			},
		}

		nextCursor := "next-cursor-encoded"

		expected := &model.InvestmentList{
			Investments: []*model.Investment{
				{
					ID:        "inv-1",
					Amount:    5000,
					CreatedAt: "2025-02-10T18:00:00Z",
					Fund: &model.Fund{
						ID:        "fund-456",
						Name:      "Cushon Equities Fund",
						Category:  "RETAIL_ISA",
						CreatedAt: "2025-02-10T18:00:00Z",
					},
				},
				{
					ID:        "inv-2",
					Amount:    7000,
					CreatedAt: "2025-02-10T18:10:00Z",
					Fund: &model.Fund{
						ID:        "fund-789",
						Name:      "Cushon Pension Growth Fund",
						Category:  "EMPLOYEE_PENSION",
						CreatedAt: "2025-02-10T18:10:00Z",
					},
				},
			},
			NextCursor: &nextCursor,
		}

		result := MapDBInvestmentsToGraphList(dbInvestments, &nextCursor)
		assert.Equal(t, expected, result)
	})

	t.Run("edge case: should handle empty list", func(t *testing.T) {
		result := MapDBInvestmentsToGraphList([]models.Investment{}, nil)
		assert.NotNil(t, result)
		assert.Empty(t, result.Investments)
		assert.Nil(t, result.NextCursor)
	})
}
