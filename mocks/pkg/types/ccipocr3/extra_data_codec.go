// Code generated by mockery v2.50.0. DO NOT EDIT.

package ccipocr3

import (
	ccipocr3 "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	mock "github.com/stretchr/testify/mock"
)

// MockExtraDataCodec is an autogenerated mock type for the ExtraDataCodec type
type MockExtraDataCodec struct {
	mock.Mock
}

type MockExtraDataCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExtraDataCodec) EXPECT() *MockExtraDataCodec_Expecter {
	return &MockExtraDataCodec_Expecter{mock: &_m.Mock}
}

// DecodeExtraArgs provides a mock function with given fields: extraArgs, sourceChainSelector
func (_m *MockExtraDataCodec) DecodeExtraArgs(extraArgs ccipocr3.Bytes, sourceChainSelector ccipocr3.ChainSelector) (map[string]interface{}, error) {
	ret := _m.Called(extraArgs, sourceChainSelector)

	if len(ret) == 0 {
		panic("no return value specified for DecodeExtraArgs")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) (map[string]interface{}, error)); ok {
		return rf(extraArgs, sourceChainSelector)
	}
	if rf, ok := ret.Get(0).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) map[string]interface{}); ok {
		r0 = rf(extraArgs, sourceChainSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) error); ok {
		r1 = rf(extraArgs, sourceChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtraDataCodec_DecodeExtraArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecodeExtraArgs'
type MockExtraDataCodec_DecodeExtraArgs_Call struct {
	*mock.Call
}

// DecodeExtraArgs is a helper method to define mock.On call
//   - extraArgs ccipocr3.Bytes
//   - sourceChainSelector ccipocr3.ChainSelector
func (_e *MockExtraDataCodec_Expecter) DecodeExtraArgs(extraArgs interface{}, sourceChainSelector interface{}) *MockExtraDataCodec_DecodeExtraArgs_Call {
	return &MockExtraDataCodec_DecodeExtraArgs_Call{Call: _e.mock.On("DecodeExtraArgs", extraArgs, sourceChainSelector)}
}

func (_c *MockExtraDataCodec_DecodeExtraArgs_Call) Run(run func(extraArgs ccipocr3.Bytes, sourceChainSelector ccipocr3.ChainSelector)) *MockExtraDataCodec_DecodeExtraArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ccipocr3.Bytes), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockExtraDataCodec_DecodeExtraArgs_Call) Return(_a0 map[string]interface{}, _a1 error) *MockExtraDataCodec_DecodeExtraArgs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtraDataCodec_DecodeExtraArgs_Call) RunAndReturn(run func(ccipocr3.Bytes, ccipocr3.ChainSelector) (map[string]interface{}, error)) *MockExtraDataCodec_DecodeExtraArgs_Call {
	_c.Call.Return(run)
	return _c
}

// DecodeTokenAmountDestExecData provides a mock function with given fields: destExecData, sourceChainSelector
func (_m *MockExtraDataCodec) DecodeTokenAmountDestExecData(destExecData ccipocr3.Bytes, sourceChainSelector ccipocr3.ChainSelector) (map[string]interface{}, error) {
	ret := _m.Called(destExecData, sourceChainSelector)

	if len(ret) == 0 {
		panic("no return value specified for DecodeTokenAmountDestExecData")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) (map[string]interface{}, error)); ok {
		return rf(destExecData, sourceChainSelector)
	}
	if rf, ok := ret.Get(0).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) map[string]interface{}); ok {
		r0 = rf(destExecData, sourceChainSelector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(ccipocr3.Bytes, ccipocr3.ChainSelector) error); ok {
		r1 = rf(destExecData, sourceChainSelector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExtraDataCodec_DecodeTokenAmountDestExecData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecodeTokenAmountDestExecData'
type MockExtraDataCodec_DecodeTokenAmountDestExecData_Call struct {
	*mock.Call
}

// DecodeTokenAmountDestExecData is a helper method to define mock.On call
//   - destExecData ccipocr3.Bytes
//   - sourceChainSelector ccipocr3.ChainSelector
func (_e *MockExtraDataCodec_Expecter) DecodeTokenAmountDestExecData(destExecData interface{}, sourceChainSelector interface{}) *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call {
	return &MockExtraDataCodec_DecodeTokenAmountDestExecData_Call{Call: _e.mock.On("DecodeTokenAmountDestExecData", destExecData, sourceChainSelector)}
}

func (_c *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call) Run(run func(destExecData ccipocr3.Bytes, sourceChainSelector ccipocr3.ChainSelector)) *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ccipocr3.Bytes), args[1].(ccipocr3.ChainSelector))
	})
	return _c
}

func (_c *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call) Return(_a0 map[string]interface{}, _a1 error) *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call) RunAndReturn(run func(ccipocr3.Bytes, ccipocr3.ChainSelector) (map[string]interface{}, error)) *MockExtraDataCodec_DecodeTokenAmountDestExecData_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExtraDataCodec creates a new instance of MockExtraDataCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExtraDataCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExtraDataCodec {
	mock := &MockExtraDataCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
