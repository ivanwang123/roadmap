package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/loaders"
)

func (r *checkpointResolver) Links(ctx context.Context, obj *model.Checkpoint) ([]*model.Link, error) {
	return obj.Links, nil
}

func (r *checkpointResolver) Status(ctx context.Context, obj *model.Checkpoint) (*string, error) {
	checkpointStatus, err := loaders.ForContext(ctx).CheckpointStatus(obj.ID)
	if err != nil {
		return nil, err
	}
	if checkpointStatus == nil {
		return nil, nil
	}

	return &checkpointStatus.Status, nil
}

func (r *checkpointResolver) Roadmap(ctx context.Context, obj *model.Checkpoint) (*model.Roadmap, error) {
	return loaders.ForContext(ctx).RoadmapById(obj.RoadmapID)
}

// Checkpoint returns generated.CheckpointResolver implementation.
func (r *Resolver) Checkpoint() generated.CheckpointResolver { return &checkpointResolver{r} }

type checkpointResolver struct{ *Resolver }
