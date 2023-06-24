// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	ai "code-connect/pkg/ai"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// GPTClient is an autogenerated mock type for the GPTClient type
type GPTClient struct {
	mock.Mock
}

type GPTClient_Expecter struct {
	mock *mock.Mock
}

func (_m *GPTClient) EXPECT() *GPTClient_Expecter {
	return &GPTClient_Expecter{mock: &_m.Mock}
}

// AddPrompt provides a mock function with given fields: prompt
func (_m *GPTClient) AddPrompt(prompt string) {
	_m.Called(prompt)
}

// GPTClient_AddPrompt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPrompt'
type GPTClient_AddPrompt_Call struct {
	*mock.Call
}

// AddPrompt is a helper method to define mock.On call
//   - prompt string
func (_e *GPTClient_Expecter) AddPrompt(prompt interface{}) *GPTClient_AddPrompt_Call {
	return &GPTClient_AddPrompt_Call{Call: _e.mock.On("AddPrompt", prompt)}
}

func (_c *GPTClient_AddPrompt_Call) Run(run func(prompt string)) *GPTClient_AddPrompt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GPTClient_AddPrompt_Call) Return() *GPTClient_AddPrompt_Call {
	_c.Call.Return()
	return _c
}

func (_c *GPTClient_AddPrompt_Call) RunAndReturn(run func(string)) *GPTClient_AddPrompt_Call {
	_c.Call.Return(run)
	return _c
}

// ClearContext provides a mock function with given fields:
func (_m *GPTClient) ClearContext() {
	_m.Called()
}

// GPTClient_ClearContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearContext'
type GPTClient_ClearContext_Call struct {
	*mock.Call
}

// ClearContext is a helper method to define mock.On call
func (_e *GPTClient_Expecter) ClearContext() *GPTClient_ClearContext_Call {
	return &GPTClient_ClearContext_Call{Call: _e.mock.On("ClearContext")}
}

func (_c *GPTClient_ClearContext_Call) Run(run func()) *GPTClient_ClearContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GPTClient_ClearContext_Call) Return() *GPTClient_ClearContext_Call {
	_c.Call.Return()
	return _c
}

func (_c *GPTClient_ClearContext_Call) RunAndReturn(run func()) *GPTClient_ClearContext_Call {
	_c.Call.Return(run)
	return _c
}

// Complete provides a mock function with given fields: ctx
func (_m *GPTClient) Complete(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GPTClient_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type GPTClient_Complete_Call struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - ctx context.Context
func (_e *GPTClient_Expecter) Complete(ctx interface{}) *GPTClient_Complete_Call {
	return &GPTClient_Complete_Call{Call: _e.mock.On("Complete", ctx)}
}

func (_c *GPTClient_Complete_Call) Run(run func(ctx context.Context)) *GPTClient_Complete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *GPTClient_Complete_Call) Return(answer string, err error) *GPTClient_Complete_Call {
	_c.Call.Return(answer, err)
	return _c
}

func (_c *GPTClient_Complete_Call) RunAndReturn(run func(context.Context) (string, error)) *GPTClient_Complete_Call {
	_c.Call.Return(run)
	return _c
}

// NewContext provides a mock function with given fields:
func (_m *GPTClient) Clone() ai.GPTClient {
	ret := _m.Called()

	var r0 ai.GPTClient
	if rf, ok := ret.Get(0).(func() ai.GPTClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ai.GPTClient)
		}
	}

	return r0
}

// GPTClient_NewContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewContext'
type GPTClient_NewContext_Call struct {
	*mock.Call
}

// NewContext is a helper method to define mock.On call
func (_e *GPTClient_Expecter) NewContext() *GPTClient_NewContext_Call {
	return &GPTClient_NewContext_Call{Call: _e.mock.On("NewContext")}
}

func (_c *GPTClient_NewContext_Call) Run(run func()) *GPTClient_NewContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GPTClient_NewContext_Call) Return(_a0 ai.GPTClient) *GPTClient_NewContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GPTClient_NewContext_Call) RunAndReturn(run func() ai.GPTClient) *GPTClient_NewContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewGPTClient creates a new instance of GPTClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGPTClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *GPTClient {
	mock := &GPTClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
