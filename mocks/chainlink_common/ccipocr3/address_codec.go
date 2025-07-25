// Code generated by mockery v2.52.3. DO NOT EDIT.

package ccipocr3

import (
	ccipocr3 "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	mock "github.com/stretchr/testify/mock"
)

// MockAddressCodec is an autogenerated mock type for the AddressCodec type
type MockAddressCodec struct {
	mock.Mock
}

type MockAddressCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAddressCodec) EXPECT() *MockAddressCodec_Expecter {
	return &MockAddressCodec_Expecter{mock: &_m.Mock}
}

// AddressBytesToString provides a mock function with given fields: _a0, _a1
func (_m *MockAddressCodec) AddressBytesToString(_a0 ccipocr3.UnknownAddress, _a1 ccipocr3.ChainSelector) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddressBytesToString")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(ccipocr3.UnknownAddress, ccipocr3.ChainSelector) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(ccipocr3.UnknownAddress, ccipocr3.ChainSelector) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(ccipocr3.UnknownAddress, ccipocr3.ChainSelector) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAddressCodec_AddressBytesToString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddressBytesToString'
type MockAddressCodec_AddressBytesToString_Call struct {
	*mock.Call
}

// AddressBytesToString is a helper method to define mock.On call
//   - _a0 ccipocr3.UnknownAddress
//   - _a1 ccipocr3.ChainSelector
func (_e *MockAddressCodec_Expecter) AddressBytesToString(_a0 interface{}, _a1 interface{}) *MockAddressCodec_AddressBytesToString_Call {
	return &MockAddressCodec_AddressBytesToString_Call{Call: _e.mock.On("AddressBytesToString", _a0, _a1)}
}

func (_c *MockAddressCodec_AddressBytesToString_Call) Run(run func(_a0 ccipocr3.UnknownAddress, _a1 ccipocr3.ChainSelector)) *MockAddressCodec_AddressBytesToString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ccipocr3.UnknownAddress), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockAddressCodec_AddressBytesToString_Call) Return(_a0 string, _a1 error) *MockAddressCodec_AddressBytesToString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAddressCodec_AddressBytesToString_Call) RunAndReturn(run func(ccipocr3.UnknownAddress, ccipocr3.ChainSelector) (string, error)) *MockAddressCodec_AddressBytesToString_Call {
	_c.Call.Return(run)
	return _c
}

// AddressStringToBytes provides a mock function with given fields: _a0, _a1
func (_m *MockAddressCodec) AddressStringToBytes(_a0 string, _a1 ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddressStringToBytes")
	}

	var r0 ccipocr3.UnknownAddress
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, ccipocr3.ChainSelector) ccipocr3.UnknownAddress); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ccipocr3.UnknownAddress)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ccipocr3.ChainSelector) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAddressCodec_AddressStringToBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddressStringToBytes'
type MockAddressCodec_AddressStringToBytes_Call struct {
	*mock.Call
}

// AddressStringToBytes is a helper method to define mock.On call
//   - _a0 string
//   - _a1 ccipocr3.ChainSelector
func (_e *MockAddressCodec_Expecter) AddressStringToBytes(_a0 interface{}, _a1 interface{}) *MockAddressCodec_AddressStringToBytes_Call {
	return &MockAddressCodec_AddressStringToBytes_Call{Call: _e.mock.On("AddressStringToBytes", _a0, _a1)}
}

func (_c *MockAddressCodec_AddressStringToBytes_Call) Run(run func(_a0 string, _a1 ccipocr3.ChainSelector)) *MockAddressCodec_AddressStringToBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockAddressCodec_AddressStringToBytes_Call) Return(_a0 ccipocr3.UnknownAddress, _a1 error) *MockAddressCodec_AddressStringToBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAddressCodec_AddressStringToBytes_Call) RunAndReturn(run func(string, ccipocr3.ChainSelector) (ccipocr3.UnknownAddress, error)) *MockAddressCodec_AddressStringToBytes_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAddressCodec creates a new instance of MockAddressCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAddressCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAddressCodec {
	mock := &MockAddressCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
