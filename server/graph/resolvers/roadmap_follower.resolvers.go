package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
)

func (r *roadmapFollowerResolver) User(ctx context.Context, obj *model.RoadmapFollower) (*model.User, error) {
	// return dataloaders.ForContext(ctx).UserById.Load(obj.UserID)
	panic(fmt.Errorf("not implemented"))
}

func (r *roadmapFollowerResolver) Roadmap(ctx context.Context, obj *model.RoadmapFollower) (*model.Roadmap, error) {
	panic(fmt.Errorf("not implemented"))
}

// RoadmapFollower returns generated.RoadmapFollowerResolver implementation.
func (r *Resolver) RoadmapFollower() generated.RoadmapFollowerResolver {
	return &roadmapFollowerResolver{r}
}

type roadmapFollowerResolver struct{ *Resolver }
