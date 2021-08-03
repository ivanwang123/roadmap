package loaders

import (
	"context"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/database"
	"github.com/ivanwang123/roadmap/models"
)

func RoadmapById(ctx context.Context) func(int) (*models.Roadmap, error) {
	loader := dataloader.NewBatchedLoader(roadmapByIdBatchFn)

	return func(roadmapId int) (*models.Roadmap, error) {
		result, err := loader.Load(ctx, dataloader.StringKey(strconv.Itoa(roadmapId)))()
		if err != nil {
			return nil, err
		}

		roadmap, _ := result.(*models.Roadmap)
		return roadmap, nil
	}
}

func roadmapByIdBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	roadmapsMap := make(map[string]interface{})
	roadmaps := []*models.Roadmap{}

	if err := Fetcher(ctx, "roadmaps", "id", keys, &roadmaps); err != nil {
		return HandleError(err, results)
	}

	for _, roadmap := range roadmaps {
		roadmapsMap[strconv.Itoa(roadmap.ID)] = roadmap
	}

	SortResults(results, keys, roadmapsMap)

	return results
}

func RoadmapByFollower(ctx context.Context) func(int) ([]*models.Roadmap, error) {
	loader := dataloader.NewBatchedLoader(roadmapByFollowerBatchFn)

	return func(userId int) ([]*models.Roadmap, error) {
		result, err := loader.Load(ctx, dataloader.StringKey(strconv.Itoa(userId)))()
		if err != nil {
			return nil, err
		}

		roadmaps, _ := result.([]*models.Roadmap)
		return roadmaps, nil
	}
}

func roadmapByFollowerBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	roadmapsList := make([][]*models.Roadmap, len(keys))

	db := database.ForContext(ctx)

	for i, key := range ToStrKeys(keys) {
		if err := db.Unsafe().Select(&roadmapsList[i], "SELECT * FROM roadmap_followers AS f LEFT JOIN roadmaps AS r ON f.roadmap_id = r.id WHERE f.user_id = $1", key); err != nil {
			roadmapsList[i] = []*models.Roadmap{}
		}
	}

	for i, roadmaps := range roadmapsList {
		results[i] = &dataloader.Result{
			Data: roadmaps,
		}
	}

	return results
}
