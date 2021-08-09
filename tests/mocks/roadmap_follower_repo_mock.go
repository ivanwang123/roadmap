// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/ivanwang123/roadmap/models"
	mock "github.com/stretchr/testify/mock"

	transaction "github.com/ivanwang123/roadmap/internal/common/transaction"
)

// RoadmapFollowerMockRepo is an autogenerated mock type for the Repository type
type RoadmapFollowerMockRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, userID, roadmapID
func (_m *RoadmapFollowerMockRepo) Create(ctx context.Context, userID int, roadmapID int) error {
	ret := _m.Called(ctx, userID, roadmapID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userID, roadmapID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, userID, roadmapID
func (_m *RoadmapFollowerMockRepo) Delete(ctx context.Context, userID int, roadmapID int) error {
	ret := _m.Called(ctx, userID, roadmapID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userID, roadmapID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, userID, roadmapID
func (_m *RoadmapFollowerMockRepo) Get(ctx context.Context, userID int, roadmapID int) (*models.RoadmapFollower, error) {
	ret := _m.Called(ctx, userID, roadmapID)

	var r0 *models.RoadmapFollower
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *models.RoadmapFollower); ok {
		r0 = rf(ctx, userID, roadmapID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RoadmapFollower)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, userID, roadmapID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRoadmap provides a mock function with given fields: ctx, roadmapID
func (_m *RoadmapFollowerMockRepo) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error) {
	ret := _m.Called(ctx, roadmapID)

	var r0 []*models.RoadmapFollower
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.RoadmapFollower); ok {
		r0 = rf(ctx, roadmapID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.RoadmapFollower)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, roadmapID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTransaction provides a mock function with given fields: ctx, fn
func (_m *RoadmapFollowerMockRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, transaction.TxFunc) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}