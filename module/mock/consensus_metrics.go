// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ConsensusMetrics is an autogenerated mock type for the ConsensusMetrics type
type ConsensusMetrics struct {
	mock.Mock
}

// CheckSealingDuration provides a mock function with given fields: duration
func (_m *ConsensusMetrics) CheckSealingDuration(duration time.Duration) {
	_m.Called(duration)
}

// EmergencySeal provides a mock function with given fields:
func (_m *ConsensusMetrics) EmergencySeal() {
	_m.Called()
}

// FinishBlockToSeal provides a mock function with given fields: blockID
func (_m *ConsensusMetrics) FinishBlockToSeal(blockID flow.Identifier) {
	_m.Called(blockID)
}

// FinishCollectionToFinalized provides a mock function with given fields: collectionID
func (_m *ConsensusMetrics) FinishCollectionToFinalized(collectionID flow.Identifier) {
	_m.Called(collectionID)
}

// OnApprovalProcessingDuration provides a mock function with given fields: duration
func (_m *ConsensusMetrics) OnApprovalProcessingDuration(duration time.Duration) {
	_m.Called(duration)
}

// OnReceiptProcessingDuration provides a mock function with given fields: duration
func (_m *ConsensusMetrics) OnReceiptProcessingDuration(duration time.Duration) {
	_m.Called(duration)
}

// StartBlockToSeal provides a mock function with given fields: blockID
func (_m *ConsensusMetrics) StartBlockToSeal(blockID flow.Identifier) {
	_m.Called(blockID)
}

// StartCollectionToFinalized provides a mock function with given fields: collectionID
func (_m *ConsensusMetrics) StartCollectionToFinalized(collectionID flow.Identifier) {
	_m.Called(collectionID)
}

type mockConstructorTestingTNewConsensusMetrics interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsensusMetrics creates a new instance of ConsensusMetrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsensusMetrics(t mockConstructorTestingTNewConsensusMetrics) *ConsensusMetrics {
	mock := &ConsensusMetrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
