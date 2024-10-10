// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	perpetualtypes "github.com/elys-network/elys/x/perpetual/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PerpetualKeeper is an autogenerated mock type for the PerpetualKeeper type
type PerpetualKeeper struct {
	mock.Mock
}

type PerpetualKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *PerpetualKeeper) EXPECT() *PerpetualKeeper_Expecter {
	return &PerpetualKeeper_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields: ctx, msg
func (_m *PerpetualKeeper) Close(ctx types.Context, msg *perpetualtypes.MsgClose) (*perpetualtypes.MsgCloseResponse, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 *perpetualtypes.MsgCloseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MsgClose) (*perpetualtypes.MsgCloseResponse, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MsgClose) *perpetualtypes.MsgCloseResponse); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualtypes.MsgCloseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *perpetualtypes.MsgClose) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerpetualKeeper_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type PerpetualKeeper_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx types.Context
//   - msg *perpetualtypes.MsgClose
func (_e *PerpetualKeeper_Expecter) Close(ctx interface{}, msg interface{}) *PerpetualKeeper_Close_Call {
	return &PerpetualKeeper_Close_Call{Call: _e.mock.On("Close", ctx, msg)}
}

func (_c *PerpetualKeeper_Close_Call) Run(run func(ctx types.Context, msg *perpetualtypes.MsgClose)) *PerpetualKeeper_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.MsgClose))
	})
	return _c
}

func (_c *PerpetualKeeper_Close_Call) Return(_a0 *perpetualtypes.MsgCloseResponse, _a1 error) *PerpetualKeeper_Close_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PerpetualKeeper_Close_Call) RunAndReturn(run func(types.Context, *perpetualtypes.MsgClose) (*perpetualtypes.MsgCloseResponse, error)) *PerpetualKeeper_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetMTP provides a mock function with given fields: ctx, mtpAddress, id
func (_m *PerpetualKeeper) GetMTP(ctx types.Context, mtpAddress types.AccAddress, id uint64) (perpetualtypes.MTP, error) {
	ret := _m.Called(ctx, mtpAddress, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMTP")
	}

	var r0 perpetualtypes.MTP
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, uint64) (perpetualtypes.MTP, error)); ok {
		return rf(ctx, mtpAddress, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, uint64) perpetualtypes.MTP); ok {
		r0 = rf(ctx, mtpAddress, id)
	} else {
		r0 = ret.Get(0).(perpetualtypes.MTP)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.AccAddress, uint64) error); ok {
		r1 = rf(ctx, mtpAddress, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerpetualKeeper_GetMTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMTP'
type PerpetualKeeper_GetMTP_Call struct {
	*mock.Call
}

// GetMTP is a helper method to define mock.On call
//   - ctx types.Context
//   - mtpAddress types.AccAddress
//   - id uint64
func (_e *PerpetualKeeper_Expecter) GetMTP(ctx interface{}, mtpAddress interface{}, id interface{}) *PerpetualKeeper_GetMTP_Call {
	return &PerpetualKeeper_GetMTP_Call{Call: _e.mock.On("GetMTP", ctx, mtpAddress, id)}
}

func (_c *PerpetualKeeper_GetMTP_Call) Run(run func(ctx types.Context, mtpAddress types.AccAddress, id uint64)) *PerpetualKeeper_GetMTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress), args[2].(uint64))
	})
	return _c
}

func (_c *PerpetualKeeper_GetMTP_Call) Return(_a0 perpetualtypes.MTP, _a1 error) *PerpetualKeeper_GetMTP_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PerpetualKeeper_GetMTP_Call) RunAndReturn(run func(types.Context, types.AccAddress, uint64) (perpetualtypes.MTP, error)) *PerpetualKeeper_GetMTP_Call {
	_c.Call.Return(run)
	return _c
}

// GetPool provides a mock function with given fields: ctx, poolId
func (_m *PerpetualKeeper) GetPool(ctx types.Context, poolId uint64) (perpetualtypes.Pool, bool) {
	ret := _m.Called(ctx, poolId)

	if len(ret) == 0 {
		panic("no return value specified for GetPool")
	}

	var r0 perpetualtypes.Pool
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (perpetualtypes.Pool, bool)); ok {
		return rf(ctx, poolId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) perpetualtypes.Pool); ok {
		r0 = rf(ctx, poolId)
	} else {
		r0 = ret.Get(0).(perpetualtypes.Pool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) bool); ok {
		r1 = rf(ctx, poolId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// PerpetualKeeper_GetPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPool'
type PerpetualKeeper_GetPool_Call struct {
	*mock.Call
}

// GetPool is a helper method to define mock.On call
//   - ctx types.Context
//   - poolId uint64
func (_e *PerpetualKeeper_Expecter) GetPool(ctx interface{}, poolId interface{}) *PerpetualKeeper_GetPool_Call {
	return &PerpetualKeeper_GetPool_Call{Call: _e.mock.On("GetPool", ctx, poolId)}
}

func (_c *PerpetualKeeper_GetPool_Call) Run(run func(ctx types.Context, poolId uint64)) *PerpetualKeeper_GetPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(uint64))
	})
	return _c
}

func (_c *PerpetualKeeper_GetPool_Call) Return(val perpetualtypes.Pool, found bool) *PerpetualKeeper_GetPool_Call {
	_c.Call.Return(val, found)
	return _c
}

func (_c *PerpetualKeeper_GetPool_Call) RunAndReturn(run func(types.Context, uint64) (perpetualtypes.Pool, bool)) *PerpetualKeeper_GetPool_Call {
	_c.Call.Return(run)
	return _c
}

// HandleCloseEstimation provides a mock function with given fields: ctx, req
func (_m *PerpetualKeeper) HandleCloseEstimation(ctx types.Context, req *perpetualtypes.QueryCloseEstimationRequest) (*perpetualtypes.QueryCloseEstimationResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleCloseEstimation")
	}

	var r0 *perpetualtypes.QueryCloseEstimationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.QueryCloseEstimationRequest) (*perpetualtypes.QueryCloseEstimationResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.QueryCloseEstimationRequest) *perpetualtypes.QueryCloseEstimationResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualtypes.QueryCloseEstimationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *perpetualtypes.QueryCloseEstimationRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerpetualKeeper_HandleCloseEstimation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleCloseEstimation'
type PerpetualKeeper_HandleCloseEstimation_Call struct {
	*mock.Call
}

// HandleCloseEstimation is a helper method to define mock.On call
//   - ctx types.Context
//   - req *perpetualtypes.QueryCloseEstimationRequest
func (_e *PerpetualKeeper_Expecter) HandleCloseEstimation(ctx interface{}, req interface{}) *PerpetualKeeper_HandleCloseEstimation_Call {
	return &PerpetualKeeper_HandleCloseEstimation_Call{Call: _e.mock.On("HandleCloseEstimation", ctx, req)}
}

func (_c *PerpetualKeeper_HandleCloseEstimation_Call) Run(run func(ctx types.Context, req *perpetualtypes.QueryCloseEstimationRequest)) *PerpetualKeeper_HandleCloseEstimation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.QueryCloseEstimationRequest))
	})
	return _c
}

func (_c *PerpetualKeeper_HandleCloseEstimation_Call) Return(res *perpetualtypes.QueryCloseEstimationResponse, err error) *PerpetualKeeper_HandleCloseEstimation_Call {
	_c.Call.Return(res, err)
	return _c
}

func (_c *PerpetualKeeper_HandleCloseEstimation_Call) RunAndReturn(run func(types.Context, *perpetualtypes.QueryCloseEstimationRequest) (*perpetualtypes.QueryCloseEstimationResponse, error)) *PerpetualKeeper_HandleCloseEstimation_Call {
	_c.Call.Return(run)
	return _c
}

// HandleOpenEstimation provides a mock function with given fields: ctx, req
func (_m *PerpetualKeeper) HandleOpenEstimation(ctx types.Context, req *perpetualtypes.QueryOpenEstimationRequest) (*perpetualtypes.QueryOpenEstimationResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for HandleOpenEstimation")
	}

	var r0 *perpetualtypes.QueryOpenEstimationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.QueryOpenEstimationRequest) (*perpetualtypes.QueryOpenEstimationResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.QueryOpenEstimationRequest) *perpetualtypes.QueryOpenEstimationResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualtypes.QueryOpenEstimationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *perpetualtypes.QueryOpenEstimationRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerpetualKeeper_HandleOpenEstimation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleOpenEstimation'
type PerpetualKeeper_HandleOpenEstimation_Call struct {
	*mock.Call
}

// HandleOpenEstimation is a helper method to define mock.On call
//   - ctx types.Context
//   - req *perpetualtypes.QueryOpenEstimationRequest
func (_e *PerpetualKeeper_Expecter) HandleOpenEstimation(ctx interface{}, req interface{}) *PerpetualKeeper_HandleOpenEstimation_Call {
	return &PerpetualKeeper_HandleOpenEstimation_Call{Call: _e.mock.On("HandleOpenEstimation", ctx, req)}
}

func (_c *PerpetualKeeper_HandleOpenEstimation_Call) Run(run func(ctx types.Context, req *perpetualtypes.QueryOpenEstimationRequest)) *PerpetualKeeper_HandleOpenEstimation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.QueryOpenEstimationRequest))
	})
	return _c
}

func (_c *PerpetualKeeper_HandleOpenEstimation_Call) Return(_a0 *perpetualtypes.QueryOpenEstimationResponse, _a1 error) *PerpetualKeeper_HandleOpenEstimation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PerpetualKeeper_HandleOpenEstimation_Call) RunAndReturn(run func(types.Context, *perpetualtypes.QueryOpenEstimationRequest) (*perpetualtypes.QueryOpenEstimationResponse, error)) *PerpetualKeeper_HandleOpenEstimation_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: ctx, msg, isBroker
func (_m *PerpetualKeeper) Open(ctx types.Context, msg *perpetualtypes.MsgOpen, isBroker bool) (*perpetualtypes.MsgOpenResponse, error) {
	ret := _m.Called(ctx, msg, isBroker)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 *perpetualtypes.MsgOpenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MsgOpen, bool) (*perpetualtypes.MsgOpenResponse, error)); ok {
		return rf(ctx, msg, isBroker)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *perpetualtypes.MsgOpen, bool) *perpetualtypes.MsgOpenResponse); ok {
		r0 = rf(ctx, msg, isBroker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualtypes.MsgOpenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *perpetualtypes.MsgOpen, bool) error); ok {
		r1 = rf(ctx, msg, isBroker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerpetualKeeper_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type PerpetualKeeper_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - ctx types.Context
//   - msg *perpetualtypes.MsgOpen
//   - isBroker bool
func (_e *PerpetualKeeper_Expecter) Open(ctx interface{}, msg interface{}, isBroker interface{}) *PerpetualKeeper_Open_Call {
	return &PerpetualKeeper_Open_Call{Call: _e.mock.On("Open", ctx, msg, isBroker)}
}

func (_c *PerpetualKeeper_Open_Call) Run(run func(ctx types.Context, msg *perpetualtypes.MsgOpen, isBroker bool)) *PerpetualKeeper_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(*perpetualtypes.MsgOpen), args[2].(bool))
	})
	return _c
}

func (_c *PerpetualKeeper_Open_Call) Return(_a0 *perpetualtypes.MsgOpenResponse, _a1 error) *PerpetualKeeper_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PerpetualKeeper_Open_Call) RunAndReturn(run func(types.Context, *perpetualtypes.MsgOpen, bool) (*perpetualtypes.MsgOpenResponse, error)) *PerpetualKeeper_Open_Call {
	_c.Call.Return(run)
	return _c
}

// NewPerpetualKeeper creates a new instance of PerpetualKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPerpetualKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *PerpetualKeeper {
	mock := &PerpetualKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}