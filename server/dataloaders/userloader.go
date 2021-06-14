package dataloaders

import (
	"time"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

func UserById(db *sqlx.DB) *UserLoader {
	return &UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []int) ([]*model.User, []error) {
			users := []*model.User{}
			errors := Fetcher(db, "users", "id", ids, &users)

			return users, errors
		},
	}
}
