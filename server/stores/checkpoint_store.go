package stores

import (
	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type CheckpointStore struct {
	*sqlx.DB
}

func (s *CheckpointStore) Create(input *model.NewCheckpoint) (*model.Checkpoint, error) {
	var checkpoint model.Checkpoint
	if err := s.Get(&checkpoint, "INSERT INTO checkpoints (title, instructions, roadmap_id) VALUES ($1, $2, $3)",
		input.Title, input.Instructions, input.RoadmapID); err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (s *CheckpointStore) GetByRoadmap(roadmapId int) ([]*model.Checkpoint, error) {
	checkpoints := []*model.Checkpoint{}
	if err := s.Get(&checkpoints, "SELECT * FROM checkpoints WHERE roadmap_id = $1",
		roadmapId); err != nil {
		return nil, err
	}
	return checkpoints, nil
}
