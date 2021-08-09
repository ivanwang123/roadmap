package usecase

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/internal/common/loader"
	"github.com/ivanwang123/roadmap/internal/roadmap"
	"github.com/ivanwang123/roadmap/internal/roadmap_follower"
	"github.com/ivanwang123/roadmap/models"
)

type roadmapUsecase struct {
	roadmapRepo          roadmap.Repository
	roadmapFollowerRepo  roadmap_follower.Repository
	checkpointRepo       checkpoint.Repository
	checkpointStatusRepo checkpoint_status.Repository
}

func NewRoadmapUsecase(r roadmap.Repository, rf roadmap_follower.Repository, c checkpoint.Repository, cs checkpoint_status.Repository) roadmap.Usecase {
	return &roadmapUsecase{
		roadmapRepo:          r,
		roadmapFollowerRepo:  rf,
		checkpointRepo:       c,
		checkpointStatusRepo: cs,
	}
}

func (u *roadmapUsecase) GetAll(ctx context.Context) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetAll(ctx)
}

func (u *roadmapUsecase) GetByID(ctx context.Context, ID int) (*models.Roadmap, error) {
	return u.roadmapRepo.GetByID(ctx, ID)
}

func (u *roadmapUsecase) GetByCreatorID(ctx context.Context, creatorID int) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetByCreatorID(ctx, creatorID)
}

func (u *roadmapUsecase) GetByFollower(ctx context.Context, userID int) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetByFollower(ctx, userID)
}

func (u *roadmapUsecase) GetByPagination(ctx context.Context, input *models.GetRoadmaps) ([]*models.Roadmap, error) {
	return u.roadmapRepo.GetByPagination(ctx, input)
}

func (u *roadmapUsecase) Create(ctx context.Context, input *models.NewRoadmap) (*models.Roadmap, error) {
	return u.roadmapRepo.Create(ctx, input)
}

func (u *roadmapUsecase) ToggleFollow(ctx context.Context, userID, roadmapID int) (*models.Roadmap, error) {
	var roadmap *models.Roadmap

	follower, _ := u.roadmapFollowerRepo.Get(ctx, userID, roadmapID)

	err := u.roadmapRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		checkpointIDs, err := u.checkpointRepo.GetIDByRoadmap(txCtx, roadmapID)
		if err != nil {
			return err
		}

		if follower != nil {
			err = u.roadmapFollowerRepo.Delete(txCtx, userID, roadmapID)
			if err != nil {
				return err
			}

			deleteCheckpointStatuses := &models.DeleteManyCheckpointStatus{
				RoadmapID:     roadmapID,
				UserIDs:       []int{userID},
				CheckpointIDs: checkpointIDs,
			}
			if err = u.checkpointStatusRepo.DeleteMany(txCtx, deleteCheckpointStatuses); err != nil {
				return err
			}
		} else {
			err = u.roadmapFollowerRepo.Create(ctx, userID, roadmapID)
			if err != nil {
				return err
			}

			newCheckpointStatuses := make([]*models.CreateCheckpointStatus, len(checkpointIDs))
			for i, checkpointId := range checkpointIDs {
				newCheckpointStatuses[i] = &models.CreateCheckpointStatus{UserID: userID, CheckpointID: int(checkpointId), RoadmapID: roadmapID}
			}
			if err = u.checkpointStatusRepo.CreateMany(txCtx, newCheckpointStatuses); err != nil {
				return err
			}
		}

		roadmap, err = u.roadmapRepo.GetByID(txCtx, roadmapID)
		return err
	})

	if err != nil {
		return nil, err
	}
	return roadmap, nil
}

func (u *roadmapUsecase) BatchGet(ctx context.Context) func(int) (*models.Roadmap, error) {
	batchLoader := loader.NewLoader(func(ctx context.Context, keys []string) (map[string]interface{}, error) {
		roadmaps, err := u.roadmapRepo.GetIn(ctx, keys)
		if err != nil {
			return nil, err
		}

		roadmapsMap := make(map[string]interface{})
		for _, roadmap := range roadmaps {
			roadmapsMap[strconv.Itoa(roadmap.ID)] = roadmap
		}
		return roadmapsMap, nil
	})

	return func(roadmapID int) (*models.Roadmap, error) {
		result, err := batchLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(roadmapID)))()
		if err != nil {
			return nil, err
		}

		roadmap, _ := result.(*models.Roadmap)
		return roadmap, nil
	}
}
