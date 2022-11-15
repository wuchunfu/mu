// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/aaronzjc/mu/internal/application/dto"
	mock "github.com/stretchr/testify/mock"
)

// FavorService is an autogenerated mock type for the FavorService type
type FavorService struct {
	mock.Mock
}

type FavorService_Expecter struct {
	mock *mock.Mock
}

func (_m *FavorService) EXPECT() *FavorService_Expecter {
	return &FavorService_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *FavorService) Add(_a0 context.Context, _a1 *dto.Favor) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.Favor) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FavorService_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type FavorService_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *dto.Favor
func (_e *FavorService_Expecter) Add(_a0 interface{}, _a1 interface{}) *FavorService_Add_Call {
	return &FavorService_Add_Call{Call: _e.mock.On("Add", _a0, _a1)}
}

func (_c *FavorService_Add_Call) Run(run func(_a0 context.Context, _a1 *dto.Favor)) *FavorService_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*dto.Favor))
	})
	return _c
}

func (_c *FavorService_Add_Call) Return(_a0 error) *FavorService_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

// Del provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *FavorService) Del(_a0 context.Context, _a1 int, _a2 string, _a3 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FavorService_Del_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Del'
type FavorService_Del_Call struct {
	*mock.Call
}

// Del is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
//   - _a2 string
//   - _a3 string
func (_e *FavorService_Expecter) Del(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *FavorService_Del_Call {
	return &FavorService_Del_Call{Call: _e.mock.On("Del", _a0, _a1, _a2, _a3)}
}

func (_c *FavorService_Del_Call) Run(run func(_a0 context.Context, _a1 int, _a2 string, _a3 string)) *FavorService_Del_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *FavorService_Del_Call) Return(_a0 error) *FavorService_Del_Call {
	_c.Call.Return(_a0)
	return _c
}

// UserFavorSites provides a mock function with given fields: _a0, _a1, _a2
func (_m *FavorService) UserFavorSites(_a0 context.Context, _a1 int, _a2 string) ([]string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, int, string) []string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FavorService_UserFavorSites_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserFavorSites'
type FavorService_UserFavorSites_Call struct {
	*mock.Call
}

// UserFavorSites is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
//   - _a2 string
func (_e *FavorService_Expecter) UserFavorSites(_a0 interface{}, _a1 interface{}, _a2 interface{}) *FavorService_UserFavorSites_Call {
	return &FavorService_UserFavorSites_Call{Call: _e.mock.On("UserFavorSites", _a0, _a1, _a2)}
}

func (_c *FavorService_UserFavorSites_Call) Run(run func(_a0 context.Context, _a1 int, _a2 string)) *FavorService_UserFavorSites_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string))
	})
	return _c
}

func (_c *FavorService_UserFavorSites_Call) Return(_a0 []string, _a1 error) *FavorService_UserFavorSites_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UserFavors provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *FavorService) UserFavors(_a0 context.Context, _a1 int, _a2 string, _a3 string) ([]*dto.Favor, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*dto.Favor
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) []*dto.Favor); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.Favor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FavorService_UserFavors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserFavors'
type FavorService_UserFavors_Call struct {
	*mock.Call
}

// UserFavors is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int
//   - _a2 string
//   - _a3 string
func (_e *FavorService_Expecter) UserFavors(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *FavorService_UserFavors_Call {
	return &FavorService_UserFavors_Call{Call: _e.mock.On("UserFavors", _a0, _a1, _a2, _a3)}
}

func (_c *FavorService_UserFavors_Call) Run(run func(_a0 context.Context, _a1 int, _a2 string, _a3 string)) *FavorService_UserFavors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *FavorService_UserFavors_Call) Return(_a0 []*dto.Favor, _a1 error) *FavorService_UserFavors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewFavorService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFavorService creates a new instance of FavorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFavorService(t mockConstructorTestingTNewFavorService) *FavorService {
	mock := &FavorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
