package checkpoint

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetByID(ctx context.Context, ID int) (*models.Checkpoint, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
}

type Usecase interface {
	GetByID(ctx context.Context, ID int) (*models.Checkpoint, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
}
