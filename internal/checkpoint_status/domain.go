package checkpoint_status

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/common/transaction"
	"github.com/ivanwang123/roadmap/models"
)

type Repository interface {
	Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error)
	GetIn(ctx context.Context, userID int, IDs []string) ([]*models.CheckpointStatus, error)
	Update(ctx context.Context, userID int, input *models.UpdateStatus) error
	CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error
	DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error
	WithTransaction(ctx context.Context, fn transaction.TxFunc) error
}

type Usecase interface {
	Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error)
	CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error
	DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error
	BatchGet(ctx context.Context) func(int) (*models.CheckpointStatus, error)
}
