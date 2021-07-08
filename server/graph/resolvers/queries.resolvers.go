package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ivanwang123/roadmap/server/auth"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *queryResolver) User(ctx context.Context, input *model.GetUser) (*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetById(input.ID)
}

func (r *queryResolver) Roadmap(ctx context.Context, input *model.GetRoadmap) (*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetById(input.ID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetAll()
}

func (r *queryResolver) Roadmaps(ctx context.Context, input *model.GetRoadmaps) ([]*model.Roadmap, error) {
	switch input.Sort {
	case model.SortNewest:
		return stores.ForContext(ctx).RoadmapStore.GetByNewest(input.CursorID, input.CursorValue)
	case model.SortOldest:
		return stores.ForContext(ctx).RoadmapStore.GetByOldest(input.CursorID, input.CursorValue)
	case model.SortMostFollowers:
		return stores.ForContext(ctx).RoadmapStore.GetByMostFollowers(input.CursorID, input.CursorValue)
	case model.SortMostCheckpoints:
		return stores.ForContext(ctx).RoadmapStore.GetByMostCheckpoints(input.CursorID, input.CursorValue)
	case model.SortLeastCheckpoints:
		return stores.ForContext(ctx).RoadmapStore.GetByLeastCheckpoints(input.CursorID, input.CursorValue)
	default:
		return nil, errors.New("Invalid sort option")
	}
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userId := auth.ForContext(ctx)
	if userId < 0 {
		return nil, nil
	}
	return stores.ForContext(ctx).UserStore.GetById(userId)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
