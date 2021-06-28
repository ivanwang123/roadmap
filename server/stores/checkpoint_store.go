package stores

import (
	"fmt"

	"github.com/ivanwang123/roadmap/server/graph/model"
	"github.com/jmoiron/sqlx"
)

type CheckpointStore struct {
	*sqlx.DB
}

func (s *CheckpointStore) Create(input *model.NewCheckpoint) (*model.Checkpoint, error) {
	// TODO: Get roadmap followers and add checkpoint status
	var checkpoint model.Checkpoint
	if err := s.Get(&checkpoint, "INSERT INTO checkpoints (title, instructions, links, roadmap_id) VALUES ($1, $2, $3, $4) RETURNING *",
		input.Title, input.Instructions, input.Links, input.RoadmapID); err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (s *CheckpointStore) GetByRoadmap(roadmapId int) ([]*model.Checkpoint, error) {
	checkpoints := []*model.Checkpoint{}
	if err := s.Select(&checkpoints, "SELECT * FROM checkpoints WHERE roadmap_id = $1",
		roadmapId); err != nil {
		return nil, err
	}
	return checkpoints, nil
}

func (s *CheckpointStore) UpdateStatus(userId int, input *model.UpdateStatus) (*model.Checkpoint, error) {
	if _, err := s.Exec("UPDATE checkpoint_status SET status = $1 WHERE user_id = $2 AND checkpoint_id = $3", input.Status, userId, input.CheckpointID); err != nil {
		return nil, nil
	}

	var checkpoint model.Checkpoint
	if err := s.Get(&checkpoint, "SELECT * FROM checkpoints WHERE id = $1", input.CheckpointID); err != nil {
		return nil, nil
	}
	checkpoint.Status = model.Status(input.Status)
	fmt.Printf("%+v \n", checkpoint)
	return &checkpoint, nil
}
