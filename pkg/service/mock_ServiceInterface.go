// Code generated by mockery v2.15.0. DO NOT EDIT.

package service

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/wachayathorn/golang-service-structure-basic/pkg/model"
)

// MockServiceInterface is an autogenerated mock type for the ServiceInterface type
type MockServiceInterface struct {
	mock.Mock
}

type MockServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServiceInterface) EXPECT() *MockServiceInterface_Expecter {
	return &MockServiceInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *MockServiceInterface) Get() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockServiceInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockServiceInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockServiceInterface_Expecter) Get() *MockServiceInterface_Get_Call {
	return &MockServiceInterface_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *MockServiceInterface_Get_Call) Run(run func()) *MockServiceInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockServiceInterface_Get_Call) Return(_a0 string) *MockServiceInterface_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

// Post provides a mock function with given fields: req
func (_m *MockServiceInterface) Post(req model.Post) (string, error) {
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

// MockServiceInterface_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type MockServiceInterface_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - req model.Post
func (_e *MockServiceInterface_Expecter) Post(req interface{}) *MockServiceInterface_Post_Call {
	return &MockServiceInterface_Post_Call{Call: _e.mock.On("Post", req)}
}

func (_c *MockServiceInterface_Post_Call) Run(run func(req model.Post)) *MockServiceInterface_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Post))
	})
	return _c
}

func (_c *MockServiceInterface_Post_Call) Return(_a0 string, _a1 error) *MockServiceInterface_Post_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockServiceInterface creates a new instance of MockServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockServiceInterface(t mockConstructorTestingTNewMockServiceInterface) *MockServiceInterface {
	mock := &MockServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
