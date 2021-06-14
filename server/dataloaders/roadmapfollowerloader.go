package dataloaders

import (
	"time"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

func RoadmapFollowerByUserId(db *sqlx.DB) *RoadmapFollowerLoader {
	return &RoadmapFollowerLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(userIds []int) ([]*model.RoadmapFollower, []error) {
			followers := []*model.RoadmapFollower{}
			errors := Fetcher(db, "roadmap_followers", "user_id", userIds, &followers)

			return followers, errors
		},
	}
}

func RoadmapFollowerByRoadmapId(db *sqlx.DB) *RoadmapFollowerLoader {
	return &RoadmapFollowerLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(roadmapIds []int) ([]*model.RoadmapFollower, []error) {
			followers := []*model.RoadmapFollower{}
			errors := Fetcher(db, "roadmap_followers", "roadmap_id", roadmapIds, &followers)

			return followers, errors
		},
	}
}
