// Code generated by mockery v2.12.3. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// BlockRequester is an autogenerated mock type for the BlockRequester type
type BlockRequester struct {
	mock.Mock
}

// Prune provides a mock function with given fields: final
func (_m *BlockRequester) Prune(final *flow.Header) {
	_m.Called(final)
}

// RequestBlock provides a mock function with given fields: blockID
func (_m *BlockRequester) RequestBlock(blockID flow.Identifier) {
	_m.Called(blockID)
}

// RequestHeight provides a mock function with given fields: height
func (_m *BlockRequester) RequestHeight(height uint64) {
	_m.Called(height)
}

type NewBlockRequesterT interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockRequester creates a new instance of BlockRequester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockRequester(t NewBlockRequesterT) *BlockRequester {
	mock := &BlockRequester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
