package usecase

import (
	"context"

	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/internal/roadmap"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/models"
)

type roadmapFollowerUsecase struct {
	roadmapFollowerRepo  roadmap_follower.Repository
	roadmapRepo          roadmap.Repository
	checkpointRepo       checkpoint.Repository
	checkpointStatusRepo checkpoint_status.Repository
}

func NewRoadmapFollowerUsecase(rf roadmap_follower.Repository, r roadmap.Repository, c checkpoint.Repository, cs checkpoint_status.Repository) roadmap_follower.Usecase {
	return &roadmapFollowerUsecase{
		roadmapFollowerRepo:  rf,
		roadmapRepo:          r,
		checkpointRepo:       c,
		checkpointStatusRepo: cs,
	}
}

func (u *roadmapFollowerUsecase) Get(ctx context.Context, userID int, roadmapID int) (*models.RoadmapFollower, error) {
	return u.roadmapFollowerRepo.Get(ctx, userID, roadmapID)
}

func (u *roadmapFollowerUsecase) GetByRoadmap(ctx context.Context, roadmapID int) ([]*models.RoadmapFollower, error) {
	return u.roadmapFollowerRepo.GetByRoadmap(ctx, roadmapID)
}

// TODO: Move to roadmap?
func (u *roadmapFollowerUsecase) ToggleFollowRoadmap(ctx context.Context, userID, roadmapID int) (*models.Roadmap, error) {
	err := u.roadmapFollowerRepo.Create(ctx, userID, roadmapID)
	if err != nil {
		err = u.roadmapFollowerRepo.Delete(ctx, userID, roadmapID)

		checkpointIDs, err := u.checkpointRepo.GetIDByRoadmap(ctx, roadmapID)
		if err != nil {
			return nil, err
		}

		deleteCheckpointStatuses := &models.DeleteManyCheckpointStatus{
			RoadmapID:     roadmapID,
			UserIDs:       []int{userID},
			CheckpointIDs: checkpointIDs,
		}
		if err := u.checkpointStatusRepo.DeleteMany(ctx, deleteCheckpointStatuses); err != nil {
			return nil, err
		}

		return u.roadmapRepo.GetByID(ctx, roadmapID)
	}

	checkpointIDs, err := u.checkpointRepo.GetIDByRoadmap(ctx, roadmapID)
	if err != nil {
		return nil, err
	}

	newCheckpointStatuses := make([]*models.CreateCheckpointStatus, len(checkpointIDs))
	for i, checkpointId := range checkpointIDs {
		newCheckpointStatuses[i] = &models.CreateCheckpointStatus{UserID: userID, CheckpointID: int(checkpointId), RoadmapID: roadmapID}
	}
	if err := u.checkpointStatusRepo.CreateMany(ctx, newCheckpointStatuses); err != nil {
		return nil, err
	}

	return u.roadmapRepo.GetByID(ctx, roadmapID)
}
