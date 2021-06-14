package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
)

func (r *userResolver) FollowingRoadmaps(ctx context.Context, obj *model.User) ([]*model.Roadmap, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) CreatedRoadmaps(ctx context.Context, obj *model.User) ([]*model.Roadmap, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
