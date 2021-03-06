// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/ivanwang123/roadmap/models"
	mock "github.com/stretchr/testify/mock"

	transaction "github.com/ivanwang123/roadmap/internal/common/transaction"
)

// CheckpointMockRepo is an autogenerated mock type for the Repository type
type CheckpointMockRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *CheckpointMockRepo) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	ret := _m.Called(ctx, input)

	var r0 *models.Checkpoint
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCheckpoint) *models.Checkpoint); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Checkpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCheckpoint) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *CheckpointMockRepo) GetByID(ctx context.Context, ID int) (*models.Checkpoint, error) {
	ret := _m.Called(ctx, ID)

	var r0 *models.Checkpoint
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Checkpoint); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Checkpoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRoadmap provides a mock function with given fields: ctx, roadmapID
func (_m *CheckpointMockRepo) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error) {
	ret := _m.Called(ctx, roadmapID)

	var r0 []*models.Checkpoint
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Checkpoint); ok {
		r0 = rf(ctx, roadmapID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Checkpoint)
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

// GetIDByRoadmap provides a mock function with given fields: ctx, roadmapID
func (_m *CheckpointMockRepo) GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error) {
	ret := _m.Called(ctx, roadmapID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int) []int); ok {
		r0 = rf(ctx, roadmapID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
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
func (_m *CheckpointMockRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, transaction.TxFunc) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
