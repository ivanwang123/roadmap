package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ivanwang123/roadmap/server/auth"
	"github.com/ivanwang123/roadmap/server/cookie"
	"github.com/ivanwang123/roadmap/server/graph/generated"
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/ivanwang123/roadmap/server/stores"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return stores.ForContext(ctx).UserStore.Create(&input)
}

func (r *mutationResolver) CreateCheckpoint(ctx context.Context, input model.NewCheckpoint) (*model.Checkpoint, error) {
	return stores.ForContext(ctx).CheckpointStore.Create(&input)
}

func (r *mutationResolver) CreateRoadmap(ctx context.Context, input model.NewRoadmap) (*model.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.Create(&input)
}

func (r *mutationResolver) ToggleFollowRoadmap(ctx context.Context, input model.FollowRoadmap) (bool, error) {
	return stores.ForContext(ctx).RoadmapFollowerStore.ToggleFollowRoadmap(&input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.User, error) {
	user, err := stores.ForContext(ctx).UserStore.Authenticate(&input)
	if err != nil {
		return nil, err
	}

	userToken, err := auth.GenerateToken(user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}
	refreshToken, err := auth.GenerateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	cookie.ForContext(ctx).SetCookie("user", userToken, time.Hour*24)
	cookie.ForContext(ctx).SetCookie("refresh", refreshToken, time.Hour*24*7)
	fmt.Println("LOGIN SET COOKIES")

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
