package graph

import (
	"github.com/JakeHumphries/investment-service/database"
	"github.com/JakeHumphries/investment-service/graph/model"
)

// MapGraphInvestmentInputToDB converts GraphQL input to a database investment.
func MapGraphInvestmentInputToDB(input *model.InvestmentInput) database.Investment {
	return database.Investment{
		CustomerID: input.CustomerID,
		FundID:     input.FundID,
		Amount:     input.Amount,
	}
}

// MapDBInvestmentsToGraphList converts database investments to GraphQL investment list.
func MapDBInvestmentsToGraphList(dbInvestments []database.Investment, nextCursor *string) *model.InvestmentList {
	graphInvestments := make([]*model.Investment, len(dbInvestments))
	for i, dbInv := range dbInvestments {
		graphInvestments[i] = &model.Investment{
			ID:        dbInv.ID,
			Amount:    dbInv.Amount,
			CreatedAt: dbInv.CreatedAt,
			Fund: &model.Fund{
				ID:   dbInv.FundID,
				Name: "Unknown", // Can be enriched if needed
			},
		}
	}

	return &model.InvestmentList{
		Investments: graphInvestments,
		NextCursor:  nextCursor,
	}
}

// TODO we currently dont get the fund as well as the investments in the db module
