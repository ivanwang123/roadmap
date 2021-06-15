package stores

import (
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
	return true, nil
}
