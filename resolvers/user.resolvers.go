package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/loaders"
	"github.com/ivanwang123/roadmap/models"
)

func (r *userResolver) FollowingRoadmaps(ctx context.Context, obj *models.User) ([]*models.Roadmap, error) {
	return loaders.ForContext(ctx).RoadmapByFollower(obj.ID)
}

func (r *userResolver) CreatedRoadmaps(ctx context.Context, obj *models.User) ([]*models.Roadmap, error) {
	return r.RoadmapUsecase.GetByCreatorID(ctx, obj.ID)
	// return stores.ForContext(ctx).RoadmapStore.GetByCreatorId(obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
