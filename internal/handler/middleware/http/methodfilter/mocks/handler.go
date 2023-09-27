// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HandlerMock is an autogenerated mock type for the Handler type
type HandlerMock struct {
	mock.Mock
}

type HandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *HandlerMock) EXPECT() *HandlerMock_Expecter {
	return &HandlerMock_Expecter{mock: &_m.Mock}
}

// ServeHTTP provides a mock function with given fields: _a0, _a1
func (_m *HandlerMock) ServeHTTP(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// HandlerMock_ServeHTTP_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServeHTTP'
type HandlerMock_ServeHTTP_Call struct {
	*mock.Call
}

// ServeHTTP is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *HandlerMock_Expecter) ServeHTTP(_a0 interface{}, _a1 interface{}) *HandlerMock_ServeHTTP_Call {
	return &HandlerMock_ServeHTTP_Call{Call: _e.mock.On("ServeHTTP", _a0, _a1)}
}

func (_c *HandlerMock_ServeHTTP_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *HandlerMock_ServeHTTP_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *HandlerMock_ServeHTTP_Call) Return() *HandlerMock_ServeHTTP_Call {
	_c.Call.Return()
	return _c
}

func (_c *HandlerMock_ServeHTTP_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *HandlerMock_ServeHTTP_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHandlerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandlerMock creates a new instance of HandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandlerMock(t mockConstructorTestingTNewHandlerMock) *HandlerMock {
	mock := &HandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
