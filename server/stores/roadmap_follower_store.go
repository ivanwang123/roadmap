package stores

import (
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type RoadmapFollowerStore struct {
	DB *sqlx.DB
}

func (s *RoadmapFollowerStore) Get(userId, roadmapId int) (*model.RoadmapFollower, error) {
	var roadmapFollower model.RoadmapFollower
	if err := s.DB.Get(&roadmapFollower, "SELECT * FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2 LIMIT 1", userId, roadmapId); err != nil {
		return nil, err
	}
	return &roadmapFollower, nil
}

// TODO: Put in transaction?
func (s *RoadmapFollowerStore) ToggleFollowRoadmap(store *Store, userId, roadmapId int) (*model.Roadmap, error) {
	if _, err := s.DB.Exec("INSERT INTO roadmap_followers (user_id, roadmap_id) VALUES ($1, $2)", userId, roadmapId); err != nil {
		if _, err := s.DB.Exec("DELETE FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2", userId, roadmapId); err != nil {
			return nil, err
		}

		checkpoints := make([]int64, 0)
		err := s.DB.Select(&checkpoints, "SELECT id FROM checkpoints WHERE roadmap_id = $1", roadmapId)
		if err != nil {
			return nil, err
		}

		for _, checkpointId := range checkpoints {
			// TODO: Batch request
			s.DB.Exec("DELETE FROM checkpoint_status WHERE user_id = $1 AND checkpoint_id = $2 AND roadmap_id = $3", userId, checkpointId, roadmapId)
		}

		return store.RoadmapStore.GetById(roadmapId)
	}

	checkpoints := make([]int64, 0)
	err := s.DB.Select(&checkpoints, "SELECT id FROM checkpoints WHERE roadmap_id = $1", roadmapId)
	if err != nil {
		return nil, err
	}

	for _, checkpointId := range checkpoints {
		// TODO: Batch request
		s.DB.Exec("INSERT INTO checkpoint_status (user_id, checkpoint_id, roadmap_id) VALUES ($1, $2, $3)", userId, checkpointId, roadmapId)
	}

	return store.RoadmapStore.GetById(roadmapId)
}
