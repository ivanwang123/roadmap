package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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

func (r *queryResolver) Roadmaps(ctx context.Context) ([]*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetAll()
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userId := auth.ForContext(ctx)
	return stores.ForContext(ctx).UserStore.GetById(userId)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
