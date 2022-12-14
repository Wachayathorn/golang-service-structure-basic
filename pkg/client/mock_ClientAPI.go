// Code generated by mockery v2.15.0. DO NOT EDIT.

package client

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wachayathorn/golang-service-structure-basic/pkg/model"
)

// MockClientAPI is an autogenerated mock type for the ClientAPI type
type MockClientAPI struct {
	mock.Mock
}

type MockClientAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientAPI) EXPECT() *MockClientAPI_Expecter {
	return &MockClientAPI_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *MockClientAPI) Get() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientAPI_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockClientAPI_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockClientAPI_Expecter) Get() *MockClientAPI_Get_Call {
	return &MockClientAPI_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockClientAPI_Get_Call) Run(run func()) *MockClientAPI_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientAPI_Get_Call) Return(_a0 string, _a1 error) *MockClientAPI_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Post provides a mock function with given fields: req
func (_m *MockClientAPI) Post(req model.Post) (string, error) {
	ret := _m.Called(req)

	var r0 string
	if rf, ok := ret.Get(0).(func(model.Post) string); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Post) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientAPI_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type MockClientAPI_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - req model.Post
func (_e *MockClientAPI_Expecter) Post(req interface{}) *MockClientAPI_Post_Call {
	return &MockClientAPI_Post_Call{Call: _e.mock.On("Post", req)}
}

func (_c *MockClientAPI_Post_Call) Run(run func(req model.Post)) *MockClientAPI_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Post))
	})
	return _c
}

func (_c *MockClientAPI_Post_Call) Return(_a0 string, _a1 error) *MockClientAPI_Post_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReplaceUrl provides a mock function with given fields: url
func (_m *MockClientAPI) ReplaceUrl(url string) {
	_m.Called(url)
}

// MockClientAPI_ReplaceUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplaceUrl'
type MockClientAPI_ReplaceUrl_Call struct {
	*mock.Call
}

// ReplaceUrl is a helper method to define mock.On call
//   - url string
func (_e *MockClientAPI_Expecter) ReplaceUrl(url interface{}) *MockClientAPI_ReplaceUrl_Call {
	return &MockClientAPI_ReplaceUrl_Call{Call: _e.mock.On("ReplaceUrl", url)}
}

func (_c *MockClientAPI_ReplaceUrl_Call) Run(run func(url string)) *MockClientAPI_ReplaceUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClientAPI_ReplaceUrl_Call) Return() *MockClientAPI_ReplaceUrl_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockClientAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClientAPI creates a new instance of MockClientAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClientAPI(t mockConstructorTestingTNewMockClientAPI) *MockClientAPI {
	mock := &MockClientAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
