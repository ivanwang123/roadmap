package dataloaders

import (
	"time"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

func UserDataloader(db *sqlx.DB) *UserLoader {
	return &UserLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []int) ([]*model.User, []error) {
			users := []*model.User{}
			errors := []error{}

			query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)
			if err != nil {
				errors = append(errors, err)
			}

			if err := db.Select(&users, db.Rebind(query), args...); err != nil {
				errors = append(errors, err)
			}

			return users, errors
		},
	}
}
