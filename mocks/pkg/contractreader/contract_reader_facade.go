// Code generated by mockery v2.52.3. DO NOT EDIT.

package contractreader

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitives "github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	query "github.com/smartcontractkit/chainlink-common/pkg/types/query"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// MockContractReaderFacade is an autogenerated mock type for the ContractReaderFacade type
type MockContractReaderFacade struct {
	mock.Mock
}

type MockContractReaderFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *MockContractReaderFacade) EXPECT() *MockContractReaderFacade_Expecter {
	return &MockContractReaderFacade_Expecter{mock: &_m.Mock}
}

// BatchGetLatestValues provides a mock function with given fields: ctx, request
func (_m *MockContractReaderFacade) BatchGetLatestValues(ctx context.Context, request types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetLatestValues")
	}

	var r0 types.BatchGetLatestValuesResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.BatchGetLatestValuesRequest) types.BatchGetLatestValuesResult); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.BatchGetLatestValuesResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.BatchGetLatestValuesRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContractReaderFacade_BatchGetLatestValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGetLatestValues'
type MockContractReaderFacade_BatchGetLatestValues_Call struct {
	*mock.Call
}

// BatchGetLatestValues is a helper method to define mock.On call
//   - ctx context.Context
//   - request types.BatchGetLatestValuesRequest
func (_e *MockContractReaderFacade_Expecter) BatchGetLatestValues(ctx interface{}, request interface{}) *MockContractReaderFacade_BatchGetLatestValues_Call {
	return &MockContractReaderFacade_BatchGetLatestValues_Call{Call: _e.mock.On("BatchGetLatestValues", ctx, request)}
}

func (_c *MockContractReaderFacade_BatchGetLatestValues_Call) Run(run func(ctx context.Context, request types.BatchGetLatestValuesRequest)) *MockContractReaderFacade_BatchGetLatestValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.BatchGetLatestValuesRequest))
	})
	return _c
}

func (_c *MockContractReaderFacade_BatchGetLatestValues_Call) Return(_a0 types.BatchGetLatestValuesResult, _a1 error) *MockContractReaderFacade_BatchGetLatestValues_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContractReaderFacade_BatchGetLatestValues_Call) RunAndReturn(run func(context.Context, types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error)) *MockContractReaderFacade_BatchGetLatestValues_Call {
	_c.Call.Return(run)
	return _c
}

// Bind provides a mock function with given fields: ctx, bindings
func (_m *MockContractReaderFacade) Bind(ctx context.Context, bindings []types.BoundContract) error {
	ret := _m.Called(ctx, bindings)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BoundContract) error); ok {
		r0 = rf(ctx, bindings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContractReaderFacade_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type MockContractReaderFacade_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - ctx context.Context
//   - bindings []types.BoundContract
func (_e *MockContractReaderFacade_Expecter) Bind(ctx interface{}, bindings interface{}) *MockContractReaderFacade_Bind_Call {
	return &MockContractReaderFacade_Bind_Call{Call: _e.mock.On("Bind", ctx, bindings)}
}

func (_c *MockContractReaderFacade_Bind_Call) Run(run func(ctx context.Context, bindings []types.BoundContract)) *MockContractReaderFacade_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BoundContract))
	})
	return _c
}

func (_c *MockContractReaderFacade_Bind_Call) Return(_a0 error) *MockContractReaderFacade_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContractReaderFacade_Bind_Call) RunAndReturn(run func(context.Context, []types.BoundContract) error) *MockContractReaderFacade_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestValue provides a mock function with given fields: ctx, readIdentifier, confidenceLevel, params, returnVal
func (_m *MockContractReaderFacade) GetLatestValue(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{}) error {
	ret := _m.Called(ctx, readIdentifier, confidenceLevel, params, returnVal)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestValue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, primitives.ConfidenceLevel, interface{}, interface{}) error); ok {
		r0 = rf(ctx, readIdentifier, confidenceLevel, params, returnVal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContractReaderFacade_GetLatestValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestValue'
type MockContractReaderFacade_GetLatestValue_Call struct {
	*mock.Call
}

// GetLatestValue is a helper method to define mock.On call
//   - ctx context.Context
//   - readIdentifier string
//   - confidenceLevel primitives.ConfidenceLevel
//   - params interface{}
//   - returnVal interface{}
func (_e *MockContractReaderFacade_Expecter) GetLatestValue(ctx interface{}, readIdentifier interface{}, confidenceLevel interface{}, params interface{}, returnVal interface{}) *MockContractReaderFacade_GetLatestValue_Call {
	return &MockContractReaderFacade_GetLatestValue_Call{Call: _e.mock.On("GetLatestValue", ctx, readIdentifier, confidenceLevel, params, returnVal)}
}

func (_c *MockContractReaderFacade_GetLatestValue_Call) Run(run func(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{})) *MockContractReaderFacade_GetLatestValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(primitives.ConfidenceLevel), args[3].(interface{}), args[4].(interface{}))
	})
	return _c
}

func (_c *MockContractReaderFacade_GetLatestValue_Call) Return(_a0 error) *MockContractReaderFacade_GetLatestValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContractReaderFacade_GetLatestValue_Call) RunAndReturn(run func(context.Context, string, primitives.ConfidenceLevel, interface{}, interface{}) error) *MockContractReaderFacade_GetLatestValue_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with no fields
func (_m *MockContractReaderFacade) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// MockContractReaderFacade_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type MockContractReaderFacade_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *MockContractReaderFacade_Expecter) HealthReport() *MockContractReaderFacade_HealthReport_Call {
	return &MockContractReaderFacade_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *MockContractReaderFacade_HealthReport_Call) Run(run func()) *MockContractReaderFacade_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContractReaderFacade_HealthReport_Call) Return(_a0 map[string]error) *MockContractReaderFacade_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContractReaderFacade_HealthReport_Call) RunAndReturn(run func() map[string]error) *MockContractReaderFacade_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// QueryKey provides a mock function with given fields: ctx, contract, filter, limitAndSort, sequenceDataType
func (_m *MockContractReaderFacade) QueryKey(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{}) ([]types.Sequence, error) {
	ret := _m.Called(ctx, contract, filter, limitAndSort, sequenceDataType)

	if len(ret) == 0 {
		panic("no return value specified for QueryKey")
	}

	var r0 []types.Sequence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.BoundContract, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)); ok {
		return rf(ctx, contract, filter, limitAndSort, sequenceDataType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.BoundContract, query.KeyFilter, query.LimitAndSort, interface{}) []types.Sequence); ok {
		r0 = rf(ctx, contract, filter, limitAndSort, sequenceDataType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Sequence)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.BoundContract, query.KeyFilter, query.LimitAndSort, interface{}) error); ok {
		r1 = rf(ctx, contract, filter, limitAndSort, sequenceDataType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContractReaderFacade_QueryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryKey'
type MockContractReaderFacade_QueryKey_Call struct {
	*mock.Call
}

// QueryKey is a helper method to define mock.On call
//   - ctx context.Context
//   - contract types.BoundContract
//   - filter query.KeyFilter
//   - limitAndSort query.LimitAndSort
//   - sequenceDataType interface{}
func (_e *MockContractReaderFacade_Expecter) QueryKey(ctx interface{}, contract interface{}, filter interface{}, limitAndSort interface{}, sequenceDataType interface{}) *MockContractReaderFacade_QueryKey_Call {
	return &MockContractReaderFacade_QueryKey_Call{Call: _e.mock.On("QueryKey", ctx, contract, filter, limitAndSort, sequenceDataType)}
}

func (_c *MockContractReaderFacade_QueryKey_Call) Run(run func(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{})) *MockContractReaderFacade_QueryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.BoundContract), args[2].(query.KeyFilter), args[3].(query.LimitAndSort), args[4].(interface{}))
	})
	return _c
}

func (_c *MockContractReaderFacade_QueryKey_Call) Return(_a0 []types.Sequence, _a1 error) *MockContractReaderFacade_QueryKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContractReaderFacade_QueryKey_Call) RunAndReturn(run func(context.Context, types.BoundContract, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)) *MockContractReaderFacade_QueryKey_Call {
	_c.Call.Return(run)
	return _c
}

// Unbind provides a mock function with given fields: ctx, bindings
func (_m *MockContractReaderFacade) Unbind(ctx context.Context, bindings []types.BoundContract) error {
	ret := _m.Called(ctx, bindings)

	if len(ret) == 0 {
		panic("no return value specified for Unbind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.BoundContract) error); ok {
		r0 = rf(ctx, bindings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContractReaderFacade_Unbind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unbind'
type MockContractReaderFacade_Unbind_Call struct {
	*mock.Call
}

// Unbind is a helper method to define mock.On call
//   - ctx context.Context
//   - bindings []types.BoundContract
func (_e *MockContractReaderFacade_Expecter) Unbind(ctx interface{}, bindings interface{}) *MockContractReaderFacade_Unbind_Call {
	return &MockContractReaderFacade_Unbind_Call{Call: _e.mock.On("Unbind", ctx, bindings)}
}

func (_c *MockContractReaderFacade_Unbind_Call) Run(run func(ctx context.Context, bindings []types.BoundContract)) *MockContractReaderFacade_Unbind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BoundContract))
	})
	return _c
}

func (_c *MockContractReaderFacade_Unbind_Call) Return(_a0 error) *MockContractReaderFacade_Unbind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContractReaderFacade_Unbind_Call) RunAndReturn(run func(context.Context, []types.BoundContract) error) *MockContractReaderFacade_Unbind_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockContractReaderFacade creates a new instance of MockContractReaderFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockContractReaderFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockContractReaderFacade {
	mock := &MockContractReaderFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
