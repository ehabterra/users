// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Db is an autogenerated mock type for the Db type
type Db struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *Db) Delete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Load provides a mock function with given fields: _a0, _a1
func (_m *Db) Load(_a0 string, _a1 string) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadAll provides a mock function with given fields: bucket
func (_m *Db) LoadAll(bucket string) (interface{}, error) {
	ret := _m.Called(bucket)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(bucket)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bucket)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewID provides a mock function with given fields: _a0
func (_m *Db) NewID(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0, _a1, _a2
func (_m *Db) Save(_a0 string, _a1 string, _a2 interface{}) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
