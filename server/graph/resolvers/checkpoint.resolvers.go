package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/server/dataloaders"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
)

func (r *checkpointResolver) Roadmap(ctx context.Context, obj *model.Checkpoint) (*model.Roadmap, error) {
	return dataloaders.ForContext(ctx).RoadmapById.Load(obj.RoadmapID)
}

// Checkpoint returns generated.CheckpointResolver implementation.
func (r *Resolver) Checkpoint() generated.CheckpointResolver { return &checkpointResolver{r} }

type checkpointResolver struct{ *Resolver }
