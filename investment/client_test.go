package investment

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JakeHumphries/investment-service/models"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

// TestCreateInvestment tests the CreateInvestment function
func TestCreateInvestment(t *testing.T) {
	ctx := context.Background()
	mockDB := &models.MockRepository{}
	client := NewClient(mockDB)

	validInvestment := models.Investment{
		CustomerID: "customer-123",
		FundID:     "fund-456",
		Amount:     25000,
	}

	validFund := &models.Fund{
		ID:           "fund-456",
		Name:         "Cushon Equities Fund",
		Category:     "ISA",
		CustomerType: CustomerTypeRetail,
	}

	t.Run("success case: it should create an investment", func(t *testing.T) {
		mockDB.EXPECT().GetFundByID(ctx, validInvestment.FundID).Return(validFund, nil).Once()
		mockDB.EXPECT().CreateInvestment(ctx, mock.AnythingOfType("*models.Investment")).Return(&validInvestment, nil).Once()

		result, err := client.CreateInvestment(ctx, validInvestment, CustomerTypeRetail)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, validInvestment.Amount, result.Amount)

		mockDB.AssertExpectations(t)
	})

	t.Run("failure case: customerType mismatch", func(t *testing.T) {
		invalidFund := &models.Fund{
			ID:           "fund-456",
			Name:         "Cushon Pension Fund",
			Category:     "PENSION",
			CustomerType: CustomerTypeEmployee,
		}
		mockDB.EXPECT().GetFundByID(ctx, validInvestment.FundID).Return(invalidFund, nil).Once()

		result, err := client.CreateInvestment(ctx, validInvestment, CustomerTypeRetail)
		assert.ErrorContains(t, err, "retail customers can only invest in employee funds")
		assert.Nil(t, result)

		mockDB.AssertExpectations(t)
	})

	t.Run("failure case: invalid customer type", func(t *testing.T) {
		mockDB.EXPECT().GetFundByID(ctx, validInvestment.FundID).Return(validFund, nil).Once()

		result, err := client.CreateInvestment(ctx, validInvestment, "invalid_customer_type")
		assert.ErrorContains(t, err, "invalid customer type")
		assert.Nil(t, result)
	})

	t.Run("failure case: invalid amount", func(t *testing.T) {
		invalidInvestment := models.Investment{
			CustomerID: "customer-123",
			FundID:     "fund-456",
			Amount:     0,
		}

		result, err := client.CreateInvestment(ctx, invalidInvestment, CustomerTypeRetail)
		assert.ErrorContains(t, err, "investment amount must be greater than zero")
		assert.Nil(t, result)
	})

	t.Run("failure case: fund validation fails", func(t *testing.T) {
		mockDB.EXPECT().GetFundByID(ctx, validInvestment.FundID).Return(nil, errors.New("fund not found")).Once()

		result, err := client.CreateInvestment(ctx, validInvestment, CustomerTypeRetail)
		assert.ErrorContains(t, err, "fund not found")
		assert.Nil(t, result)

		mockDB.AssertExpectations(t)
	})
}

// TestGetInvestments tests the GetInvestments function
func TestGetInvestments(t *testing.T) {
	ctx := context.Background()
	mockDB := &models.MockRepository{}
	client := NewClient(mockDB)

	customerID := "customer-123"
	limit := 10
	cursor := encodeCursor(time.Now().UTC().Format(time.RFC3339))
	investments := []models.Investment{
		{ID: "inv-1", CustomerID: customerID, Amount: 5000},
		{ID: "inv-2", CustomerID: customerID, Amount: 15000},
	}
	nextCursor := encodeCursor(time.Now().Add(1 * time.Hour).Format(time.RFC3339))

	t.Run("success case: it should retrieve investments with pagination", func(t *testing.T) {
		mockDB.EXPECT().GetInvestments(ctx, customerID, limit, mock.Anything).Return(investments, &nextCursor, nil).Once()

		result, _, err := client.GetInvestments(ctx, customerID, &cursor, limit)
		assert.NoError(t, err)
		assert.Equal(t, investments, result)

		mockDB.AssertExpectations(t)
	})

	t.Run("failure case: db.GetInvestments should error", func(t *testing.T) {
		mockDB.EXPECT().GetInvestments(ctx, customerID, limit, mock.Anything).Return(nil, nil, errors.New("db error")).Once()

		result, returnedCursor, err := client.GetInvestments(ctx, customerID, &cursor, limit)
		assert.ErrorContains(t, err, "db error")
		assert.Nil(t, result)
		assert.Nil(t, returnedCursor)

		mockDB.AssertExpectations(t)
	})

	t.Run("edge case: invalid cursor format", func(t *testing.T) {
		invalidCursor := "invalid_cursor"

		result, returnedCursor, err := client.GetInvestments(ctx, customerID, &invalidCursor, limit)
		assert.ErrorContains(t, err, "invalid cursor")
		assert.Nil(t, result)
		assert.Nil(t, returnedCursor)
	})
}

// TestGetFunds tests the GetFunds function
func TestGetFunds(t *testing.T) {
	ctx := context.Background()
	mockDB := &models.MockRepository{}
	client := NewClient(mockDB)

	funds := []models.Fund{
		{ID: "fund-1", Name: "Cushon Equities Fund", Category: "ISA", CustomerType: CustomerTypeRetail},
		{ID: "fund-2", Name: "Cushon Pension Fund", Category: "PENSION", CustomerType: CustomerTypeEmployee},
	}

	t.Run("success case: it should retrieve funds for Retail customers", func(t *testing.T) {
		mockDB.EXPECT().GetFunds(ctx, CustomerTypeRetail).Return([]models.Fund{funds[0]}, nil).Once()

		result, err := client.GetFunds(ctx, CustomerTypeRetail)
		assert.NoError(t, err)
		assert.Equal(t, []models.Fund{funds[0]}, result)

		mockDB.AssertExpectations(t)
	})

	t.Run("success case: it should retrieve funds for Employee customers", func(t *testing.T) {
		mockDB.EXPECT().GetFunds(ctx, CustomerTypeEmployee).Return([]models.Fund{funds[1]}, nil).Once()

		result, err := client.GetFunds(ctx, CustomerTypeEmployee)
		assert.NoError(t, err)
		assert.Equal(t, []models.Fund{funds[1]}, result)

		mockDB.AssertExpectations(t)
	})

	t.Run("failure case: db.GetFunds should error", func(t *testing.T) {
		mockDB.EXPECT().GetFunds(ctx, CustomerTypeRetail).Return(nil, errors.New("db error")).Once()

		result, err := client.GetFunds(ctx, CustomerTypeRetail)
		assert.ErrorContains(t, err, "db error")
		assert.Nil(t, result)

		mockDB.AssertExpectations(t)
	})
}
