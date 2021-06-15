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

func UserByRoadmapFollowing(db *sqlx.DB) *UserRoadmapFollowingLoader {
	return &UserRoadmapFollowingLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(roadmapIds []int) ([][]*model.User, []error) {
			users := make([][]*model.User, len(roadmapIds))
			errors := make([]error, len(roadmapIds))

			for i, roadmapId := range roadmapIds {
				if err := db.Unsafe().Select(&users[i], "SELECT * FROM roadmap_followers AS f LEFT JOIN users AS u ON f.user_id = u.id WHERE f.roadmap_id = $1", roadmapId); err != nil {
					errors[i] = err
				}
			}

			return users, errors
		},
	}
}
