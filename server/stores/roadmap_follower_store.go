package stores

import (
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type RoadmapFollowerStore struct {
	*sqlx.DB
}

// TODO: Get user id from auth context
func (s *RoadmapFollowerStore) ToggleFollowRoadmap(input *model.FollowRoadmap) (bool, error) {
	if _, err := s.Exec("INSERT INTO roadmap_followers (user_id, roadmap_id) VALUES ($1, $2)", input.UserID, input.RoadmapID); err != nil {
		if _, err := s.Exec("DELETE FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2", input.UserID, input.RoadmapID); err != nil {
			return false, err
		}
		return false, nil
	}

	// TODO: Get roadmap and add checkpoint status to all checkpoints
	checkpoints := make([]int64, 0)
	err := s.Select(&checkpoints, "SELECT id FROM checkpoints WHERE roadmap_id = $1", input.RoadmapID)
	fmt.Println("CHECKPOINTS", checkpoints, err)
	if err != nil {
		return true, err
	}

	for _, checkpointID := range checkpoints {
		// TODO: Batch request in transaction
		_, err := s.Exec("INSERT INTO checkpoint_status (user_id, checkpoint_id, roadmap_id) VALUES ($1, $2, $3)", input.UserID, checkpointID, input.RoadmapID)
		if err != nil {
			return true, err
		}
	}

	return true, nil
}
