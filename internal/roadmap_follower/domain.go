package roadmap_follower

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	Get(ctx context.Context, userID, roadmapID int) (*models.RoadmapFollower, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error)
	Create(ctx context.Context, userID, roadmapID int) error
	Delete(ctx context.Context, userID, roadmapID int) error
	WithTransaction(ctx context.Context, fn transaction.TxFunc) error
}

type Usecase interface {
	Get(ctx context.Context, userID, roadmapID int) (*models.RoadmapFollower, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error)
}
