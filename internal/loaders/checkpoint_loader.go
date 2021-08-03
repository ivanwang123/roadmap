package loaders

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/database"
	"github.com/ivanwang123/roadmap/internal/common/auth"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

func CheckpointStatus(ctx context.Context) func(int) (*models.CheckpointStatus, error) {
	loader := dataloader.NewBatchedLoader(checkpointStatusBatchFn)

	return func(checkpointId int) (*models.CheckpointStatus, error) {
		result, err := loader.Load(ctx, dataloader.StringKey(strconv.Itoa(checkpointId)))()
		if err != nil {
			return nil, err
		}

		status, _ := result.(*models.CheckpointStatus)
		return status, nil
	}
}

func checkpointStatusBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	statusesMap := make(map[string]interface{})
	statuses := []*models.CheckpointStatus{}

	db := database.ForContext(ctx)
	userId := auth.ForContext(ctx)
	if userId < 0 {
		return handleNilData(results)
	}

	query, args, err := sqlx.In("SELECT * FROM checkpoint_status WHERE user_id = ? AND checkpoint_id IN (?)", userId, ToStrKeys(keys))
	if err != nil {
		return HandleError(err, results)
	}

	if err := db.Select(&statuses, db.Rebind(query), args...); err != nil {
		return HandleError(err, results)
	}

	if len(statuses) == 0 {
		return handleNilData(results)
	}

	for _, status := range statuses {
		statusesMap[strconv.Itoa(status.CheckpointID)] = status
	}

	SortResults(results, keys, statusesMap)

	return results
}

func handleNilData(results []*dataloader.Result) []*dataloader.Result {
	for i := range results {
		result := dataloader.Result{
			Data: nil,
		}
		results[i] = &result
	}
	return results
}
