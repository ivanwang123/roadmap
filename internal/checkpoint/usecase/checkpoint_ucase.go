package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/models"
)

type checkpointUsecase struct {
	checkpointRepo       checkpoint.Repository
	checkpointStatusRepo checkpoint_status.Repository
	roadmapFollowerRepo  roadmap_follower.Repository
}

func NewCheckpointUsecase(checkpointRepo checkpoint.Repository, checkpointStatusRepo checkpoint_status.Repository, roadmapFollowerRepo roadmap_follower.Repository) checkpoint.Usecase {
	return &checkpointUsecase{
		checkpointRepo:       checkpointRepo,
		checkpointStatusRepo: checkpointStatusRepo,
		roadmapFollowerRepo:  roadmapFollowerRepo,
	}
}

func (u *checkpointUsecase) GetByID(ctx context.Context, ID int) (*models.Checkpoint, error) {
	return u.checkpointRepo.GetByID(ctx, ID)
}

func (u *checkpointUsecase) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.Checkpoint, error) {
	return u.checkpointRepo.GetByRoadmap(ctx, roadmapID)
}

func (u *checkpointUsecase) GetIDByRoadmap(ctx context.Context, roadmapID int) ([]int, error) {
	return u.checkpointRepo.GetIDByRoadmap(ctx, roadmapID)
}

// TODO: Put in transaction?
func (u *checkpointUsecase) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	checkpoint, err := u.checkpointRepo.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	roadmapFollowers, err := u.roadmapFollowerRepo.GetByRoadmap(ctx, input.RoadmapID)
	if err != nil {
		return nil, err
	}

	newCheckpointStatuses := make([]*models.CreateCheckpointStatus, len(roadmapFollowers))
	for i, follower := range roadmapFollowers {
		newCheckpointStatuses[i] = &models.CreateCheckpointStatus{
			UserID:       follower.UserID,
			CheckpointID: checkpoint.ID,
			RoadmapID:    input.RoadmapID,
		}
	}

	if err := u.checkpointStatusRepo.CreateMany(ctx, newCheckpointStatuses); err != nil {
		return nil, err
	}

	return checkpoint, nil
}
