// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IJwtManager is an autogenerated mock type for the IJwtManager type
type IJwtManager struct {
	mock.Mock
}

// Create provides a mock function with given fields:
func (_m *IJwtManager) Create() (string, error) {
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

// Verify provides a mock function with given fields: token
func (_m *IJwtManager) Verify(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIJwtManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewIJwtManager creates a new instance of IJwtManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIJwtManager(t mockConstructorTestingTNewIJwtManager) *IJwtManager {
	mock := &IJwtManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
