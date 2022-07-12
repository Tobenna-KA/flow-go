// Code generated by mockery v2.13.0. DO NOT EDIT.

package mocktracker

import (
	cid "github.com/ipfs/go-cid"
	mock "github.com/stretchr/testify/mock"
)

// TrackBlobsFn is an autogenerated mock type for the TrackBlobsFn type
type TrackBlobsFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields: blockHeight, cids
func (_m *TrackBlobsFn) Execute(blockHeight uint64, cids ...cid.Cid) error {
	_va := make([]interface{}, len(cids))
	for _i := range cids {
		_va[_i] = cids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, blockHeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, ...cid.Cid) error); ok {
		r0 = rf(blockHeight, cids...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewTrackBlobsFnT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTrackBlobsFn creates a new instance of TrackBlobsFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTrackBlobsFn(t NewTrackBlobsFnT) *TrackBlobsFn {
	mock := &TrackBlobsFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
