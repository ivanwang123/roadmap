package roadmap_follower

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	Get(ctx context.Context, userID, roadmapID int) (*models.RoadmapFollower, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error)
	Create(ctx context.Context, userID, roadmapID int) error
	Delete(ctx context.Context, userID, roadmapID int) error
}

type Usecase interface {
	Get(ctx context.Context, userID, roadmapID int) (*models.RoadmapFollower, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error)
	ToggleFollowRoadmap(ctx context.Context, userID, roadmapID int) (*models.Roadmap, error)
}
