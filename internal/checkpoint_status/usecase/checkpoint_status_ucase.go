package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/models"
)

type checkpointStatusUsecase struct {
	checkpointStatusRepo checkpoint_status.Repository
	checkpointRepo       checkpoint.Repository
}

func NewCheckpointStatusUsecase(checkpointStatusRepo checkpoint_status.Repository, checkpointRepo checkpoint.Repository) checkpoint_status.Usecase {
	return &checkpointStatusUsecase{
		checkpointStatusRepo: checkpointStatusRepo,
		checkpointRepo:       checkpointRepo,
	}
}

func (u *checkpointStatusUsecase) Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error) {
	return u.checkpointStatusRepo.Get(ctx, input)
}

// TODO: Change return type?
func (u *checkpointStatusUsecase) Update(ctx context.Context, userID int, input *models.UpdateStatus) (*models.Checkpoint, error) {
	checkpoint, err := u.checkpointRepo.GetByID(ctx, input.CheckpointID)
	if err != nil {
		return nil, err
	}

	err = u.checkpointStatusRepo.Update(ctx, userID, input)
	if err != nil {
		return nil, err
	}

	checkpoint.Status = models.StatusType(input.Status)
	return checkpoint, nil
}

func (u *checkpointStatusUsecase) CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error {
	return u.checkpointStatusRepo.CreateMany(ctx, input)
}

func (u *checkpointStatusUsecase) DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error {
	return u.checkpointStatusRepo.DeleteMany(ctx, input)
}
