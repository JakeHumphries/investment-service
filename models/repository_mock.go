// Code generated by mockery v2.52.1. DO NOT EDIT.

package models

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// CreateInvestment provides a mock function with given fields: ctx, investment
func (_m *MockRepository) CreateInvestment(ctx context.Context, investment *Investment) (*Investment, error) {
	ret := _m.Called(ctx, investment)

	if len(ret) == 0 {
		panic("no return value specified for CreateInvestment")
	}

	var r0 *Investment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *Investment) (*Investment, error)); ok {
		return rf(ctx, investment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *Investment) *Investment); ok {
		r0 = rf(ctx, investment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Investment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *Investment) error); ok {
		r1 = rf(ctx, investment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_CreateInvestment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInvestment'
type MockRepository_CreateInvestment_Call struct {
	*mock.Call
}

// CreateInvestment is a helper method to define mock.On call
//   - ctx context.Context
//   - investment *Investment
func (_e *MockRepository_Expecter) CreateInvestment(ctx interface{}, investment interface{}) *MockRepository_CreateInvestment_Call {
	return &MockRepository_CreateInvestment_Call{Call: _e.mock.On("CreateInvestment", ctx, investment)}
}

func (_c *MockRepository_CreateInvestment_Call) Run(run func(ctx context.Context, investment *Investment)) *MockRepository_CreateInvestment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Investment))
	})
	return _c
}

func (_c *MockRepository_CreateInvestment_Call) Return(_a0 *Investment, _a1 error) *MockRepository_CreateInvestment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_CreateInvestment_Call) RunAndReturn(run func(context.Context, *Investment) (*Investment, error)) *MockRepository_CreateInvestment_Call {
	_c.Call.Return(run)
	return _c
}

// GetFundByID provides a mock function with given fields: ctx, fundID
func (_m *MockRepository) GetFundByID(ctx context.Context, fundID string) (*Fund, error) {
	ret := _m.Called(ctx, fundID)

	if len(ret) == 0 {
		panic("no return value specified for GetFundByID")
	}

	var r0 *Fund
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Fund, error)); ok {
		return rf(ctx, fundID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Fund); ok {
		r0 = rf(ctx, fundID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Fund)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fundID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetFundByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFundByID'
type MockRepository_GetFundByID_Call struct {
	*mock.Call
}

// GetFundByID is a helper method to define mock.On call
//   - ctx context.Context
//   - fundID string
func (_e *MockRepository_Expecter) GetFundByID(ctx interface{}, fundID interface{}) *MockRepository_GetFundByID_Call {
	return &MockRepository_GetFundByID_Call{Call: _e.mock.On("GetFundByID", ctx, fundID)}
}

func (_c *MockRepository_GetFundByID_Call) Run(run func(ctx context.Context, fundID string)) *MockRepository_GetFundByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetFundByID_Call) Return(_a0 *Fund, _a1 error) *MockRepository_GetFundByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetFundByID_Call) RunAndReturn(run func(context.Context, string) (*Fund, error)) *MockRepository_GetFundByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetFunds provides a mock function with given fields: ctx, customerType
func (_m *MockRepository) GetFunds(ctx context.Context, customerType string) ([]Fund, error) {
	ret := _m.Called(ctx, customerType)

	if len(ret) == 0 {
		panic("no return value specified for GetFunds")
	}

	var r0 []Fund
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]Fund, error)); ok {
		return rf(ctx, customerType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []Fund); ok {
		r0 = rf(ctx, customerType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Fund)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, customerType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_GetFunds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFunds'
type MockRepository_GetFunds_Call struct {
	*mock.Call
}

// GetFunds is a helper method to define mock.On call
//   - ctx context.Context
//   - customerType string
func (_e *MockRepository_Expecter) GetFunds(ctx interface{}, customerType interface{}) *MockRepository_GetFunds_Call {
	return &MockRepository_GetFunds_Call{Call: _e.mock.On("GetFunds", ctx, customerType)}
}

func (_c *MockRepository_GetFunds_Call) Run(run func(ctx context.Context, customerType string)) *MockRepository_GetFunds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepository_GetFunds_Call) Return(_a0 []Fund, _a1 error) *MockRepository_GetFunds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_GetFunds_Call) RunAndReturn(run func(context.Context, string) ([]Fund, error)) *MockRepository_GetFunds_Call {
	_c.Call.Return(run)
	return _c
}

// GetInvestments provides a mock function with given fields: ctx, customerID, limit, cursor
func (_m *MockRepository) GetInvestments(ctx context.Context, customerID string, limit int, cursor *string) ([]Investment, *string, error) {
	ret := _m.Called(ctx, customerID, limit, cursor)

	if len(ret) == 0 {
		panic("no return value specified for GetInvestments")
	}

	var r0 []Investment
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, *string) ([]Investment, *string, error)); ok {
		return rf(ctx, customerID, limit, cursor)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, *string) []Investment); ok {
		r0 = rf(ctx, customerID, limit, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Investment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, *string) *string); ok {
		r1 = rf(ctx, customerID, limit, cursor)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int, *string) error); ok {
		r2 = rf(ctx, customerID, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRepository_GetInvestments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInvestments'
type MockRepository_GetInvestments_Call struct {
	*mock.Call
}

// GetInvestments is a helper method to define mock.On call
//   - ctx context.Context
//   - customerID string
//   - limit int
//   - cursor *string
func (_e *MockRepository_Expecter) GetInvestments(ctx interface{}, customerID interface{}, limit interface{}, cursor interface{}) *MockRepository_GetInvestments_Call {
	return &MockRepository_GetInvestments_Call{Call: _e.mock.On("GetInvestments", ctx, customerID, limit, cursor)}
}

func (_c *MockRepository_GetInvestments_Call) Run(run func(ctx context.Context, customerID string, limit int, cursor *string)) *MockRepository_GetInvestments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(*string))
	})
	return _c
}

func (_c *MockRepository_GetInvestments_Call) Return(_a0 []Investment, _a1 *string, _a2 error) *MockRepository_GetInvestments_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockRepository_GetInvestments_Call) RunAndReturn(run func(context.Context, string, int, *string) ([]Investment, *string, error)) *MockRepository_GetInvestments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
