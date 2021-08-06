package usecase

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/internal/checkpoint"
	"github.com/ivanwang123/roadmap/internal/checkpoint_status"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/internal/common/loader"
	"github.com/ivanwang123/roadmap/models"
)

type checkpointStatusUsecase struct {
	checkpointStatusRepo checkpoint_status.Repository
	checkpointRepo       checkpoint.Repository
}

func NewCheckpointStatusUsecase(cs checkpoint_status.Repository, c checkpoint.Repository) checkpoint_status.Usecase {
	return &checkpointStatusUsecase{
		checkpointStatusRepo: cs,
		checkpointRepo:       c,
	}
}

func (u *checkpointStatusUsecase) Get(ctx context.Context, input *models.GetCheckpointStatus) (*models.CheckpointStatus, error) {
	return u.checkpointStatusRepo.Get(ctx, input)
}

func (u *checkpointStatusUsecase) CreateMany(ctx context.Context, input []*models.CreateCheckpointStatus) error {
	return u.checkpointStatusRepo.CreateMany(ctx, input)
}

func (u *checkpointStatusUsecase) DeleteMany(ctx context.Context, input *models.DeleteManyCheckpointStatus) error {
	return u.checkpointStatusRepo.DeleteMany(ctx, input)
}

func (u *checkpointStatusUsecase) BatchGet(ctx context.Context) func(int) (*models.CheckpointStatus, error) {
	batchedLoader := loader.NewLoader(func(ctx context.Context, keys []string) (map[string]interface{}, error) {
		userID, err := auth.GetCurrentUser(ctx)
		if err != nil {
			return nil, err
		}

		statuses, err := u.checkpointStatusRepo.GetIn(ctx, userID, keys)
		if err != nil {
			return nil, err
		}

		statusesMap := make(map[string]interface{})
		for _, status := range statuses {
			statusesMap[strconv.Itoa(status.CheckpointID)] = status
		}
		return statusesMap, nil
	})

	return func(checkpointID int) (*models.CheckpointStatus, error) {
		result, err := batchedLoader.Load(ctx, dataloader.StringKey(strconv.Itoa(checkpointID)))()
		if err != nil {
			return nil, err
		}

		status, _ := result.(*models.CheckpointStatus)
		return status, nil
	}
}
