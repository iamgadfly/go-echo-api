// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/iamgadfly/go-echo-api/pkg/api/user_v1/proto"
	mock "github.com/stretchr/testify/mock"
)

// UserServer is an autogenerated mock type for the UserServer type
type UserServer struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *UserServer) Get(_a0 context.Context, _a1 *proto.UserIdRequest) (*proto.UserResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *proto.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UserIdRequest) (*proto.UserResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UserIdRequest) *proto.UserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.UserIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: _a0, _a1
func (_m *UserServer) GetUsers(_a0 context.Context, _a1 *proto.EmptyParams) (*proto.UserResponseList, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 *proto.UserResponseList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.EmptyParams) (*proto.UserResponseList, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.EmptyParams) *proto.UserResponseList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UserResponseList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.EmptyParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedUserServer provides a mock function with given fields:
func (_m *UserServer) mustEmbedUnimplementedUserServer() {
	_m.Called()
}

// NewUserServer creates a new instance of UserServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServer {
	mock := &UserServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
