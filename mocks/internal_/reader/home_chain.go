// Code generated by mockery v2.43.2. DO NOT EDIT.

package reader

import (
	context "context"

	ccipocr3 "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	mapset "github.com/deckarep/golang-set/v2"

	mock "github.com/stretchr/testify/mock"

	reader "github.com/smartcontractkit/chainlink-ccip/internal/reader"

	types "github.com/smartcontractkit/libocr/ragep2p/types"
)

// MockHomeChain is an autogenerated mock type for the HomeChain type
type MockHomeChain struct {
	mock.Mock
}

type MockHomeChain_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHomeChain) EXPECT() *MockHomeChain_Expecter {
	return &MockHomeChain_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockHomeChain) Close() error {
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

// MockHomeChain_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockHomeChain_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) Close() *MockHomeChain_Close_Call {
	return &MockHomeChain_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockHomeChain_Close_Call) Run(run func()) *MockHomeChain_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_Close_Call) Return(_a0 error) *MockHomeChain_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHomeChain_Close_Call) RunAndReturn(run func() error) *MockHomeChain_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllChainConfigs provides a mock function with given fields:
func (_m *MockHomeChain) GetAllChainConfigs() (map[ccipocr3.ChainSelector]reader.ChainConfig, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllChainConfigs")
	}

	var r0 map[ccipocr3.ChainSelector]reader.ChainConfig
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[ccipocr3.ChainSelector]reader.ChainConfig, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[ccipocr3.ChainSelector]reader.ChainConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ccipocr3.ChainSelector]reader.ChainConfig)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetAllChainConfigs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllChainConfigs'
type MockHomeChain_GetAllChainConfigs_Call struct {
	*mock.Call
}

// GetAllChainConfigs is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) GetAllChainConfigs() *MockHomeChain_GetAllChainConfigs_Call {
	return &MockHomeChain_GetAllChainConfigs_Call{Call: _e.mock.On("GetAllChainConfigs")}
}

func (_c *MockHomeChain_GetAllChainConfigs_Call) Run(run func()) *MockHomeChain_GetAllChainConfigs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_GetAllChainConfigs_Call) Return(_a0 map[ccipocr3.ChainSelector]reader.ChainConfig, _a1 error) *MockHomeChain_GetAllChainConfigs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetAllChainConfigs_Call) RunAndReturn(run func() (map[ccipocr3.ChainSelector]reader.ChainConfig, error)) *MockHomeChain_GetAllChainConfigs_Call {
	_c.Call.Return(run)
	return _c
}

// GetChainConfig provides a mock function with given fields: chainSelector
func (_m *MockHomeChain) GetChainConfig(chainSelector ccipocr3.ChainSelector) (reader.ChainConfig, error) {
	ret := _m.Called(chainSelector)

	if len(ret) == 0 {
		panic("no return value specified for GetChainConfig")
	}

	var r0 reader.ChainConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(ccipocr3.ChainSelector) (reader.ChainConfig, error)); ok {
		return rf(chainSelector)
	}
	if rf, ok := ret.Get(0).(func(ccipocr3.ChainSelector) reader.ChainConfig); ok {
		r0 = rf(chainSelector)
	} else {
		r0 = ret.Get(0).(reader.ChainConfig)
	}

	if rf, ok := ret.Get(1).(func(ccipocr3.ChainSelector) error); ok {
		r1 = rf(chainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetChainConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChainConfig'
type MockHomeChain_GetChainConfig_Call struct {
	*mock.Call
}

// GetChainConfig is a helper method to define mock.On call
//   - chainSelector ccipocr3.ChainSelector
func (_e *MockHomeChain_Expecter) GetChainConfig(chainSelector interface{}) *MockHomeChain_GetChainConfig_Call {
	return &MockHomeChain_GetChainConfig_Call{Call: _e.mock.On("GetChainConfig", chainSelector)}
}

func (_c *MockHomeChain_GetChainConfig_Call) Run(run func(chainSelector ccipocr3.ChainSelector)) *MockHomeChain_GetChainConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockHomeChain_GetChainConfig_Call) Return(_a0 reader.ChainConfig, _a1 error) *MockHomeChain_GetChainConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetChainConfig_Call) RunAndReturn(run func(ccipocr3.ChainSelector) (reader.ChainConfig, error)) *MockHomeChain_GetChainConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetFChain provides a mock function with given fields:
func (_m *MockHomeChain) GetFChain() (map[ccipocr3.ChainSelector]int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFChain")
	}

	var r0 map[ccipocr3.ChainSelector]int
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[ccipocr3.ChainSelector]int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[ccipocr3.ChainSelector]int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[ccipocr3.ChainSelector]int)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetFChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFChain'
type MockHomeChain_GetFChain_Call struct {
	*mock.Call
}

// GetFChain is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) GetFChain() *MockHomeChain_GetFChain_Call {
	return &MockHomeChain_GetFChain_Call{Call: _e.mock.On("GetFChain")}
}

func (_c *MockHomeChain_GetFChain_Call) Run(run func()) *MockHomeChain_GetFChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_GetFChain_Call) Return(_a0 map[ccipocr3.ChainSelector]int, _a1 error) *MockHomeChain_GetFChain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetFChain_Call) RunAndReturn(run func() (map[ccipocr3.ChainSelector]int, error)) *MockHomeChain_GetFChain_Call {
	_c.Call.Return(run)
	return _c
}

// GetKnownCCIPChains provides a mock function with given fields:
func (_m *MockHomeChain) GetKnownCCIPChains() (mapset.Set[ccipocr3.ChainSelector], error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetKnownCCIPChains")
	}

	var r0 mapset.Set[ccipocr3.ChainSelector]
	var r1 error
	if rf, ok := ret.Get(0).(func() (mapset.Set[ccipocr3.ChainSelector], error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() mapset.Set[ccipocr3.ChainSelector]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mapset.Set[ccipocr3.ChainSelector])
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetKnownCCIPChains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKnownCCIPChains'
type MockHomeChain_GetKnownCCIPChains_Call struct {
	*mock.Call
}

// GetKnownCCIPChains is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) GetKnownCCIPChains() *MockHomeChain_GetKnownCCIPChains_Call {
	return &MockHomeChain_GetKnownCCIPChains_Call{Call: _e.mock.On("GetKnownCCIPChains")}
}

func (_c *MockHomeChain_GetKnownCCIPChains_Call) Run(run func()) *MockHomeChain_GetKnownCCIPChains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_GetKnownCCIPChains_Call) Return(_a0 mapset.Set[ccipocr3.ChainSelector], _a1 error) *MockHomeChain_GetKnownCCIPChains_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetKnownCCIPChains_Call) RunAndReturn(run func() (mapset.Set[ccipocr3.ChainSelector], error)) *MockHomeChain_GetKnownCCIPChains_Call {
	_c.Call.Return(run)
	return _c
}

// GetOCRConfigs provides a mock function with given fields: ctx, donID, pluginType
func (_m *MockHomeChain) GetOCRConfigs(ctx context.Context, donID uint32, pluginType uint8) (reader.ActiveAndCandidate, error) {
	ret := _m.Called(ctx, donID, pluginType)

	if len(ret) == 0 {
		panic("no return value specified for GetOCRConfigs")
	}

	var r0 reader.ActiveAndCandidate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8) (reader.ActiveAndCandidate, error)); ok {
		return rf(ctx, donID, pluginType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8) reader.ActiveAndCandidate); ok {
		r0 = rf(ctx, donID, pluginType)
	} else {
		r0 = ret.Get(0).(reader.ActiveAndCandidate)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8) error); ok {
		r1 = rf(ctx, donID, pluginType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetOCRConfigs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOCRConfigs'
type MockHomeChain_GetOCRConfigs_Call struct {
	*mock.Call
}

// GetOCRConfigs is a helper method to define mock.On call
//   - ctx context.Context
//   - donID uint32
//   - pluginType uint8
func (_e *MockHomeChain_Expecter) GetOCRConfigs(ctx interface{}, donID interface{}, pluginType interface{}) *MockHomeChain_GetOCRConfigs_Call {
	return &MockHomeChain_GetOCRConfigs_Call{Call: _e.mock.On("GetOCRConfigs", ctx, donID, pluginType)}
}

func (_c *MockHomeChain_GetOCRConfigs_Call) Run(run func(ctx context.Context, donID uint32, pluginType uint8)) *MockHomeChain_GetOCRConfigs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8))
	})
	return _c
}

func (_c *MockHomeChain_GetOCRConfigs_Call) Return(_a0 reader.ActiveAndCandidate, _a1 error) *MockHomeChain_GetOCRConfigs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetOCRConfigs_Call) RunAndReturn(run func(context.Context, uint32, uint8) (reader.ActiveAndCandidate, error)) *MockHomeChain_GetOCRConfigs_Call {
	_c.Call.Return(run)
	return _c
}

// GetSupportedChainsForPeer provides a mock function with given fields: id
func (_m *MockHomeChain) GetSupportedChainsForPeer(id types.PeerID) (mapset.Set[ccipocr3.ChainSelector], error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetSupportedChainsForPeer")
	}

	var r0 mapset.Set[ccipocr3.ChainSelector]
	var r1 error
	if rf, ok := ret.Get(0).(func(types.PeerID) (mapset.Set[ccipocr3.ChainSelector], error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(types.PeerID) mapset.Set[ccipocr3.ChainSelector]); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mapset.Set[ccipocr3.ChainSelector])
		}
	}

	if rf, ok := ret.Get(1).(func(types.PeerID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHomeChain_GetSupportedChainsForPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSupportedChainsForPeer'
type MockHomeChain_GetSupportedChainsForPeer_Call struct {
	*mock.Call
}

// GetSupportedChainsForPeer is a helper method to define mock.On call
//   - id types.PeerID
func (_e *MockHomeChain_Expecter) GetSupportedChainsForPeer(id interface{}) *MockHomeChain_GetSupportedChainsForPeer_Call {
	return &MockHomeChain_GetSupportedChainsForPeer_Call{Call: _e.mock.On("GetSupportedChainsForPeer", id)}
}

func (_c *MockHomeChain_GetSupportedChainsForPeer_Call) Run(run func(id types.PeerID)) *MockHomeChain_GetSupportedChainsForPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.PeerID))
	})
	return _c
}

func (_c *MockHomeChain_GetSupportedChainsForPeer_Call) Return(_a0 mapset.Set[ccipocr3.ChainSelector], _a1 error) *MockHomeChain_GetSupportedChainsForPeer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHomeChain_GetSupportedChainsForPeer_Call) RunAndReturn(run func(types.PeerID) (mapset.Set[ccipocr3.ChainSelector], error)) *MockHomeChain_GetSupportedChainsForPeer_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *MockHomeChain) HealthReport() map[string]error {
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

// MockHomeChain_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type MockHomeChain_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) HealthReport() *MockHomeChain_HealthReport_Call {
	return &MockHomeChain_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *MockHomeChain_HealthReport_Call) Run(run func()) *MockHomeChain_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_HealthReport_Call) Return(_a0 map[string]error) *MockHomeChain_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHomeChain_HealthReport_Call) RunAndReturn(run func() map[string]error) *MockHomeChain_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockHomeChain) Name() string {
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

// MockHomeChain_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockHomeChain_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) Name() *MockHomeChain_Name_Call {
	return &MockHomeChain_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockHomeChain_Name_Call) Run(run func()) *MockHomeChain_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_Name_Call) Return(_a0 string) *MockHomeChain_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHomeChain_Name_Call) RunAndReturn(run func() string) *MockHomeChain_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *MockHomeChain) Ready() error {
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

// MockHomeChain_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type MockHomeChain_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *MockHomeChain_Expecter) Ready() *MockHomeChain_Ready_Call {
	return &MockHomeChain_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *MockHomeChain_Ready_Call) Run(run func()) *MockHomeChain_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHomeChain_Ready_Call) Return(_a0 error) *MockHomeChain_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHomeChain_Ready_Call) RunAndReturn(run func() error) *MockHomeChain_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *MockHomeChain) Start(_a0 context.Context) error {
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

// MockHomeChain_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockHomeChain_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockHomeChain_Expecter) Start(_a0 interface{}) *MockHomeChain_Start_Call {
	return &MockHomeChain_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *MockHomeChain_Start_Call) Run(run func(_a0 context.Context)) *MockHomeChain_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockHomeChain_Start_Call) Return(_a0 error) *MockHomeChain_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHomeChain_Start_Call) RunAndReturn(run func(context.Context) error) *MockHomeChain_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHomeChain creates a new instance of MockHomeChain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHomeChain(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHomeChain {
	mock := &MockHomeChain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
