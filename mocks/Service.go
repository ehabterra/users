// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	users "users/gen/users"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Activate provides a mock function with given fields: _a0, _a1
func (_m *Service) Activate(_a0 context.Context, _a1 []string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *Service) Add(_a0 context.Context, _a1 *users.User) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *users.User) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *Service) List(_a0 context.Context, _a1 *users.ListPayload) (users.StoredUserCollection, string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 users.StoredUserCollection
	if rf, ok := ret.Get(0).(func(context.Context, *users.ListPayload) users.StoredUserCollection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(users.StoredUserCollection)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *users.ListPayload) string); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *users.ListPayload) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Remove provides a mock function with given fields: _a0, _a1
func (_m *Service) Remove(_a0 context.Context, _a1 *users.RemovePayload) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.RemovePayload) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Show provides a mock function with given fields: _a0, _a1
func (_m *Service) Show(_a0 context.Context, _a1 *users.ShowPayload) (*users.StoredUser, string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *users.StoredUser
	if rf, ok := ret.Get(0).(func(context.Context, *users.ShowPayload) *users.StoredUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.StoredUser)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *users.ShowPayload) string); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *users.ShowPayload) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Service) Update(_a0 context.Context, _a1 *users.User) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *users.User) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
