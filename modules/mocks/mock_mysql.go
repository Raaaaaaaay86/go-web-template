// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// IMySQLGorm is an autogenerated mock type for the IMySQLGorm type
type IMySQLGorm struct {
	mock.Mock
}

// CreateMySQLConnection provides a mock function with given fields:
func (_m *IMySQLGorm) CreateMySQLConnection() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *IMySQLGorm) Get() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewIMySQLGorm interface {
	mock.TestingT
	Cleanup(func())
}

// NewIMySQLGorm creates a new instance of IMySQLGorm. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIMySQLGorm(t mockConstructorTestingTNewIMySQLGorm) *IMySQLGorm {
	mock := &IMySQLGorm{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
