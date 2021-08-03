package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/models"
)

type checkpointUsecase struct {
	checkpointRepo checkpoint.Repository
}

func NewCheckpointUsecase(checkpointRepo checkpoint.Repository) checkpoint.Usecase {
	return &checkpointUsecase{
		checkpointRepo: checkpointRepo,
	}
}

func (u *checkpointUsecase) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error) {
	return u.checkpointRepo.GetByRoadmap(ctx, roadmapID)
}

func (u *checkpointUsecase) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	return u.checkpointRepo.Create(ctx, input)
}
