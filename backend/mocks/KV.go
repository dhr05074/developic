// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KV is an autogenerated mock type for the KV type
type KV struct {
	mock.Mock
}

type KV_Expecter struct {
	mock *mock.Mock
}

func (_m *KV) EXPECT() *KV_Expecter {
	return &KV_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key
func (_m *KV) Get(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KV_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type KV_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *KV_Expecter) Get(ctx interface{}, key interface{}) *KV_Get_Call {
	return &KV_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *KV_Get_Call) Run(run func(ctx context.Context, key string)) *KV_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *KV_Get_Call) Return(_a0 string, _a1 error) *KV_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KV_Get_Call) RunAndReturn(run func(context.Context, string) (string, error)) *KV_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewKV creates a new instance of KV. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKV(t interface {
	mock.TestingT
	Cleanup(func())
}) *KV {
	mock := &KV{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
