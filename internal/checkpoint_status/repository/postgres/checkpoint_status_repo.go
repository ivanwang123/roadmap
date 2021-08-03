package postgres

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type checkpointStatusRepo struct {
	db *sqlx.DB
}

func NewCheckpointStatusRepo(db *sqlx.DB) checkpoint_status.Repository {
	return &checkpointStatusRepo{db}
}

func (r *checkpointStatusRepo) Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error) {
	return nil, nil
}

func (r *checkpointStatusRepo) Update(ctx context.Context, input *models.UpdateStatus) error {
	return nil
}

func (r *checkpointStatusRepo) CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error {
	return nil
}

func (r *checkpointStatusRepo) DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error {
	return nil
}
