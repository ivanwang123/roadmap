package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/loaders"
	"github.com/ivanwang123/roadmap/models"
)

func (r *checkpointResolver) Links(ctx context.Context, obj *models.Checkpoint) ([]*models.Link, error) {
	return obj.Links, nil
}

func (r *checkpointResolver) Status(ctx context.Context, obj *models.Checkpoint) (*models.Status, error) {
	checkpointStatus, err := loaders.ForContext(ctx).CheckpointStatus(obj.ID)
	if err != nil {
		return nil, nil
	}
	if checkpointStatus == nil {
		return nil, nil
	}

	return &checkpointStatus.Status, nil
}

func (r *checkpointResolver) Roadmap(ctx context.Context, obj *models.Checkpoint) (*models.Roadmap, error) {
	return loaders.ForContext(ctx).RoadmapById(obj.RoadmapID)
}

// Checkpoint returns generated.CheckpointResolver implementation.
func (r *Resolver) Checkpoint() generated.CheckpointResolver { return &checkpointResolver{r} }

type checkpointResolver struct{ *Resolver }
