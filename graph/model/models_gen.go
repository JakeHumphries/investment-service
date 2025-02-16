// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Represents an investment fund that customers can invest in.
type Fund struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	CreatedAt string `json:"createdAt"`
}

// Represents a list of investment funds.
type FundList struct {
	// The list of funds available for investment.
	Funds []*Fund `json:"funds"`
}

// Represents an investment made by a customer.
type Investment struct {
	ID        string  `json:"id"`
	Fund      *Fund   `json:"fund"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"createdAt"`
}

// Input type for making an investment.
type InvestmentInput struct {
	// The ID of the customer making the investment.
	CustomerID string `json:"customerId"`
	// The ID of the fund the customer is investing in.
	FundID string `json:"fundId"`
	// The amount the customer wants to invest.
	Amount float64 `json:"amount"`
	// The type of customer: retail or employee.
	CustomerType CustomerType `json:"customerType"`
}

// A paginated list of investment records.
type InvestmentList struct {
	// The list of investments returned in this query.
	Investments []*Investment `json:"investments"`
	// The cursor to use for the next page of results.
	NextCursor *string `json:"nextCursor,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

// Defines the possible customer types.
type CustomerType string

const (
	CustomerTypeRetail   CustomerType = "RETAIL"
	CustomerTypeEmployee CustomerType = "EMPLOYEE"
)

var AllCustomerType = []CustomerType{
	CustomerTypeRetail,
	CustomerTypeEmployee,
}

func (e CustomerType) IsValid() bool {
	switch e {
	case CustomerTypeRetail, CustomerTypeEmployee:
		return true
	}
	return false
}

func (e CustomerType) String() string {
	return string(e)
}

func (e *CustomerType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomerType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomerType", str)
	}
	return nil
}

func (e CustomerType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
