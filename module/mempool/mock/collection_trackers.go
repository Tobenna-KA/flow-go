// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import flow "github.com/dapperlabs/flow-go/model/flow"

import mock "github.com/stretchr/testify/mock"
import tracker "github.com/dapperlabs/flow-go/model/verification/tracker"

// CollectionTrackers is an autogenerated mock type for the CollectionTrackers type
type CollectionTrackers struct {
	mock.Mock
}

// Add provides a mock function with given fields: collt
func (_m *CollectionTrackers) Add(collt *tracker.CollectionTracker) bool {
	ret := _m.Called(collt)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*tracker.CollectionTracker) bool); ok {
		r0 = rf(collt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *CollectionTrackers) All() []*tracker.CollectionTracker {
	ret := _m.Called()

	var r0 []*tracker.CollectionTracker
	if rf, ok := ret.Get(0).(func() []*tracker.CollectionTracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*tracker.CollectionTracker)
		}
	}

	return r0
}

// ByCollectionID provides a mock function with given fields: collID
func (_m *CollectionTrackers) ByCollectionID(collID flow.Identifier) (*tracker.CollectionTracker, bool) {
	ret := _m.Called(collID)

	var r0 *tracker.CollectionTracker
	if rf, ok := ret.Get(0).(func(flow.Identifier) *tracker.CollectionTracker); ok {
		r0 = rf(collID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tracker.CollectionTracker)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(collID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Has provides a mock function with given fields: collID
func (_m *CollectionTrackers) Has(collID flow.Identifier) bool {
	ret := _m.Called(collID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(collID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Rem provides a mock function with given fields: collID
func (_m *CollectionTrackers) Rem(collID flow.Identifier) bool {
	ret := _m.Called(collID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(collID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
