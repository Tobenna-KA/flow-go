// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// NewJobListener is an autogenerated mock type for the NewJobListener type
type NewJobListener struct {
	mock.Mock
}

// Check provides a mock function with given fields:
func (_m *NewJobListener) Check() {
	_m.Called()
}

type mockConstructorTestingTNewNewJobListener interface {
	mock.TestingT
	Cleanup(func())
}

// NewNewJobListener creates a new instance of NewJobListener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNewJobListener(t mockConstructorTestingTNewNewJobListener) *NewJobListener {
	mock := &NewJobListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
