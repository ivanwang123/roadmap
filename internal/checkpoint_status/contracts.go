package checkpoint_status

import (
	"context"

	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error)
	Update(ctx context.Context, input *models.UpdateStatus) error
	CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error
	DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error
}

type Usecase interface {
	Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error)
	Update(ctx context.Context, input *models.UpdateStatus) error
	CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error
	DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error
}
