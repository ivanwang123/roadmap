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
