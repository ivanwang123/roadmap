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

func NewCheckpointUsecase(c checkpoint.Repository, cs checkpoint_status.Repository, rf roadmap_follower.Repository) checkpoint.Usecase {
	return &checkpointUsecase{
		checkpointRepo:       c,
		checkpointStatusRepo: cs,
		roadmapFollowerRepo:  rf,
	}
}

func SetRepo(c *checkpointUsecase) {

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

func (u *checkpointUsecase) Create(ctx context.Context, input *models.NewCheckpoint) (*models.Checkpoint, error) {
	var checkpoint *models.Checkpoint

	err := u.checkpointRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		var err error
		checkpoint, err = u.checkpointRepo.Create(txCtx, input)
		if err != nil {
			return err
		}

		roadmapFollowers, err := u.roadmapFollowerRepo.GetByRoadmap(txCtx, input.RoadmapID)
		if err != nil {
			return err
		}

		newCheckpointStatuses := make([]*models.CreateCheckpointStatus, len(roadmapFollowers))
		for i, follower := range roadmapFollowers {
			newCheckpointStatuses[i] = &models.CreateCheckpointStatus{
				UserID:       follower.UserID,
				CheckpointID: checkpoint.ID,
				RoadmapID:    input.RoadmapID,
			}
		}

		if err := u.checkpointStatusRepo.CreateMany(txCtx, newCheckpointStatuses); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return checkpoint, nil
}

func (u *checkpointUsecase) UpdateStatus(ctx context.Context, userID int, input *models.UpdateStatus) (*models.Checkpoint, error) {
	var checkpoint *models.Checkpoint

	err := u.checkpointRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		var err error
		checkpoint, err = u.checkpointRepo.GetByID(ctx, input.CheckpointID)
		if err != nil {
			return err
		}

		err = u.checkpointStatusRepo.Update(ctx, userID, input)
		if err != nil {
			return err
		}

		checkpoint.Status = models.StatusType(input.Status)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return checkpoint, nil
}
