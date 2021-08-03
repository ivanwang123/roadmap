package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/loaders"
	"github.com/ivanwang123/roadmap/models"
)

func (r *roadmapFollowerResolver) User(ctx context.Context, obj *models.RoadmapFollower) (*models.User, error) {
	return loaders.ForContext(ctx).UserById(obj.UserID)
}

func (r *roadmapFollowerResolver) Roadmap(ctx context.Context, obj *models.RoadmapFollower) (*models.Roadmap, error) {
	return loaders.ForContext(ctx).RoadmapById(obj.RoadmapID)
}

// RoadmapFollower returns generated.RoadmapFollowerResolver implementation.
func (r *Resolver) RoadmapFollower() generated.RoadmapFollowerResolver {
	return &roadmapFollowerResolver{r}
}

type roadmapFollowerResolver struct{ *Resolver }
