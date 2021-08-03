package loaders

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/graph-gophers/dataloader"
	"github.com/ivanwang123/roadmap/database"
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

const loaderCtxKey = "loader"

type Loader struct {
	UserById               func(int) (*models.User, error)
	UserByRoadmapFollowing func(int) ([]*models.User, error)
	RoadmapById            func(int) (*models.Roadmap, error)
	RoadmapByFollower      func(int) ([]*models.Roadmap, error)
	CheckpointStatus       func(int) (*models.CheckpointStatus, error)
}

func Middleware(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loader := &Loader{
				UserById:               UserById(r.Context()),
				UserByRoadmapFollowing: UserByRoadmapFollowing(r.Context()),
				RoadmapById:            RoadmapById(r.Context()),
				RoadmapByFollower:      RoadmapByFollower(r.Context()),
				CheckpointStatus:       CheckpointStatus(r.Context()),
			}
			ctx := context.WithValue(r.Context(), loaderCtxKey, loader)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Loader {
	raw, _ := ctx.Value(loaderCtxKey).(*Loader)
	return raw
}

func Fetcher(ctx context.Context, table, column string, keys dataloader.Keys, dest interface{}) error {
	db := database.ForContext(ctx)

	queryStr := fmt.Sprintf("SELECT * FROM %s WHERE %s IN (?)", table, column)
	query, args, err := sqlx.In(queryStr, ToStrKeys(keys))
	if err != nil {
		return err
	}

	if err := db.Select(dest, db.Rebind(query), args...); err != nil {
		return err
	}

	return nil
}

func SortResults(results []*dataloader.Result, keys dataloader.Keys, dataMap map[string]interface{}) {
	for i, key := range keys {
		data, ok := dataMap[key.String()]
		var result dataloader.Result
		if !ok {
			result.Error = errors.New("Not found")
		} else {
			result.Data = data
		}

		results[i] = &result
	}
}

func HandleError(err error, results []*dataloader.Result) []*dataloader.Result {
	for i := range results {
		result := dataloader.Result{
			Error: err,
		}
		results[i] = &result
	}
	return results
}

func ToStrKeys(keys dataloader.Keys) []string {
	strKeys := make([]string, len(keys))

	for i, key := range keys {
		strKeys[i] = key.String()
	}

	return strKeys
}
