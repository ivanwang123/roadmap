package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/server/dataloaders"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *userResolver) FollowingRoadmaps(ctx context.Context, obj *model.User) ([]*model.Roadmap, error) {
	return dataloaders.ForContext(ctx).RoadmapByFollower.Load(obj.ID)
	// return stores.ForContext(ctx).RoadmapFollowerStore.GetByUserId(obj.ID)
}

func (r *userResolver) CreatedRoadmaps(ctx context.Context, obj *model.User) ([]*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.GetByCreatorId(obj.ID)
	// return dataloaders.ForContext(ctx).RoadmapByCreatorId.Load(obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
