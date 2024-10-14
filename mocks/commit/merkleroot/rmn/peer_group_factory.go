// Code generated by mockery v2.43.2. DO NOT EDIT.

package rmn

import (
	commontypes "github.com/smartcontractkit/libocr/commontypes"
	mock "github.com/stretchr/testify/mock"

	networking "github.com/smartcontractkit/libocr/networking"

	types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// MockPeerGroupFactory is an autogenerated mock type for the PeerGroupFactory type
type MockPeerGroupFactory struct {
	mock.Mock
}

type MockPeerGroupFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPeerGroupFactory) EXPECT() *MockPeerGroupFactory_Expecter {
	return &MockPeerGroupFactory_Expecter{mock: &_m.Mock}
}

// NewPeerGroup provides a mock function with given fields: configDigest, peerIDs, bootstrappers
func (_m *MockPeerGroupFactory) NewPeerGroup(configDigest types.ConfigDigest, peerIDs []string, bootstrappers []commontypes.BootstrapperLocator) (networking.PeerGroup, error) {
	ret := _m.Called(configDigest, peerIDs, bootstrappers)

	if len(ret) == 0 {
		panic("no return value specified for NewPeerGroup")
	}

	var r0 networking.PeerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(types.ConfigDigest, []string, []commontypes.BootstrapperLocator) (networking.PeerGroup, error)); ok {
		return rf(configDigest, peerIDs, bootstrappers)
	}
	if rf, ok := ret.Get(0).(func(types.ConfigDigest, []string, []commontypes.BootstrapperLocator) networking.PeerGroup); ok {
		r0 = rf(configDigest, peerIDs, bootstrappers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(networking.PeerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(types.ConfigDigest, []string, []commontypes.BootstrapperLocator) error); ok {
		r1 = rf(configDigest, peerIDs, bootstrappers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPeerGroupFactory_NewPeerGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPeerGroup'
type MockPeerGroupFactory_NewPeerGroup_Call struct {
	*mock.Call
}

// NewPeerGroup is a helper method to define mock.On call
//   - configDigest types.ConfigDigest
//   - peerIDs []string
//   - bootstrappers []commontypes.BootstrapperLocator
func (_e *MockPeerGroupFactory_Expecter) NewPeerGroup(configDigest interface{}, peerIDs interface{}, bootstrappers interface{}) *MockPeerGroupFactory_NewPeerGroup_Call {
	return &MockPeerGroupFactory_NewPeerGroup_Call{Call: _e.mock.On("NewPeerGroup", configDigest, peerIDs, bootstrappers)}
}

func (_c *MockPeerGroupFactory_NewPeerGroup_Call) Run(run func(configDigest types.ConfigDigest, peerIDs []string, bootstrappers []commontypes.BootstrapperLocator)) *MockPeerGroupFactory_NewPeerGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.ConfigDigest), args[1].([]string), args[2].([]commontypes.BootstrapperLocator))
	})
	return _c
}

func (_c *MockPeerGroupFactory_NewPeerGroup_Call) Return(_a0 networking.PeerGroup, _a1 error) *MockPeerGroupFactory_NewPeerGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPeerGroupFactory_NewPeerGroup_Call) RunAndReturn(run func(types.ConfigDigest, []string, []commontypes.BootstrapperLocator) (networking.PeerGroup, error)) *MockPeerGroupFactory_NewPeerGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPeerGroupFactory creates a new instance of MockPeerGroupFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPeerGroupFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPeerGroupFactory {
	mock := &MockPeerGroupFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
