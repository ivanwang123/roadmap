package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *queryResolver) OneUser(ctx context.Context, input *model.GetUser) (*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetById(input.UserID)
}

func (r *queryResolver) OneRoadmap(ctx context.Context, input *model.GetRoadmap) (*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetById(input.RoadmapID)
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetAll()
}

func (r *queryResolver) AllRoadmaps(ctx context.Context) ([]*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetAll()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
