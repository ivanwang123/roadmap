package dataloaders

import (
	"time"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

func RoadmapById(db *sqlx.DB) *RoadmapLoader {
	return &RoadmapLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []int) ([]*model.Roadmap, []error) {
			roadmaps := []*model.Roadmap{}
			errors := Fetcher(db, "roadmaps", "id", ids, &roadmaps)

			return roadmaps, errors
		},
	}
}

func RoadmapByFollower(db *sqlx.DB) *RoadmapFollowerLoader {
	return &RoadmapFollowerLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(userIds []int) ([][]*model.Roadmap, []error) {
			roadmaps := make([][]*model.Roadmap, len(userIds))
			errors := make([]error, len(userIds))

			for i, userId := range userIds {
				if err := db.Unsafe().Select(&roadmaps[i], "SELECT * FROM roadmap_followers AS f LEFT JOIN roadmaps AS r ON f.roadmap_id = r.id WHERE f.user_id = $1", userId); err != nil {
					errors[i] = err
				}
			}

			return roadmaps, errors
		},
	}
}
