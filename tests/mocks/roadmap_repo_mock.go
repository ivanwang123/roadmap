// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/ivanwang123/roadmap/models"
	mock "github.com/stretchr/testify/mock"

	transaction "github.com/ivanwang123/roadmap/internal/common/transaction"
)

// RoadmapMockRepo is an autogenerated mock type for the Repository type
type RoadmapMockRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *RoadmapMockRepo) Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error) {
	ret := _m.Called(ctx, input)

	var r0 *models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewRoadmap) *models.Roadmap); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewRoadmap) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *RoadmapMockRepo) GetAll(ctx context.Context) ([]*models.Roadmap, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Roadmap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCreatorID provides a mock function with given fields: ctx, creatorID
func (_m *RoadmapMockRepo) GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error) {
	ret := _m.Called(ctx, creatorID)

	var r0 []*models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Roadmap); ok {
		r0 = rf(ctx, creatorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, creatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByFollower provides a mock function with given fields: ctx, userID
func (_m *RoadmapMockRepo) GetByFollower(ctx context.Context, userID int) ([]*models.Roadmap, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.Roadmap); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *RoadmapMockRepo) GetByID(ctx context.Context, ID int) (*models.Roadmap, error) {
	ret := _m.Called(ctx, ID)

	var r0 *models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Roadmap); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Roadmap)
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

// GetByPagination provides a mock function with given fields: ctx, input
func (_m *RoadmapMockRepo) GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error) {
	ret := _m.Called(ctx, input)

	var r0 []*models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetRoadmaps) []*models.Roadmap); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.GetRoadmaps) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIn provides a mock function with given fields: ctx, IDs
func (_m *RoadmapMockRepo) GetIn(ctx context.Context, IDs []string) ([]*models.Roadmap, error) {
	ret := _m.Called(ctx, IDs)

	var r0 []*models.Roadmap
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*models.Roadmap); ok {
		r0 = rf(ctx, IDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Roadmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTransaction provides a mock function with given fields: ctx, fn
func (_m *RoadmapMockRepo) WithTransaction(ctx context.Context, fn transaction.TxFunc) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, transaction.TxFunc) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}