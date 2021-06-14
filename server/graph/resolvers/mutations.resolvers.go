package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

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

func (r *mutationResolver) FollowRoadmap(ctx context.Context, input model.FollowRoadmap) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCheckpoint(ctx context.Context, input model.AddCheckpoint) (*model.Checkpoint, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
