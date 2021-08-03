package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/models"
)

type checkpointStatusUsecase struct {
	checkpointStatusRepo checkpoint_status.Repository
}

func NewCheckpointStatusUsecase(checkpointStatusRepo checkpoint_status.Repository) checkpoint_status.Usecase {
	return &checkpointStatusUsecase{
		checkpointStatusRepo: checkpointStatusRepo,
	}
}

// TODO: Move repository inputs to models
func (u *checkpointStatusUsecase) Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error) {
	return u.checkpointStatusRepo.Get(ctx, input)
}

func (u *checkpointStatusUsecase) Update(ctx context.Context, input *models.UpdateStatus) error {
	return u.checkpointStatusRepo.Update(ctx, input)
}

func (u *checkpointStatusUsecase) CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error {
	return u.checkpointStatusRepo.CreateMany(ctx, input)
}

func (u *checkpointStatusUsecase) DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error {
	return u.checkpointStatusRepo.DeleteMany(ctx, input)
}
