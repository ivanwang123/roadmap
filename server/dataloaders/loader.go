package dataloaders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

const loaderCtxKey = "dataloader"

type Loader struct {
	UserById               *UserLoader
	UserByRoadmapFollowing *UserRoadmapFollowingLoader
	RoadmapById            *RoadmapLoader
	RoadmapByFollower      *RoadmapFollowerLoader
}

func Middleware(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loaderCtxKey, &Loader{
				UserById:               UserById(db),
				UserByRoadmapFollowing: UserByRoadmapFollowing(db),
				RoadmapById:            RoadmapById(db),
				RoadmapByFollower:      RoadmapByFollower(db),
			})

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *Loader {
	raw, _ := ctx.Value(loaderCtxKey).(*Loader)
	return raw
}

func Fetcher(db *sqlx.DB, table string, column string, ids []int, dest interface{}) []error {
	errors := make([]error, len(ids))

	queryStr := fmt.Sprintf("SELECT * FROM %s WHERE %s IN (?)", table, column)
	query, args, err := sqlx.In(queryStr, ids)
	if err != nil {
		for i := range errors {
			errors[i] = err
		}
	}

	if err := db.Select(dest, db.Rebind(query), args...); err != nil {
		for i := range errors {
			errors[i] = err
		}
	}

	return errors
}
