package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/models"
)

func (r *roadmapResolver) Creator(ctx context.Context, obj *models.Roadmap) (*models.User, error) {
	return r.UserUsecase.BatchGet(ctx)(obj.CreatorID)
}

func (r *roadmapResolver) Checkpoints(ctx context.Context, obj *models.Roadmap) ([]*models.Checkpoint, error) {
	return r.CheckpointUsecase.GetByRoadmap(ctx, obj.ID)
}

func (r *roadmapResolver) Followers(ctx context.Context, obj *models.Roadmap) ([]*models.User, error) {
	return r.UserUsecase.GetByRoadmapFollowing(ctx, obj.ID)
}

// Roadmap returns generated.RoadmapResolver implementation.
func (r *Resolver) Roadmap() generated.RoadmapResolver { return &roadmapResolver{r} }

type roadmapResolver struct{ *Resolver }
