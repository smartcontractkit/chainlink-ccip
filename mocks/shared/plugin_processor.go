// Code generated by mockery v2.43.2. DO NOT EDIT.

package shared

import (
	context "context"

	shared "github.com/smartcontractkit/chainlink-ccip/shared"
	mock "github.com/stretchr/testify/mock"
)

// MockPluginProcessor is an autogenerated mock type for the PluginProcessor type
type MockPluginProcessor[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	mock.Mock
}

type MockPluginProcessor_Expecter[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	mock *mock.Mock
}

func (_m *MockPluginProcessor[QueryType, ObservationType, OutcomeType]) EXPECT() *MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType] {
	return &MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType]{mock: &_m.Mock}
}

// Observation provides a mock function with given fields: ctx, prevOutcome, query
func (_m *MockPluginProcessor[QueryType, ObservationType, OutcomeType]) Observation(ctx context.Context, prevOutcome OutcomeType, query QueryType) (ObservationType, error) {
	ret := _m.Called(ctx, prevOutcome, query)

	if len(ret) == 0 {
		panic("no return value specified for Observation")
	}

	var r0 ObservationType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, OutcomeType, QueryType) (ObservationType, error)); ok {
		return rf(ctx, prevOutcome, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, OutcomeType, QueryType) ObservationType); ok {
		r0 = rf(ctx, prevOutcome, query)
	} else {
		r0 = ret.Get(0).(ObservationType)
	}

	if rf, ok := ret.Get(1).(func(context.Context, OutcomeType, QueryType) error); ok {
		r1 = rf(ctx, prevOutcome, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPluginProcessor_Observation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Observation'
type MockPluginProcessor_Observation_Call[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	*mock.Call
}

// Observation is a helper method to define mock.On call
//   - ctx context.Context
//   - prevOutcome OutcomeType
//   - query QueryType
func (_e *MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType]) Observation(ctx interface{}, prevOutcome interface{}, query interface{}) *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType] {
	return &MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType]{Call: _e.mock.On("Observation", ctx, prevOutcome, query)}
}

func (_c *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType]) Run(run func(ctx context.Context, prevOutcome OutcomeType, query QueryType)) *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(OutcomeType), args[2].(QueryType))
	})
	return _c
}

func (_c *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType]) Return(_a0 ObservationType, _a1 error) *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType]) RunAndReturn(run func(context.Context, OutcomeType, QueryType) (ObservationType, error)) *MockPluginProcessor_Observation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(run)
	return _c
}

// Outcome provides a mock function with given fields: prevOutcome, query, aos
func (_m *MockPluginProcessor[QueryType, ObservationType, OutcomeType]) Outcome(prevOutcome OutcomeType, query QueryType, aos []shared.AttributedObservation[ObservationType]) (OutcomeType, error) {
	ret := _m.Called(prevOutcome, query, aos)

	if len(ret) == 0 {
		panic("no return value specified for Outcome")
	}

	var r0 OutcomeType
	var r1 error
	if rf, ok := ret.Get(0).(func(OutcomeType, QueryType, []shared.AttributedObservation[ObservationType]) (OutcomeType, error)); ok {
		return rf(prevOutcome, query, aos)
	}
	if rf, ok := ret.Get(0).(func(OutcomeType, QueryType, []shared.AttributedObservation[ObservationType]) OutcomeType); ok {
		r0 = rf(prevOutcome, query, aos)
	} else {
		r0 = ret.Get(0).(OutcomeType)
	}

	if rf, ok := ret.Get(1).(func(OutcomeType, QueryType, []shared.AttributedObservation[ObservationType]) error); ok {
		r1 = rf(prevOutcome, query, aos)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPluginProcessor_Outcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Outcome'
type MockPluginProcessor_Outcome_Call[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	*mock.Call
}

// Outcome is a helper method to define mock.On call
//   - prevOutcome OutcomeType
//   - query QueryType
//   - aos []shared.AttributedObservation[ObservationType]
func (_e *MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType]) Outcome(prevOutcome interface{}, query interface{}, aos interface{}) *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType] {
	return &MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType]{Call: _e.mock.On("Outcome", prevOutcome, query, aos)}
}

func (_c *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType]) Run(run func(prevOutcome OutcomeType, query QueryType, aos []shared.AttributedObservation[ObservationType])) *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(OutcomeType), args[1].(QueryType), args[2].([]shared.AttributedObservation[ObservationType]))
	})
	return _c
}

func (_c *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType]) Return(_a0 OutcomeType, _a1 error) *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType]) RunAndReturn(run func(OutcomeType, QueryType, []shared.AttributedObservation[ObservationType]) (OutcomeType, error)) *MockPluginProcessor_Outcome_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, prevOutcome
func (_m *MockPluginProcessor[QueryType, ObservationType, OutcomeType]) Query(ctx context.Context, prevOutcome OutcomeType) (QueryType, error) {
	ret := _m.Called(ctx, prevOutcome)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 QueryType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, OutcomeType) (QueryType, error)); ok {
		return rf(ctx, prevOutcome)
	}
	if rf, ok := ret.Get(0).(func(context.Context, OutcomeType) QueryType); ok {
		r0 = rf(ctx, prevOutcome)
	} else {
		r0 = ret.Get(0).(QueryType)
	}

	if rf, ok := ret.Get(1).(func(context.Context, OutcomeType) error); ok {
		r1 = rf(ctx, prevOutcome)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPluginProcessor_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockPluginProcessor_Query_Call[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - prevOutcome OutcomeType
func (_e *MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType]) Query(ctx interface{}, prevOutcome interface{}) *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType] {
	return &MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType]{Call: _e.mock.On("Query", ctx, prevOutcome)}
}

func (_c *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType]) Run(run func(ctx context.Context, prevOutcome OutcomeType)) *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(OutcomeType))
	})
	return _c
}

func (_c *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType]) Return(_a0 QueryType, _a1 error) *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType]) RunAndReturn(run func(context.Context, OutcomeType) (QueryType, error)) *MockPluginProcessor_Query_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(run)
	return _c
}

// ValidateObservation provides a mock function with given fields: prevOutcome, query, ao
func (_m *MockPluginProcessor[QueryType, ObservationType, OutcomeType]) ValidateObservation(prevOutcome OutcomeType, query QueryType, ao shared.AttributedObservation[ObservationType]) error {
	ret := _m.Called(prevOutcome, query, ao)

	if len(ret) == 0 {
		panic("no return value specified for ValidateObservation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(OutcomeType, QueryType, shared.AttributedObservation[ObservationType]) error); ok {
		r0 = rf(prevOutcome, query, ao)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPluginProcessor_ValidateObservation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateObservation'
type MockPluginProcessor_ValidateObservation_Call[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}] struct {
	*mock.Call
}

// ValidateObservation is a helper method to define mock.On call
//   - prevOutcome OutcomeType
//   - query QueryType
//   - ao shared.AttributedObservation[ObservationType]
func (_e *MockPluginProcessor_Expecter[QueryType, ObservationType, OutcomeType]) ValidateObservation(prevOutcome interface{}, query interface{}, ao interface{}) *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType] {
	return &MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType]{Call: _e.mock.On("ValidateObservation", prevOutcome, query, ao)}
}

func (_c *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType]) Run(run func(prevOutcome OutcomeType, query QueryType, ao shared.AttributedObservation[ObservationType])) *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(OutcomeType), args[1].(QueryType), args[2].(shared.AttributedObservation[ObservationType]))
	})
	return _c
}

func (_c *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType]) Return(_a0 error) *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType]) RunAndReturn(run func(OutcomeType, QueryType, shared.AttributedObservation[ObservationType]) error) *MockPluginProcessor_ValidateObservation_Call[QueryType, ObservationType, OutcomeType] {
	_c.Call.Return(run)
	return _c
}

// NewMockPluginProcessor creates a new instance of MockPluginProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPluginProcessor[QueryType interface{}, ObservationType interface{}, OutcomeType interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPluginProcessor[QueryType, ObservationType, OutcomeType] {
	mock := &MockPluginProcessor[QueryType, ObservationType, OutcomeType]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}