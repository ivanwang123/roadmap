package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/models"
)

type roadmapFollowerUsecase struct {
	roadmapFollowerRepo roadmap_follower.Repository
}

func NewRoadmapFollowerUsecase(rf roadmap_follower.Repository) roadmap_follower.Usecase {
	return &roadmapFollowerUsecase{
		roadmapFollowerRepo: rf,
	}
}

func (u *roadmapFollowerUsecase) Get(ctx context.Context, userID int, roadmapID int) (*models.RoadmapFollower, error) {
	return u.roadmapFollowerRepo.Get(ctx, userID, roadmapID)
}

func (u *roadmapFollowerUsecase) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error) {
	return u.roadmapFollowerRepo.GetByRoadmap(ctx, roadmapID)
}
