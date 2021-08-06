package checkpoint

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	GetByID(ctx context.Context, ID int) (*models.Checkpoint, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
	WithTransaction(ctx context.Context, fn transaction.TxFunc) error
}

type Usecase interface {
	GetByID(ctx context.Context, ID int) (*models.Checkpoint, error)
	GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error)
	GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error)
	Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error)
	UpdateStatus(ctx context.Context, userID int, input *models.UpdateStatus) (*models.Checkpoint, error)
}
