// Code generated by mockery v2.43.2. DO NOT EDIT.

package ccipocr3

import (
	context "context"

	ccipocr3 "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	mock "github.com/stretchr/testify/mock"
)

// MockExecutePluginCodec is an autogenerated mock type for the ExecutePluginCodec type
type MockExecutePluginCodec struct {
	mock.Mock
}

type MockExecutePluginCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExecutePluginCodec) EXPECT() *MockExecutePluginCodec_Expecter {
	return &MockExecutePluginCodec_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: _a0, _a1
func (_m *MockExecutePluginCodec) Decode(_a0 context.Context, _a1 []byte) (ccipocr3.ExecutePluginReport, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 ccipocr3.ExecutePluginReport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) (ccipocr3.ExecutePluginReport, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte) ccipocr3.ExecutePluginReport); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(ccipocr3.ExecutePluginReport)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExecutePluginCodec_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockExecutePluginCodec_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []byte
func (_e *MockExecutePluginCodec_Expecter) Decode(_a0 interface{}, _a1 interface{}) *MockExecutePluginCodec_Decode_Call {
	return &MockExecutePluginCodec_Decode_Call{Call: _e.mock.On("Decode", _a0, _a1)}
}

func (_c *MockExecutePluginCodec_Decode_Call) Run(run func(_a0 context.Context, _a1 []byte)) *MockExecutePluginCodec_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *MockExecutePluginCodec_Decode_Call) Return(_a0 ccipocr3.ExecutePluginReport, _a1 error) *MockExecutePluginCodec_Decode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExecutePluginCodec_Decode_Call) RunAndReturn(run func(context.Context, []byte) (ccipocr3.ExecutePluginReport, error)) *MockExecutePluginCodec_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// Encode provides a mock function with given fields: _a0, _a1
func (_m *MockExecutePluginCodec) Encode(_a0 context.Context, _a1 ccipocr3.ExecutePluginReport) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ccipocr3.ExecutePluginReport) ([]byte, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ccipocr3.ExecutePluginReport) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ccipocr3.ExecutePluginReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExecutePluginCodec_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type MockExecutePluginCodec_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ccipocr3.ExecutePluginReport
func (_e *MockExecutePluginCodec_Expecter) Encode(_a0 interface{}, _a1 interface{}) *MockExecutePluginCodec_Encode_Call {
	return &MockExecutePluginCodec_Encode_Call{Call: _e.mock.On("Encode", _a0, _a1)}
}

func (_c *MockExecutePluginCodec_Encode_Call) Run(run func(_a0 context.Context, _a1 ccipocr3.ExecutePluginReport)) *MockExecutePluginCodec_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ccipocr3.ExecutePluginReport))
	})
	return _c
}

func (_c *MockExecutePluginCodec_Encode_Call) Return(_a0 []byte, _a1 error) *MockExecutePluginCodec_Encode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExecutePluginCodec_Encode_Call) RunAndReturn(run func(context.Context, ccipocr3.ExecutePluginReport) ([]byte, error)) *MockExecutePluginCodec_Encode_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExecutePluginCodec creates a new instance of MockExecutePluginCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecutePluginCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExecutePluginCodec {
	mock := &MockExecutePluginCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
