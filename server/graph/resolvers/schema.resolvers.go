package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ivanwang123/roadmap/server/dataloaders"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return stores.ForContext(ctx).UserStore.Create(&input)
}

func (r *mutationResolver) CreateCheckpoint(ctx context.Context, input model.NewCheckpoint) (*model.Checkpoint, error) {
	return stores.ForContext(ctx).CheckpointStore.Create(&input)
}

func (r *mutationResolver) CreateRoadmap(ctx context.Context, input model.NewRoadmap) (*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.Create(&input)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetAll()
}

func (r *queryResolver) Roadmaps(ctx context.Context) ([]*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetAll()
}

func (r *roadmapResolver) Creator(ctx context.Context, obj *model.Roadmap) (*model.User, error) {
	return dataloaders.ForContext(ctx).UserById.Load(obj.CreatorID)
}

func (r *roadmapResolver) Checkpoints(ctx context.Context, obj *model.Roadmap) ([]*model.Checkpoint, error) {
	return stores.ForContext(ctx).CheckpointStore.GetByRoadmap(obj.ID)
}

func (r *roadmapResolver) Followers(ctx context.Context, obj *model.Roadmap) ([]*model.RoadmapFollower, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roadmapFollowerResolver) User(ctx context.Context, obj *model.RoadmapFollower) (*model.User, error) {
	return stores.ForContext(ctx).UserStore.GetById(obj.UserID)
}

func (r *roadmapFollowerResolver) Roadmap(ctx context.Context, obj *model.RoadmapFollower) (*model.Roadmap, error) {
	// TODO: Get roadmap by id
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Roadmap returns generated.RoadmapResolver implementation.
func (r *Resolver) Roadmap() generated.RoadmapResolver { return &roadmapResolver{r} }

// RoadmapFollower returns generated.RoadmapFollowerResolver implementation.
func (r *Resolver) RoadmapFollower() generated.RoadmapFollowerResolver {
	return &roadmapFollowerResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roadmapResolver struct{ *Resolver }
type roadmapFollowerResolver struct{ *Resolver }
