package checkpoint

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
}

type Usecase interface {
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
}
