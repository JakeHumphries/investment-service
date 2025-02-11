// Package graph holds the setup and handlers for the graphQL API
package graph

import (
	"github.com/JakeHumphries/investment-service/graph/model"
	"github.com/JakeHumphries/investment-service/models"
)

// MapGraphInvestmentInputToDB converts GraphQL input to a database investment.
func MapGraphInvestmentInputToDB(input *model.InvestmentInput) models.Investment {
	return models.Investment{
		CustomerID: input.CustomerID,
		FundID:     input.FundID,
		Amount:     input.Amount,
	}
}

// MapDBInvestmentToGraph converts a DB investment to a GraphQL model.
func MapDBInvestmentToGraph(dbInvestment *models.Investment) *model.Investment {
	return &model.Investment{
		ID:        dbInvestment.ID,
		Amount:    dbInvestment.Amount,
		CreatedAt: dbInvestment.CreatedAt,
		Fund: &model.Fund{
			ID:        dbInvestment.Fund.ID,
			Name:      dbInvestment.Fund.Name,
			Category:  dbInvestment.Fund.Category,
			CreatedAt: dbInvestment.Fund.CreatedAt,
		},
	}
}

// MapDBFundsToGraphList converts a DB fund list to GraphQL.
func MapDBFundsToGraphList(dbFunds []models.Fund) *model.FundList {
	graphFunds := make([]*model.Fund, len(dbFunds))
	for i, dbFund := range dbFunds {
		graphFunds[i] = &model.Fund{
			ID:        dbFund.ID,
			Name:      dbFund.Name,
			Category:  dbFund.Category,
			CreatedAt: dbFund.CreatedAt,
		}
	}

	return &model.FundList{
		Funds: graphFunds,
	}
}

// MapDBInvestmentsToGraphList converts a DB investment list to GraphQL.
func MapDBInvestmentsToGraphList(dbInvestments []models.Investment, nextCursor *string) *model.InvestmentList {
	graphInvestments := make([]*model.Investment, len(dbInvestments))
	for i, dbInv := range dbInvestments {
		graphInvestments[i] = MapDBInvestmentToGraph(&dbInv)
	}

	return &model.InvestmentList{
		Investments: graphInvestments,
		NextCursor:  nextCursor,
	}
}
