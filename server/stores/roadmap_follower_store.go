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
func (s *RoadmapFollowerStore) ToggleFollowRoadmap(store *Store, userId, roadmapId int) (*model.Roadmap, error) {
	if _, err := s.Exec("INSERT INTO roadmap_followers (user_id, roadmap_id) VALUES ($1, $2)", userId, roadmapId); err != nil {
		if _, err := s.Exec("DELETE FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2", userId, roadmapId); err != nil {
			return nil, err
		}

		// TODO: Delete checkpoint statuses?
		fmt.Println("UNFOLLOW")
		return store.RoadmapStore.GetById(roadmapId)
	}

	// TODO: Get roadmap and add checkpoint status to all checkpoints
	checkpoints := make([]int64, 0)
	err := s.Select(&checkpoints, "SELECT id FROM checkpoints WHERE roadmap_id = $1", roadmapId)
	fmt.Println("CHECKPOINTS", checkpoints, err)
	if err != nil {
		return nil, err
	}

	for _, checkpointId := range checkpoints {
		// TODO: Batch request in transaction
		s.Exec("INSERT INTO checkpoint_status (user_id, checkpoint_id, roadmap_id) VALUES ($1, $2, $3)", userId, checkpointId, roadmapId)
		// if err != nil {
		// 	return nil, err
		// }
	}

	fmt.Println("FOLLOW")
	return store.RoadmapStore.GetById(roadmapId)
}
