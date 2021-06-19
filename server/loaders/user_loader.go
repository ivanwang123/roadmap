package loaders

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/server/database"
	"github.com/ivanwang123/roadmap/server/graph/model"
)

func UserById(ctx context.Context) func(int) (*model.User, error) {
	loader := dataloader.NewBatchedLoader(userByIdBatchFn)

	return func(userId int) (*model.User, error) {
		result, err := loader.Load(ctx, dataloader.StringKey(strconv.Itoa(userId)))()
		if err != nil {
			fmt.Println("LOADER ERR", err)
			return nil, err
		}

		user, _ := result.(*model.User)
		return user, nil
	}
}

func userByIdBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	usersMap := make(map[string]interface{})
	users := []*model.User{}

	err := Fetcher(ctx, "users", "id", keys, &users)
	if err != nil {
		return HandleError(err, results)
	}

	for _, user := range users {
		usersMap[strconv.Itoa(user.ID)] = user
	}

	SortResults(results, keys, usersMap)

	return results
}

func UserByRoadmapFollowing(ctx context.Context) func(int) ([]*model.User, error) {
	loader := dataloader.NewBatchedLoader(userByRoadmapFollowingBatchFn)
	return func(roadmapId int) ([]*model.User, error) {
		result, err := loader.Load(ctx, dataloader.StringKey(strconv.Itoa(roadmapId)))()
		if err != nil {
			return nil, err
		}

		users, _ := result.([]*model.User)
		return users, nil
	}
}

func userByRoadmapFollowingBatchFn(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, len(keys))
	usersList := make([][]*model.User, len(keys))

	db := database.ForContext(ctx)

	for i, key := range ToStrKeys(keys) {
		if err := db.Unsafe().Select(&usersList[i], "SELECT * FROM roadmap_followers AS f LEFT JOIN users AS u ON f.user_id = u.id WHERE f.roadmap_id = $1", key); err != nil {
			fmt.Println(err)
			usersList[i] = []*model.User{}
		}
	}

	for i, users := range usersList {
		results[i] = &dataloader.Result{
			Data: users,
		}
	}

	return results
}
