package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/loaders"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *roadmapResolver) Creator(ctx context.Context, obj *model.Roadmap) (*model.User, error) {
	return loaders.ForContext(ctx).UserById(obj.CreatorID)
}

func (r *roadmapResolver) Checkpoints(ctx context.Context, obj *model.Roadmap) ([]*model.Checkpoint, error) {
	return stores.ForContext(ctx).CheckpointStore.GetByRoadmap(obj.ID)
}

func (r *roadmapResolver) Followers(ctx context.Context, obj *model.Roadmap) ([]*model.User, error) {
	return loaders.ForContext(ctx).UserByRoadmapFollowing(obj.ID)
}

// Roadmap returns generated.RoadmapResolver implementation.
func (r *Resolver) Roadmap() generated.RoadmapResolver { return &roadmapResolver{r} }

type roadmapResolver struct{ *Resolver }
