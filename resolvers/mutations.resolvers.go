package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ivanwang123/roadmap/graphql/generated"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/internal/common/cookie"
	"github.com/ivanwang123/roadmap/internal/stores"
	"github.com/ivanwang123/roadmap/models"
)

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (*models.User, error) {
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

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	cookie.ForContext(ctx).SetCookie("user", "", -time.Second)
	cookie.ForContext(ctx).SetCookie("refresh", "", -time.Second)
	fmt.Println("LOGOUT")
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	return stores.ForContext(ctx).UserStore.Create(&input)
}

func (r *mutationResolver) CreateCheckpoint(ctx context.Context, input models.NewCheckpoint) (*models.Checkpoint, error) {
	return stores.ForContext(ctx).CheckpointStore.Create(&input)
}

func (r *mutationResolver) CreateRoadmap(ctx context.Context, input models.NewRoadmap) (*models.Roadmap, error) {
	return stores.ForContext(ctx).RoadmapStore.Create(&input)
}

func (r *mutationResolver) ToggleFollowRoadmap(ctx context.Context, input models.FollowRoadmap) (*models.Roadmap, error) {
	userId := auth.ForContext(ctx)
	return stores.ForContext(ctx).RoadmapFollowerStore.ToggleFollowRoadmap(userId, &input)
}

func (r *mutationResolver) UpdateCheckpointStatus(ctx context.Context, input models.UpdateStatus) (*models.Checkpoint, error) {
	userId := auth.ForContext(ctx)
	return stores.ForContext(ctx).CheckpointStore.UpdateStatus(userId, &input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
