// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	userspb "users/gen/grpc/users/pb"

	mock "github.com/stretchr/testify/mock"
)

// UsersServer is an autogenerated mock type for the UsersServer type
type UsersServer struct {
	mock.Mock
}

// Activate provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) Activate(_a0 context.Context, _a1 *userspb.ActivateRequest) (*userspb.ActivateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.ActivateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.ActivateRequest) *userspb.ActivateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.ActivateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.ActivateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) Add(_a0 context.Context, _a1 *userspb.AddRequest) (*userspb.AddResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.AddResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.AddRequest) *userspb.AddResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.AddResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.AddRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) List(_a0 context.Context, _a1 *userspb.ListRequest) (*userspb.StoredUserCollection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.StoredUserCollection
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.ListRequest) *userspb.StoredUserCollection); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.StoredUserCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.ListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) Remove(_a0 context.Context, _a1 *userspb.RemoveRequest) (*userspb.RemoveResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.RemoveResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.RemoveRequest) *userspb.RemoveResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.RemoveResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.RemoveRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Show provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) Show(_a0 context.Context, _a1 *userspb.ShowRequest) (*userspb.ShowResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.ShowResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.ShowRequest) *userspb.ShowResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.ShowResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.ShowRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *UsersServer) Update(_a0 context.Context, _a1 *userspb.UpdateRequest) (*userspb.UpdateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *userspb.UpdateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *userspb.UpdateRequest) *userspb.UpdateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userspb.UpdateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *userspb.UpdateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
