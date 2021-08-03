package stores

import (
	"github.com/ivanwang123/roadmap/models"
	"github.com/jmoiron/sqlx"
)

type RoadmapFollowerStore struct {
	DB *sqlx.DB
}

func (s *RoadmapFollowerStore) Get(userId, roadmapId int) (*models.RoadmapFollower, error) {
	var roadmapFollower models.RoadmapFollower
	if err := s.DB.Get(&roadmapFollower, "SELECT * FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2 LIMIT 1", userId, roadmapId); err != nil {
		return nil, err
	}
	return &roadmapFollower, nil
}

// TODO: Put in transaction?
func (s *RoadmapFollowerStore) ToggleFollowRoadmap(userId int, input *models.FollowRoadmap) (*models.Roadmap, error) {
	roadmapStore := &RoadmapStore{DB: s.DB}
	checkpointStatusStore := &CheckpointStatusStore{DB: s.DB}

	if _, err := s.DB.Exec("INSERT INTO roadmap_followers (user_id, roadmap_id) VALUES ($1, $2)", userId, input.RoadmapID); err != nil {
		if _, err := s.DB.Exec("DELETE FROM roadmap_followers WHERE user_id = $1 AND roadmap_id = $2", userId, input.RoadmapID); err != nil {
			return nil, err
		}

		checkpointIds := []int{}
		err := s.DB.Select(&checkpointIds, "SELECT id FROM checkpoints WHERE roadmap_id = $1", input.RoadmapID)
		if err != nil {
			return nil, err
		}

		deleteCheckpointStatusInput := &DeleteCheckpointStatus{
			roadmapId:     input.RoadmapID,
			userIds:       []int{userId},
			checkpointIds: checkpointIds,
		}
		if err := checkpointStatusStore.DeleteManyCheckpointStatus(deleteCheckpointStatusInput); err != nil {
			return nil, err
		}

		return roadmapStore.GetById(input.RoadmapID)
	}

	// TODO: Combine with above
	checkpointIds := []int64{}
	err := s.DB.Select(&checkpointIds, "SELECT id FROM checkpoints WHERE roadmap_id = $1", input.RoadmapID)
	if err != nil {
		return nil, err
	}

	newCheckpointStatuses := make([]*NewCheckpointStatus, len(checkpointIds))
	for i, checkpointId := range checkpointIds {
		newCheckpointStatuses[i] = &NewCheckpointStatus{userId: userId, checkpointId: int(checkpointId), roadmapId: input.RoadmapID}
	}
	if err := checkpointStatusStore.CreateManyCheckpointStatus(newCheckpointStatuses); err != nil {
		return nil, err
	}

	return roadmapStore.GetById(input.RoadmapID)
}
