// Code generated by mockery v2.43.2. DO NOT EDIT.

package rmn

import (
	context "context"

	rmn "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	mock "github.com/stretchr/testify/mock"

	rmnpb "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// ComputeReportSignatures provides a mock function with given fields: ctx, destChain, requestedUpdates
func (_m *MockClient) ComputeReportSignatures(ctx context.Context, destChain *rmnpb.LaneDest, requestedUpdates []rmnpb.FixedDestLaneUpdateRequest) (*rmn.ReportSignatures, error) {
	ret := _m.Called(ctx, destChain, requestedUpdates)

	if len(ret) == 0 {
		panic("no return value specified for ComputeReportSignatures")
	}

	var r0 *rmn.ReportSignatures
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rmnpb.LaneDest, []rmnpb.FixedDestLaneUpdateRequest) (*rmn.ReportSignatures, error)); ok {
		return rf(ctx, destChain, requestedUpdates)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rmnpb.LaneDest, []rmnpb.FixedDestLaneUpdateRequest) *rmn.ReportSignatures); ok {
		r0 = rf(ctx, destChain, requestedUpdates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rmn.ReportSignatures)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rmnpb.LaneDest, []rmnpb.FixedDestLaneUpdateRequest) error); ok {
		r1 = rf(ctx, destChain, requestedUpdates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ComputeReportSignatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComputeReportSignatures'
type MockClient_ComputeReportSignatures_Call struct {
	*mock.Call
}

// ComputeReportSignatures is a helper method to define mock.On call
//   - ctx context.Context
//   - destChain *rmnpb.LaneDest
//   - requestedUpdates []rmnpb.FixedDestLaneUpdateRequest
func (_e *MockClient_Expecter) ComputeReportSignatures(ctx interface{}, destChain interface{}, requestedUpdates interface{}) *MockClient_ComputeReportSignatures_Call {
	return &MockClient_ComputeReportSignatures_Call{Call: _e.mock.On("ComputeReportSignatures", ctx, destChain, requestedUpdates)}
}

func (_c *MockClient_ComputeReportSignatures_Call) Run(run func(ctx context.Context, destChain *rmnpb.LaneDest, requestedUpdates []rmnpb.FixedDestLaneUpdateRequest)) *MockClient_ComputeReportSignatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rmnpb.LaneDest), args[2].([]rmnpb.FixedDestLaneUpdateRequest))
	})
	return _c
}

func (_c *MockClient_ComputeReportSignatures_Call) Return(_a0 *rmn.ReportSignatures, _a1 error) *MockClient_ComputeReportSignatures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ComputeReportSignatures_Call) RunAndReturn(run func(context.Context, *rmnpb.LaneDest, []rmnpb.FixedDestLaneUpdateRequest) (*rmn.ReportSignatures, error)) *MockClient_ComputeReportSignatures_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}