// Code generated by mockery v2.43.2. DO NOT EDIT.

package contractreader

import (
	context "context"

	contractreader "github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	mock "github.com/stretchr/testify/mock"

	primitives "github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	query "github.com/smartcontractkit/chainlink-common/pkg/types/query"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// MockExtended is an autogenerated mock type for the Extended type
type MockExtended struct {
	mock.Mock
}

type MockExtended_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExtended) EXPECT() *MockExtended_Expecter {
	return &MockExtended_Expecter{mock: &_m.Mock}
}

// BatchGetLatestValues provides a mock function with given fields: ctx, request
func (_m *MockExtended) BatchGetLatestValues(ctx context.Context, request types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error) {
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

// MockExtended_BatchGetLatestValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGetLatestValues'
type MockExtended_BatchGetLatestValues_Call struct {
	*mock.Call
}

// BatchGetLatestValues is a helper method to define mock.On call
//   - ctx context.Context
//   - request types.BatchGetLatestValuesRequest
func (_e *MockExtended_Expecter) BatchGetLatestValues(ctx interface{}, request interface{}) *MockExtended_BatchGetLatestValues_Call {
	return &MockExtended_BatchGetLatestValues_Call{Call: _e.mock.On("BatchGetLatestValues", ctx, request)}
}

func (_c *MockExtended_BatchGetLatestValues_Call) Run(run func(ctx context.Context, request types.BatchGetLatestValuesRequest)) *MockExtended_BatchGetLatestValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.BatchGetLatestValuesRequest))
	})
	return _c
}

func (_c *MockExtended_BatchGetLatestValues_Call) Return(_a0 types.BatchGetLatestValuesResult, _a1 error) *MockExtended_BatchGetLatestValues_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtended_BatchGetLatestValues_Call) RunAndReturn(run func(context.Context, types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error)) *MockExtended_BatchGetLatestValues_Call {
	_c.Call.Return(run)
	return _c
}

// Bind provides a mock function with given fields: ctx, bindings
func (_m *MockExtended) Bind(ctx context.Context, bindings []types.BoundContract) error {
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

// MockExtended_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type MockExtended_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - ctx context.Context
//   - bindings []types.BoundContract
func (_e *MockExtended_Expecter) Bind(ctx interface{}, bindings interface{}) *MockExtended_Bind_Call {
	return &MockExtended_Bind_Call{Call: _e.mock.On("Bind", ctx, bindings)}
}

func (_c *MockExtended_Bind_Call) Run(run func(ctx context.Context, bindings []types.BoundContract)) *MockExtended_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BoundContract))
	})
	return _c
}

func (_c *MockExtended_Bind_Call) Return(_a0 error) *MockExtended_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Bind_Call) RunAndReturn(run func(context.Context, []types.BoundContract) error) *MockExtended_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockExtended) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExtended_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockExtended_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockExtended_Expecter) Close() *MockExtended_Close_Call {
	return &MockExtended_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockExtended_Close_Call) Run(run func()) *MockExtended_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtended_Close_Call) Return(_a0 error) *MockExtended_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Close_Call) RunAndReturn(run func() error) *MockExtended_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetBindings provides a mock function with given fields: contractName
func (_m *MockExtended) GetBindings(contractName string) []contractreader.ExtendedBoundContract {
	ret := _m.Called(contractName)

	if len(ret) == 0 {
		panic("no return value specified for GetBindings")
	}

	var r0 []contractreader.ExtendedBoundContract
	if rf, ok := ret.Get(0).(func(string) []contractreader.ExtendedBoundContract); ok {
		r0 = rf(contractName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contractreader.ExtendedBoundContract)
		}
	}

	return r0
}

// MockExtended_GetBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBindings'
type MockExtended_GetBindings_Call struct {
	*mock.Call
}

// GetBindings is a helper method to define mock.On call
//   - contractName string
func (_e *MockExtended_Expecter) GetBindings(contractName interface{}) *MockExtended_GetBindings_Call {
	return &MockExtended_GetBindings_Call{Call: _e.mock.On("GetBindings", contractName)}
}

func (_c *MockExtended_GetBindings_Call) Run(run func(contractName string)) *MockExtended_GetBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockExtended_GetBindings_Call) Return(_a0 []contractreader.ExtendedBoundContract) *MockExtended_GetBindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_GetBindings_Call) RunAndReturn(run func(string) []contractreader.ExtendedBoundContract) *MockExtended_GetBindings_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestValue provides a mock function with given fields: ctx, readIdentifier, confidenceLevel, params, returnVal
func (_m *MockExtended) GetLatestValue(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{}) error {
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

// MockExtended_GetLatestValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestValue'
type MockExtended_GetLatestValue_Call struct {
	*mock.Call
}

// GetLatestValue is a helper method to define mock.On call
//   - ctx context.Context
//   - readIdentifier string
//   - confidenceLevel primitives.ConfidenceLevel
//   - params interface{}
//   - returnVal interface{}
func (_e *MockExtended_Expecter) GetLatestValue(ctx interface{}, readIdentifier interface{}, confidenceLevel interface{}, params interface{}, returnVal interface{}) *MockExtended_GetLatestValue_Call {
	return &MockExtended_GetLatestValue_Call{Call: _e.mock.On("GetLatestValue", ctx, readIdentifier, confidenceLevel, params, returnVal)}
}

func (_c *MockExtended_GetLatestValue_Call) Run(run func(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params interface{}, returnVal interface{})) *MockExtended_GetLatestValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(primitives.ConfidenceLevel), args[3].(interface{}), args[4].(interface{}))
	})
	return _c
}

func (_c *MockExtended_GetLatestValue_Call) Return(_a0 error) *MockExtended_GetLatestValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_GetLatestValue_Call) RunAndReturn(run func(context.Context, string, primitives.ConfidenceLevel, interface{}, interface{}) error) *MockExtended_GetLatestValue_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *MockExtended) HealthReport() map[string]error {
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

// MockExtended_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type MockExtended_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *MockExtended_Expecter) HealthReport() *MockExtended_HealthReport_Call {
	return &MockExtended_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *MockExtended_HealthReport_Call) Run(run func()) *MockExtended_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtended_HealthReport_Call) Return(_a0 map[string]error) *MockExtended_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_HealthReport_Call) RunAndReturn(run func() map[string]error) *MockExtended_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockExtended) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockExtended_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockExtended_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockExtended_Expecter) Name() *MockExtended_Name_Call {
	return &MockExtended_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockExtended_Name_Call) Run(run func()) *MockExtended_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtended_Name_Call) Return(_a0 string) *MockExtended_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Name_Call) RunAndReturn(run func() string) *MockExtended_Name_Call {
	_c.Call.Return(run)
	return _c
}

// QueryKey provides a mock function with given fields: ctx, contract, filter, limitAndSort, sequenceDataType
func (_m *MockExtended) QueryKey(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{}) ([]types.Sequence, error) {
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

// MockExtended_QueryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryKey'
type MockExtended_QueryKey_Call struct {
	*mock.Call
}

// QueryKey is a helper method to define mock.On call
//   - ctx context.Context
//   - contract types.BoundContract
//   - filter query.KeyFilter
//   - limitAndSort query.LimitAndSort
//   - sequenceDataType interface{}
func (_e *MockExtended_Expecter) QueryKey(ctx interface{}, contract interface{}, filter interface{}, limitAndSort interface{}, sequenceDataType interface{}) *MockExtended_QueryKey_Call {
	return &MockExtended_QueryKey_Call{Call: _e.mock.On("QueryKey", ctx, contract, filter, limitAndSort, sequenceDataType)}
}

func (_c *MockExtended_QueryKey_Call) Run(run func(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType interface{})) *MockExtended_QueryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.BoundContract), args[2].(query.KeyFilter), args[3].(query.LimitAndSort), args[4].(interface{}))
	})
	return _c
}

func (_c *MockExtended_QueryKey_Call) Return(_a0 []types.Sequence, _a1 error) *MockExtended_QueryKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtended_QueryKey_Call) RunAndReturn(run func(context.Context, types.BoundContract, query.KeyFilter, query.LimitAndSort, interface{}) ([]types.Sequence, error)) *MockExtended_QueryKey_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *MockExtended) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExtended_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type MockExtended_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *MockExtended_Expecter) Ready() *MockExtended_Ready_Call {
	return &MockExtended_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *MockExtended_Ready_Call) Run(run func()) *MockExtended_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtended_Ready_Call) Return(_a0 error) *MockExtended_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Ready_Call) RunAndReturn(run func() error) *MockExtended_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *MockExtended) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExtended_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockExtended_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockExtended_Expecter) Start(_a0 interface{}) *MockExtended_Start_Call {
	return &MockExtended_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *MockExtended_Start_Call) Run(run func(_a0 context.Context)) *MockExtended_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockExtended_Start_Call) Return(_a0 error) *MockExtended_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Start_Call) RunAndReturn(run func(context.Context) error) *MockExtended_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Unbind provides a mock function with given fields: ctx, bindings
func (_m *MockExtended) Unbind(ctx context.Context, bindings []types.BoundContract) error {
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

// MockExtended_Unbind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unbind'
type MockExtended_Unbind_Call struct {
	*mock.Call
}

// Unbind is a helper method to define mock.On call
//   - ctx context.Context
//   - bindings []types.BoundContract
func (_e *MockExtended_Expecter) Unbind(ctx interface{}, bindings interface{}) *MockExtended_Unbind_Call {
	return &MockExtended_Unbind_Call{Call: _e.mock.On("Unbind", ctx, bindings)}
}

func (_c *MockExtended_Unbind_Call) Run(run func(ctx context.Context, bindings []types.BoundContract)) *MockExtended_Unbind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.BoundContract))
	})
	return _c
}

func (_c *MockExtended_Unbind_Call) Return(_a0 error) *MockExtended_Unbind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExtended_Unbind_Call) RunAndReturn(run func(context.Context, []types.BoundContract) error) *MockExtended_Unbind_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedContractReaderServer provides a mock function with given fields:
func (_m *MockExtended) mustEmbedUnimplementedContractReaderServer() {
	_m.Called()
}

// MockExtended_mustEmbedUnimplementedContractReaderServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedContractReaderServer'
type MockExtended_mustEmbedUnimplementedContractReaderServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedContractReaderServer is a helper method to define mock.On call
func (_e *MockExtended_Expecter) mustEmbedUnimplementedContractReaderServer() *MockExtended_mustEmbedUnimplementedContractReaderServer_Call {
	return &MockExtended_mustEmbedUnimplementedContractReaderServer_Call{Call: _e.mock.On("mustEmbedUnimplementedContractReaderServer")}
}

func (_c *MockExtended_mustEmbedUnimplementedContractReaderServer_Call) Run(run func()) *MockExtended_mustEmbedUnimplementedContractReaderServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExtended_mustEmbedUnimplementedContractReaderServer_Call) Return() *MockExtended_mustEmbedUnimplementedContractReaderServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockExtended_mustEmbedUnimplementedContractReaderServer_Call) RunAndReturn(run func()) *MockExtended_mustEmbedUnimplementedContractReaderServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExtended creates a new instance of MockExtended. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExtended(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExtended {
	mock := &MockExtended{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
